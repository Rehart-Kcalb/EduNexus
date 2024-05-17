package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func HandleRegister(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		login := r.PostFormValue("login")
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
		hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{"Problem with server"}, http.StatusUnauthorized).Respond(w)
			return
		}
		err = DB.CreateUser(context.Background(), db.CreateUserParams{Login: login, Password: string(hashed_password)})
		if err != nil {
			types.NewJsonResponse(struct {
				Message any `json:"error_message"`
			}{"Problem with server"}, http.StatusUnauthorized).Respond(w)
			return
		}
		types.NewJsonResponse(struct {
			Status string `json:"status"`
		}{"success"}, http.StatusOK).Respond(w)
	}
}
