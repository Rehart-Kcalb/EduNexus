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

func handleMatchingSubmission(w http.ResponseWriter, r *http.Request, assignment db.GetAssignmentByIdRow, DB *db.Queries) {
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
	if err := json.Unmarshal([]byte(assignment.Content), &matchingQuiz); err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Ошибка при парсинге задания"}, http.StatusInternalServerError).Respond(w)
		return
	}
	var userMatchingSubmission UserMatchingSubmission
	if err := json.NewDecoder(r.Body).Decode(&userMatchingSubmission); err != nil {
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Ошибка при парсинге подачи"}, http.StatusBadRequest).Respond(w)
		return
	}
	if len(matchingQuiz.MatchingPairs) != len(userMatchingSubmission.MatchingPairs) {
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"Невалидная подаяа"}, http.StatusBadRequest).Respond(w)
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
			Message string `json:"message"`
		}{"Ошибка при создания попытки"}, http.StatusInternalServerError).Respond(w)
		return
	}
}
