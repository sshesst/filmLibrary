package test

import (
	"filmLibrary/internal/handler"
	"filmLibrary/internal/models"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMovieHandler(t *testing.T) {
	tests := []struct {
		name     string
		movie    models.Movie
		expected error
	}{
		{
			name: "Valid movie",
			movie: models.Movie{
				Title:       "Valid Movie",
				Description: "This is a valid movie description.",
				Rating:      7.5,
			},
			expected: nil,
		},
		{
			name: "Empty title",
			movie: models.Movie{
				Title:       "",
				Description: "This is a valid movie description.",
				Rating:      7.5,
			},
			expected: fmt.Errorf("Название фильма должно содержать от 1 до 150 символов"),
		},
		{
			name: "Long description",
			movie: models.Movie{
				Title:       "Valid Movie",
				Description: strings.Repeat("a", 1001),
				Rating:      7.5,
			},
			expected: fmt.Errorf("Описание фильма должно содержать не более 1000 символов"),
		},
		{
			name: "Invalid rating",
			movie: models.Movie{
				Title:       "Valid Movie",
				Description: "This is a valid movie description.",
				Rating:      15,
			},
			expected: fmt.Errorf("Рейтинг фильма должен быть в диапазоне от 0 до 10"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := handler.MovieHandler(tc.movie)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected error: %v, got: %v", tc.expected, result)
			}
		})
	}
}
