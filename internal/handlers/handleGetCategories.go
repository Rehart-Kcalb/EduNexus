package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetAllCategories(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := DB.AllCategories(context.Background())
		if err != nil {
			types.NewJsonResponse("Internal error with database", http.StatusInternalServerError)
		}
		types.NewJsonResponse(categories, http.StatusOK).Respond(w)
	}
}
