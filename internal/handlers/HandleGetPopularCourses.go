package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetPopularCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := DB.GetPopularCourses(context.Background())
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Ошибка при получении данных с БД"}, http.StatusInternalServerError).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Data any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
