package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/jackc/pgx/v5/pgtype"
)

func HandleUpdateProfileInfo(DB *db.Queries) http.HandlerFunc {
	type Profile struct {
		FirstName   string `json:"firstname"`
		Description string `json:"description"`
		Profile     string `json:"profile"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var post_data Profile
		err := json.NewDecoder(r.Body).Decode(&post_data)
		if err != nil {
			return
		}
		err = DB.UpdateProfileInfo(context.Background(), db.UpdateProfileInfoParams{Firstname: pgtype.Text{String: post_data.FirstName}, Description: pgtype.Text{String: post_data.Description}, Profile: pgtype.Text{String: post_data.Profile}})
		if err != nil {
			return
		}
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
