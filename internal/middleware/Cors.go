package middleware

import "net/http"

func CORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, HEAD")

		if r.Method == http.MethodOptions {
			return
		}

		handler.ServeHTTP(w, r)
	})
}
