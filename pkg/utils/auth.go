package utils

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	database "filmLibrary"
	"filmLibrary/internal/models"
	"net/http"
	"strings"
)

// Authenticate проверяет данные аутентификации
func Authenticate(username, password string) (models.User, error) {
	pool, err := database.GetPool()
	if err != nil {
		return models.User{}, err
	}
	defer pool.Close()

	var user models.User
	row := pool.QueryRow(context.Background(), "SELECT id, is_admin FROM users WHERE username = $1 AND password = $2", username, password)
	err = row.Scan(&user.ID, &user.IsAdmin)
	switch {
	case err == sql.ErrNoRows:
		return models.User{}, errors.New("user not found or invalid credentials")
	case err != nil:
		return models.User{}, err
	default:
		return user, nil
	}
}

// BasicAuthMiddleware проверяет данные Basic Auth и роль пользователя
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

		user, err := Authenticate(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "UserRole", user.IsAdmin)

		next(w, r.WithContext(ctx))
	}
}
