package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
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

		filePath, err := utils.SaveBase64ToFile(post_data.Profile, "storage")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error saving file: %v", err), http.StatusInternalServerError)
			return
		}
		log.Println(filePath)
		err = DB.UpdateProfileInfo(context.Background(), db.UpdateProfileInfoParams{Firstname: pgtype.Text{String: post_data.FirstName, Valid: true}, Description: pgtype.Text{String: post_data.Description, Valid: true}, Profile: pgtype.Text{String: filePath, Valid: true}, ID: r.Context().Value("id").(int64)})
		if err != nil {
			return
		}
		types.NewJsonResponse(struct {
			Message string `json:"message"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
