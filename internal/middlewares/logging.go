package middlewares

import (
	"calculator_final/internal/entities"
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(entities.User)
		if !ok {
			http.Error(w, "no user in context", http.StatusForbidden)
			return
		}
		fmt.Println("AuthMiddleware: " + user.Username)
		next.ServeHTTP(w, r)
	}
}
