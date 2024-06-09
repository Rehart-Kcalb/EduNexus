package handlers

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/jackc/pgx/v5/pgtype"
)

func handleCodeSubmission(w http.ResponseWriter, r *http.Request, assignment db.GetAssignmentByIdRow, DB *db.Queries) {
	type codeQuiz struct {
		Language     string `json:"language"`
		CodeTest     string `json:"code_test"`
		QuizQuestion string `json:"quiz_question"`
	}

	type CodeSub struct {
		Code string `json:"source"`
	}

	var codeQuiz1 codeQuiz
	/* log.Println(assignment.Content) */
	if err := json.Unmarshal([]byte(assignment.Content), &codeQuiz1); err != nil {
		http.Error(w, "Failed to parse assignment content", http.StatusInternalServerError)
		return
	}

	var codeSub1 CodeSub
	if err := json.NewDecoder(r.Body).Decode(&codeSub1); err != nil {
		http.Error(w, "Failed to parse code submission", http.StatusBadRequest)
		return
	}
	log.Println(codeSub1)

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "code_submission")
	if err != nil {
		http.Error(w, "Failed to create temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	// Write the user's code and test code to the temporary directory
	userCodeFilename, err := generateRandomFilename(5, codeQuiz1.Language, "")
	if err != nil {
		http.Error(w, "Failed to generate user code filename", http.StatusInternalServerError)
		return
	}
	testCodeFilename, err := generateRandomFilename(5, codeQuiz1.Language, "_test")
	if err != nil {
		http.Error(w, "Failed to generate test code filename", http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(tempDir+"/"+userCodeFilename, []byte(codeSub1.Code), 0644); err != nil {
		http.Error(w, "Failed to write user code to file", http.StatusInternalServerError)
		return
	}
	if err := os.WriteFile(tempDir+"/"+testCodeFilename, []byte(codeQuiz1.CodeTest), 0644); err != nil {
		http.Error(w, "Failed to write test code to file", http.StatusInternalServerError)
		return
	}

	// Run the tests in Docker
	expectedOutput, _, err := runCodeInDocker(tempDir, userCodeFilename, testCodeFilename, codeQuiz1.Language)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute code docker: %v", err), http.StatusInternalServerError)
		log.Println("expected output:", expectedOutput)
		log.Println(err)
		log.Println(codeQuiz1)
		return
	}

	log.Println(expectedOutput)
	testCaseNum := strings.Count(expectedOutput, "RUN")
	failCaseNum := strings.Count(expectedOutput, "FAIL")

	content, err := json.Marshal(struct {
		Content string `json:"content"`
	}{codeSub1.Code})
	if err != nil {
		log.Println(err)
	}
	// Store the submission in the database
	err = DB.CreateSubmission(r.Context(), db.CreateSubmissionParams{
		Content:      content,
		AssignmentID: assignment.ID,
		Info:         pgtype.Text{String: fmt.Sprintf("%d/%d", testCaseNum-failCaseNum, testCaseNum), Valid: true},
		UserID:       r.Context().Value("id").(int64),
	})
	if err != nil {
		http.Error(w, "Failed to save submission to the database", http.StatusInternalServerError)
		log.Println(`{"content":"` + codeSub1.Code + `"}`)
		log.Println(err)
		return
	}

	types.NewJsonResponse(struct {
		Passed string `json:"test_passed"`
	}{fmt.Sprintf("%d/%d", testCaseNum-failCaseNum, testCaseNum)}, http.StatusOK).Respond(w)
}

func runCodeInDocker(tempDir, userCodeFilename, testCodeFilename, language string) (string, string, error) {
	language = strings.ToLower(language)
	switch language {
	case "go":
		goModPath := tempDir + "/go.mod"
		log.Println(goModPath)
		if _, err := os.Stat(goModPath); os.IsNotExist(err) {
			if err := os.WriteFile(goModPath, []byte("module gotest\n\ngo 1.22.0"), 0644); err != nil {
				return "", "", fmt.Errorf("failed to create go.mod: %v", err)
			}
		} else if err != nil {
			return "", "", fmt.Errorf("error checking go.mod: %v", err)
		}

		// Docker client initialization
		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			return "", "", fmt.Errorf("failed to create Docker client: %v", err)
		}

		// Pull the golang:1.22 image if not already present
		if err := pullImage(cli, "golang:1.22"); err != nil {
			return "", "", fmt.Errorf("failed to pull Docker image: %v", err)
		}

		// Create a container to run the tests
		resp, err := cli.ContainerCreate(
			context.Background(),
			&container.Config{
				Image:      "golang:1.22",
				Cmd:        []string{"go", "test", "-v"},
				WorkingDir: "/usr/src/app",
				Volumes: map[string]struct{}{
					tempDir: {},
				},
			},
			&container.HostConfig{
				Binds: []string{fmt.Sprintf("%s:/usr/src/app", tempDir)},
			},
			nil,
			nil,
			"",
		)
		if err != nil {
			return "", "", fmt.Errorf("failed to create Docker container: %v", err)
		}

		// Start the container
		if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
			return "", "", fmt.Errorf("failed to start Docker container: %v", err)
		}

		// Wait for the container to finish
		statusCh, errCh := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
		select {
		case err := <-errCh:
			if err != nil {
				return "", "", fmt.Errorf("failed while waiting for Docker container: %v", err)
			}
		case <-statusCh:
		}

		// Retrieve the logs from the container
		logsOptions := container.LogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Since:      "",
			Until:      "",
			Timestamps: true,
			Follow:     false,
			Tail:       "",
			Details:    true,
		}
		out, err := cli.ContainerLogs(context.Background(), resp.ID, logsOptions)
		if err != nil {
			return "", "", fmt.Errorf("failed to retrieve Docker container logs: %v", err)
		}
		defer out.Close()

		output := new(strings.Builder)
		if _, err := stdcopy.StdCopy(output, output, out); err != nil {
			return "", "", fmt.Errorf("failed to copy Docker container logs: %v", err)
		}

		return output.String(), "", nil

	default:
		return "", "", fmt.Errorf("unsupported language: %s", language)
	}
}

func pullImage(cli *client.Client, image_ string) error {
	out, err := cli.ImagePull(context.Background(), image_, image.PullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(os.Stdout, out)

	return nil
}

func generateRandomFilename(length int, extension string, suffix string) (string, error) {
	randBytes := make([]byte, length)
	if _, err := rand.Read(randBytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("random-%x%s.%s", randBytes, suffix, extension), nil
}
