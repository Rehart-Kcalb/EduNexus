package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) < len("Bearer") {
			w.Header().Add("Location", "localhost/api/login")
			return
		}
		if is_valid := VerifyToken(token); !is_valid {
			w.Header().Add("Location", "localhost/api/login")
			return
		}
		crypto := []byte(token[len("Bearer"):])
		utils.Decrypt((&(crypto)))
		var claims types.Claims
		json.Unmarshal([]byte(crypto), &claims)
		r.WithContext(context.WithValue(r.Context(), "id", claims.Id))
		handler(w, r)
	}
}

func VerifyToken(token string) bool {
	return true
}
