package controllers

import (
	"encoding/json"
	"filmLibrary/internal/models"
	"filmLibrary/internal/service"
	"filmLibrary/pkg/logging"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// @Summary Удалить актера
// @Description Удаляет актера из базы данных по его ID
// @Tags actors
// @Param id path string true "ID актера"
// @Success 200 {string} string "Информация об актере успешно удалена"
// @Failure 400 {string} string "Ошибка декодирования JSON"
// @Failure 500 {string} string "Ошибка удаления актера в БД"
// @Router /actor/delete-actor/{id} [delete]

func AddActor(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		logger.Error("Ошибка декодирования JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AddActorToDB(actor, logger)
	if err != nil {
		logger.Error("Ошибка добавления актера в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Актер успешно добавлен")
	fmt.Fprintf(w, "Актер успешно добавлен")
}

// UpdateActor обновляет информацию об актере в базе данных.
// @Summary Обновить актера
// @Description Обновляет информацию об актере в базе данных
// @Tags actors
// @Accept json
// @Produce json
// @Param actor body Actor true "Данные актера"
// @Success 200 {string} string "Информация об актере успешно обновлена"
// @Failure 400 {string} string "Ошибка декодирования JSON"
// @Failure 500 {string} string "Ошибка обновления актера в БД"
// @Router /actor/update-actor [post]

func UpdateActor(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	var actor models.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		logger.Error("Ошибка декодирования JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateActorInDB(actor, logger)
	if err != nil {
		logger.Error("Ошибка обновления актера в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Информация об актере успешно обновлена")
	fmt.Fprintf(w, "Информация об актере успешно обновлена")
}

// DeleteActor удаляет актера из базы данных по его ID.
// @Summary Удалить актера
// @Description Удаляет актера из базы данных по его ID
// @Tags actors
// @Param id path string true "ID актера"
// @Success 200 {string} string "Информация об актере успешно удалена"
// @Failure 400 {string} string "Ошибка декодирования JSON"
// @Failure 500 {string} string "Ошибка удаления актера в БД"
// @Router /actor/delete-actor/{id} [delete]

func DeleteActor(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		logger.Error("Ошибка декодирования JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteActorFromDB(uint(id), logger)
	if err != nil {
		logger.Error("Ошибка удаления актера в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Info("Информация об актере успешно удалена")
	fmt.Fprintf(w, "Информация об актере успешно удалена")
}
