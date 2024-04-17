package handlers

import (
	"context"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func handleLogin(DB *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		login := r.PostFormValue("login")
		if is_login_valid := utils.ValidateLogin(login); !is_login_valid {
			// TODO: RETURN MESSAGE ABOUT INVALID LOGIN
		}
		password := r.PostFormValue("password")
		if is_password_valid := utils.ValidatePassword(password); !is_password_valid {
			// TODO: RETURN MESSAGE ABOUT INVALID PASSWORD
		}
		hash_password, err := DB.GetPasswordByLogin(context.Background(), login)
		if err != nil {
			// TODO: RETURN PROBLEM WITH SERVER
		}
		_ = hash_password
		// TODO: WRITE FUNCTION FOR CHECKING PASSWORD
	}
}
