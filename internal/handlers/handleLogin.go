package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func HandleLogin(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			types.NewJsonResponse(struct {Message any `json:"error_message"`}{"Problem with server"}, http.StatusInternalServerError).Respond(w)
			//w.WriteHeader(http.StatusInternalServerError)
			return
		}
		login := r.PostFormValue("login")
		log.Println(login)
		if is_login_valid := utils.ValidateLogin(login); is_login_valid != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{is_login_valid.Error()}, http.StatusUnauthorized).Respond(w)
			return
		}
		password := r.PostFormValue("password")
		if is_password_valid := utils.ValidatePassword(password); is_password_valid != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{is_password_valid.Error()}, http.StatusUnauthorized).Respond(w)
			return
		}
		hash_password, err := DB.GetPasswordByLogin(context.Background(), login)
		if err != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{err.Error()}, http.StatusInternalServerError).Respond(w)
			return
		}
		if is_correct := utils.CheckPassword(password, hash_password); !is_correct {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{"Password or Login is wrong"}, http.StatusUnauthorized).Respond(w)
			return
		}
		claims, err := DB.GetClaimsByLogin(context.Background(), login)
		if err != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{err.Error()}, http.StatusInternalServerError).Respond(w)
			return
		}

		claim := (types.NewClaims(int(claims.ID), claims.Title.String))
		token, err := json.Marshal(*claim)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(token)
		utils.Encrypt(&token)
		log.Println(token)
		types.NewJsonResponse(struct {
			Token any `json:"token"`
		}{(token)}, http.StatusOK).Respond(w)
	}
}
