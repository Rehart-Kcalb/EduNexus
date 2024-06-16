package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetMyCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("id").(int64)
		courses, err := DB.GetMyCourses(context.Background(), user_id)
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Ошибка при получении данных с БД"}, http.StatusInternalServerError).Respond(w)
			// TODO: Handle Error
			return
		}
		types.NewJsonResponse(struct {
			Data any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
