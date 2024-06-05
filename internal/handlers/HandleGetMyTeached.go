package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetMyTeached(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teacher_id := r.Context().Value("id").(int64)
		_ = teacher_id
		courses, err := DB.GetMyTeached(context.Background(), teacher_id)
		if err != nil {
			types.NewJsonResponse(struct {
				Courses any `json:"courses"`
			}{[]any{}}, http.StatusOK).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Courses any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
