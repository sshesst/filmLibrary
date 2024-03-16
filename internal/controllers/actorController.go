package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	database "filmLibrary"
	"filmLibrary/internal/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = addActorToDB(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Актер успешно добавлен")
}

func UpdateActor(w http.ResponseWriter, r *http.Request) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateActorInDB(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Информация об актере успешно обновлена")
}

func DeleteActor(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр id из URL
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию для удаления актера из базы данных по ID
	err = deleteActorFromDB(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Информация об актере успешно удалена")
}

func addActorToDB(actor models.Actor) error {
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

func updateActorInDB(actor models.Actor) error {
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

func deleteActorFromDB(actorID uint) error {
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
