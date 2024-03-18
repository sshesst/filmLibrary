package middleware

import (
	"context"
	"encoding/base64"
	"filmLibrary/internal/service"
	"net/http"
	"strings"
)

func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			http.Error(w, "Неавторизован", http.StatusUnauthorized)
			return
		}
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Неавторизован", http.StatusUnauthorized)
			return
		}
		username := pair[0]
		password := pair[1]

		user, err := service.Authenticate(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "UserRole", user.IsAdmin)

		next(w, r.WithContext(ctx))
	}
}
