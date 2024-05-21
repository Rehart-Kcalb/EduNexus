package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/jackc/pgx/v5/pgtype"
)

func handleMatchingSubmission(w http.ResponseWriter, r *http.Request, assignment db.Assignment, DB *db.Queries) {
	type MatchingQuiz struct {
		MatchingPairs []struct {
			Left  string `json:"left"`
			Right string `json:"right"`
		} `json:"matching_pairs"`
	}
	type UserMatchingSubmission struct {
		MatchingPairs []struct {
			Left  string `json:"left"`
			Right string `json:"right"`
		} `json:"matching_pairs"`
	}
	var matchingQuiz MatchingQuiz
	if err := json.Unmarshal(assignment.Content, &matchingQuiz); err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"error message"`
		}{"Failed to parse assignment content"}, http.StatusInternalServerError).Respond(w)
		return
	}
	var userMatchingSubmission UserMatchingSubmission
	if err := json.NewDecoder(r.Body).Decode(&userMatchingSubmission); err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"error message"`
		}{"Failed to parse matching submission"}, http.StatusBadRequest).Respond(w)
		return
	}
	if len(matchingQuiz.MatchingPairs) != len(userMatchingSubmission.MatchingPairs) {
		types.NewJsonResponse(struct {
			Message string `json:"error message"`
		}{"Invalid submission"}, http.StatusBadRequest).Respond(w)
		return
	}
	wrong := 0
	for i, pair := range matchingQuiz.MatchingPairs {
		if pair.Left != userMatchingSubmission.MatchingPairs[i].Left || pair.Right != userMatchingSubmission.MatchingPairs[i].Right {
			wrong++
		}
	}

	err := DB.CreateSubmission(context.Background(), db.CreateSubmissionParams{AssignmentID: assignment.ID, UserID: 1, Info: pgtype.Text{String: fmt.Sprintf("%d/%d", len(matchingQuiz.MatchingPairs)-wrong, len(matchingQuiz.MatchingPairs))}})
	if err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"error message"`
		}{"Failed to create submission"}, http.StatusInternalServerError).Respond(w)
		return
	}
}
