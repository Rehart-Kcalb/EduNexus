package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/types"
	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/utils"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) < len("Bearer ") {
			log.Println(token)
			http.Redirect(w, r, "/api/login", 301)
			return
		}
		if is_valid := VerifyToken(token); !is_valid {
			log.Println(token)
			http.Redirect(w, r, "/api/login", 301)
			return
		}
		crypto := (token[len("Bearer "):])
		b, err := base64.StdEncoding.DecodeString(crypto)
		if err != nil {
			return
		}
		utils.Decrypt((&(b)))
		var claims types.Claims
		json.Unmarshal((b), &claims)
		log.Println(claims)
		ctx := context.WithValue(r.Context(), "id", int64(claims.Id))

		handler(w, r.WithContext(ctx))
	}
}

func VerifyToken(token string) bool {
	return true
}
