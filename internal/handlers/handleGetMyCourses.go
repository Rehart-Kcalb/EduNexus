package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetMyCourses(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := DB.GetMyCourses(context.Background(), 1)
		if err != nil {
			// TODO: Handle Error
		}
		types.NewJsonResponse(struct {
			Data any `json:"courses"`
		}{courses}, http.StatusOK).Respond(w)
	}
}
