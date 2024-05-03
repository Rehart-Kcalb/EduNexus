package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := DB.GetCourses(context.Background())
		if err != nil {
			types.NewJsonResponse("Problem with DB", http.StatusInternalServerError).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Courses any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
