package service

import (
	"bytes"
	"context"
	"filmLibrary/internal/models"
	database "filmLibrary/internal/storage"
	"fmt"
)

func AddActorToDB(actor models.Actor) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), `
		INSERT INTO actors (name, gender, birthdate)
		VALUES ($1, $2, $3)`,
		actor.Name, actor.Gender, actor.Birthdate)
	if err != nil {
		return err
	}

	return nil
}

func UpdateActorInDB(actor models.Actor) error {
	pool, err := database.GetPool()
	if err != nil {
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
		return err
	}

	return nil
}

func DeleteActorFromDB(actorID uint) error {
	pool, err := database.GetPool()
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(context.Background(), `
		DELETE FROM actors
		WHERE id = $1`,
		actorID)
	if err != nil {
		return err
	}

	return nil
}
