package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	database "filmLibrary"
	"filmLibrary/internal/middleware"
	"filmLibrary/internal/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = middleware.ValidateMovie(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = addMovieToDB(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Фильм успешно добавлен")
}

func addMovieToDB(movie models.Movie) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	// Проверяем, существует ли фильм с таким же названием и датой выпуска
	var count int
	err = pool.QueryRow(context.Background(), `
        SELECT COUNT(*)
        FROM movies
        WHERE title = $1 AND release_date = $2`,
		movie.Title, movie.ReleaseDate).Scan(&count)
	if err != nil {
		return err
	}

	// Если фильм уже существует, возвращаем ошибку
	if count > 0 {
		return errors.New("такой фильм уже существует")
	}

	// 1. Вставляем фильм в таблицу movies
	_, err = pool.Exec(context.Background(), `
        INSERT INTO movies (title, description, release_date, rating)
        VALUES ($1, $2, $3, $4)`,
		movie.Title, movie.Description, movie.ReleaseDate, movie.Rating)
	if err != nil {
		return err
	}

	// 2. Получаем ID вставленного фильма
	var movieID uint
	err = pool.QueryRow(context.Background(), `
        SELECT id FROM movies WHERE title = $1 AND release_date = $2`,
		movie.Title, movie.ReleaseDate).Scan(&movieID)
	if err != nil {
		return err
	}

	// 3. Вставляем записи в таблицу movie_actors для связи с актерами
	for _, actor := range movie.Actors {
		_, err = pool.Exec(context.Background(), `
            INSERT INTO movie_actors (movie_id, actor_id)
            VALUES ($1, $2)`,
			movieID, actor.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateMovieInDB(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Информация о фильме успешно обновлена")
}

func updateMovieInDB(movie models.Movie) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	var updateQuery bytes.Buffer
	updateQuery.WriteString("UPDATE movies SET ")

	var params []interface{}
	var index int

	if movie.Title != "" {
		index++
		updateQuery.WriteString(fmt.Sprintf("title = $%d", index))
		params = append(params, movie.Title)
	}

	if movie.Description != "" {
		if index > 0 {
			updateQuery.WriteString(", ")
		}
		index++
		updateQuery.WriteString(fmt.Sprintf("description = $%d", index))
		params = append(params, movie.Description)
	}

	if movie.ReleaseDate != "" {
		if index > 0 {
			updateQuery.WriteString(", ")
		}
		index++
		updateQuery.WriteString(fmt.Sprintf("release_date = $%d", index))
		params = append(params, movie.ReleaseDate)
	}

	if movie.Rating != 0 {
		if index > 0 {
			updateQuery.WriteString(", ")
		}
		index++
		updateQuery.WriteString(fmt.Sprintf("rating = $%d", index))
		params = append(params, movie.Rating)
	}

	updateQuery.WriteString(" WHERE id = $")
	index++
	updateQuery.WriteString(fmt.Sprintf("%d", index))
	params = append(params, movie.ID)

	// Выполняем обновление записи о фильме
	_, err = pool.Exec(context.Background(), updateQuery.String(), params...)
	if err != nil {
		return err
	}

	// Удаляем существующие записи в таблице movie_actors для этого фильма
	_, err = pool.Exec(context.Background(), `
        DELETE FROM movie_actors
        WHERE movie_id = $1`,
		movie.ID)
	if err != nil {
		return err
	}

	// Вставляем новые записи в таблицу movie_actors для обновленного списка актеров
	for _, actor := range movie.Actors {
		_, err = pool.Exec(context.Background(), `
            INSERT INTO movie_actors (movie_id, actor_id)
            VALUES ($1, $2)`,
			movie.ID, actor.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр id из URL
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию для удаления фильма из базы данных по ID
	err = deleteMovieFromDB(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Информация о фильме успешно удалена")
}

func deleteMovieFromDB(movieID uint) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	// Удаляем связанные записи в таблице movie_actors
	_, err = pool.Exec(context.Background(), `
        DELETE FROM movie_actors
        WHERE movie_id = $1`,
		movieID)
	if err != nil {
		return err
	}

	// Удаляем записи о фильме из таблицы movies
	_, err = pool.Exec(context.Background(), `
        DELETE FROM movies
        WHERE id = $1`,
		movieID)
	if err != nil {
		return err
	}

	return nil
}
