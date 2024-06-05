package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetProfileInfo(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user_prof_info, err := DB.GetProfileInfo(context.Background(), r.Context().Value("id").(int64))
		if err != nil {
			return
		}
		_ = user_prof_info
		types.NewJsonResponse(struct {
			Info any `json:"profile_info"`
		}{user_prof_info}, http.StatusOK).Respond(w)
	}
}
