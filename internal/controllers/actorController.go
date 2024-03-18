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
