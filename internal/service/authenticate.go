package service

import (
	"context"
	"database/sql"
	"errors"
	"filmLibrary/internal/models"
	database "filmLibrary/internal/storage"
)

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
		return models.User{}, errors.New("пользователь не найден")
	case err != nil:
		return models.User{}, err
	default:
		return user, nil
	}
}
