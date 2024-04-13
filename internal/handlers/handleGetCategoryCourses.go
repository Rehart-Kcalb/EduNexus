package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetCategoryCourses(DB db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category_name := r.PathValue("category_name")
		if category_name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		courses, err := DB.GetCategoryCourses(context.Background(), category_name)
		if err != nil {
			// TODO: Use json return error
			return
		}
		types.NewJsonResponse(courses, http.StatusOK).Respond(w)
	}
}
