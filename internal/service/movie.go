package service

import (
	"bytes"
	"context"
	"errors"
	"filmLibrary/internal/models"
	database "filmLibrary/internal/storage"
	"fmt"
)

func AddMovieToDB(movie models.Movie) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	var count int
	err = pool.QueryRow(context.Background(), `
        SELECT COUNT(*)
        FROM movies
        WHERE title = $1 AND release_date = $2`,
		movie.Title, movie.ReleaseDate).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("такой фильм уже существует")
	}

	_, err = pool.Exec(context.Background(), `
        INSERT INTO movies (title, description, release_date, rating)
        VALUES ($1, $2, $3, $4)`,
		movie.Title, movie.Description, movie.ReleaseDate, movie.Rating)
	if err != nil {
		return err
	}

	var movieID uint
	err = pool.QueryRow(context.Background(), `
        SELECT id FROM movies WHERE title = $1 AND release_date = $2`,
		movie.Title, movie.ReleaseDate).Scan(&movieID)
	if err != nil {
		return err
	}

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

func UpdateMovieInDB(movie models.Movie) error {
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

	_, err = pool.Exec(context.Background(), updateQuery.String(), params...)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), `
        DELETE FROM movie_actors
        WHERE movie_id = $1`,
		movie.ID)
	if err != nil {
		return err
	}

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

func DeleteMovieFromDB(movieID uint) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), `
        DELETE FROM movie_actors
        WHERE movie_id = $1`,
		movieID)
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), `
        DELETE FROM movies
        WHERE id = $1`,
		movieID)
	if err != nil {
		return err
	}

	return nil
}
