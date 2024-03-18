package handler

import (
	"filmLibrary/internal/models"
	"fmt"
	"strings"
)

func MovieHandler(movie models.Movie) error {
	if len(strings.TrimSpace(movie.Title)) < 1 || len(strings.TrimSpace(movie.Title)) > 150 {
		return fmt.Errorf("Название фильма должно содержать от 1 до 150 символов")
	}

	if len(movie.Description) > 1000 {
		return fmt.Errorf("Описание фильма должно содержать не более 1000 символов")
	}

	if movie.Rating < 0 || movie.Rating > 10 {
		return fmt.Errorf("Рейтинг фильма должен быть в диапазоне от 0 до 10")
	}

	return nil
}
