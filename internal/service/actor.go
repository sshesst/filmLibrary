package service

import (
	"bytes"
	"context"
	"filmLibrary/internal/models"
	database "filmLibrary/internal/storage"
	"filmLibrary/pkg/logging"
	"fmt"
)

func AddActorToDB(actor models.Actor, logger logging.Logger) error {
	pool, err := database.GetPool()
	if err != nil {
		logger.Error("Ошибка получения пула соединений:", err)
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), `
		INSERT INTO actors (name, gender, birthdate)
		VALUES ($1, $2, $3)`,
		actor.Name, actor.Gender, actor.Birthdate)
	if err != nil {
		logger.Error("Ошибка подготовки SQL-запроса в добавлении актёра в бд", err)
		return err
	}
	logger.Info("Актер успешно добавлен в БД")

	return nil
}

func UpdateActorInDB(actor models.Actor, logger logging.Logger) error {
	pool, err := database.GetPool()
	if err != nil {
		logger.Error("Ошибка получения пула соединений:", err)
		return err
	}
	defer pool.Close()

	var updateQuery bytes.Buffer
	updateQuery.WriteString("UPDATE actors SET ")

	var params []interface{}
	var index int

	if actor.Name != "" {
		index++
		updateQuery.WriteString(fmt.Sprintf("name = $%d", index))
		params = append(params, actor.Name)
	}

	if actor.Gender != "" {
		if index > 0 {
			updateQuery.WriteString(", ")
		}
		index++
		updateQuery.WriteString(fmt.Sprintf("gender = $%d", index))
		params = append(params, actor.Gender)
	}

	if actor.Birthdate != "" {
		if index > 0 {
			updateQuery.WriteString(", ")
		}
		index++
		updateQuery.WriteString(fmt.Sprintf("birthdate = $%d", index))
		params = append(params, actor.Birthdate)
	}

	updateQuery.WriteString(" WHERE id = $")
	index++
	updateQuery.WriteString(fmt.Sprintf("%d", index))
	params = append(params, actor.ID)

	_, err = pool.Exec(context.Background(), updateQuery.String(), params...)
	if err != nil {
		logger.Error("Ошибка подготовки SQL-запроса:", err)
		return err
	}

	logger.Info("Актер успешно обновлен в БД")

	return nil
}

func DeleteActorFromDB(actorID uint, logger logging.Logger) error {
	pool, err := database.GetPool()
	if err != nil {
		logger.Error("Ошибка получения пула соединений:", err)
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), `
    DELETE FROM movie_actors
    WHERE actor_id = $1`, actorID)
	if err != nil {
		logger.Error("Ошибка удаления записей из movie_actors:", err)
		return err
	}

	_, err = pool.Exec(context.Background(), `
    DELETE FROM actors
    WHERE id = $1`, actorID)
	if err != nil {
		logger.Error("Ошибка удаления актера из БД:", err)
		return err
	}

	logger.Info("Актер успешно удален из БД")
	return nil
}
