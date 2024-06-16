package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/jackc/pgx/v5/pgtype"
)

func handleQuizSubmission(w http.ResponseWriter, r *http.Request, assignment db.GetAssignmentByIdRow, DB *db.Queries) {
	type quiz struct {
		Questions []string   `json:"questions"`
		Variant   [][]string `json:"variant"`
		Answers   []string   `json:"answers"`
	}
	type UserAnswers struct {
		Answers []string `json:"answers"`
	}
	var quiz1 quiz
	if err := json.Unmarshal([]byte(assignment.Content), &quiz1); err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Ошибка при парсинге задания"}, http.StatusBadRequest).Respond(w)
		return
	}

	var user_answer UserAnswers
	if err := json.NewDecoder(r.Body).Decode(&user_answer); err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Ошибка при парсинге ответов"}, http.StatusBadRequest).Respond(w)
		return
	}
	var buff []byte = make([]byte, 1000)
	n, err := r.Body.Read(buff)
	if err != nil && !errors.Is(err, io.EOF) {
		_ = n
		log.Println(err)
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Ошибка при чтении тела запроса"}, http.StatusBadRequest).Respond(w)
		return
	}

	if len(quiz1.Answers) != len(user_answer.Answers) {
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Отправлены не все ответы"}, http.StatusBadRequest).Respond(w)
		return
	}

	grade := 0
	for i := range quiz1.Answers {
		if quiz1.Answers[i] == user_answer.Answers[i] {
			grade++
		}
	}

	// TODO: Save grade to the database
	//	fmt.Sprintf("%d/%d", grade, len(quiz1.Answers))
	DB.CreateSubmission(context.Background(), db.CreateSubmissionParams{Info: pgtype.Text{String: fmt.Sprintf("%d/%d", grade, len(quiz1.Answers))}, AssignmentID: assignment.ID, Content: buff, UserID: r.Context().Value("id").(int64)})
}
