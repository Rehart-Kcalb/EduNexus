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
		if categories == nil {
			w.Write([]byte(err.Error()))
			return
		}
		types.NewJsonResponse(struct {
			Categories any `json:"categories"`
		}{categories}, http.StatusOK).Respond(w)
	}
}
