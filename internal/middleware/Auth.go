package middleware

import "net/http"

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			// TODO: Add redirect
		}
		if len(token) < len("Bearer") {
			// TODO: Add redirect
		}
		if is_valid := VerifyToken(token); !is_valid {
			// TODO: Add redirect
		}
		claims := "Sigma"
		w.Header().Add("Claims", claims)
		handler(w, r)
	}
}

func VerifyToken(token string) bool {
	return true
}
