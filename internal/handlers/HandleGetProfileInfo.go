package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
)

func HandleGetProfileInfo(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user_id int64
		user_name := r.PathValue("user_name")
		if user_name != "" {
			claims, err := DB.GetClaimsByLogin(context.Background(), user_name)
			if err != nil {
				types.NewJsonResponse(struct {
					Message string `json:"message"`
				}{"Нет такого пользователя"}, http.StatusBadRequest).Respond(w)
				return
			}
			user_id = claims.ID
		} else {
			user_id = r.Context().Value("id").(int64)
		}

		user_prof_info, err := DB.GetProfileInfo(context.Background(), user_id)
		if err != nil {
			types.NewJsonResponse(struct {
				Message string `json:"message"`
			}{"Ошибка при получении данных пользователя"}, http.StatusInternalServerError).Respond(w)
			return
		}
		_ = user_prof_info
		types.NewJsonResponse(struct {
			Info any `json:"profile_info"`
		}{user_prof_info}, http.StatusOK).Respond(w)
	}
}
