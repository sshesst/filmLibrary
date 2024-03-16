package handler

//import (
//	"encoding/json"
//	"filmLibrary/internal/controllers"
//	"filmLibrary/internal/models"
//	"fmt"
//	"net/http"
//)
//
//func AddActorHandler(w http.ResponseWriter, r *http.Request) {
//	var actor models.Actor
//	err := json.NewDecoder(r.Body).Decode(&actor)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err = controllers.AddActor(actor)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Fprintf(w, "Актер успешно добавлен")
//}
//
//func UpdateActorHandler(w http.ResponseWriter, r *http.Request) {
//	var actor models.Actor
//	err := json.NewDecoder(r.Body).Decode(&actor)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err = controllers.UpdateActor(actor)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Fprintf(w, "Информация об актере успешно обновлена")
//}
//
//func DeleteActorHandler(w http.ResponseWriter, r *http.Request) {
//	var actor models.Actor
//	err := json.NewDecoder(r.Body).Decode(&actor)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err = controllers.DeleteActor(actor.ID)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	fmt.Fprintf(w, "Информация об актере успешно удалена")
//}
