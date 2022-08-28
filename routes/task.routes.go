package routes

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/rpccode/pg-api-task/db"
	"github.com/rpccode/pg-api-task/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	// getUsers :=
	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	db.DB.Find(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)//404

		w.Write([]byte("Tarea No Encontrado"))
		return
	}
	json.NewEncoder(w).Encode(&task)

}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createTask := db.DB.Create(&task)
	err := createTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)//404
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getUsersHnadler"))
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	db.DB.Find(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)///404

		w.Write([]byte("Tarea No Encontrado"))
		return
	}
	/// no elimina el dato solo lo cambia de estado
	db.DB.Delete(&task)

	///elimina el dato de la base de datos
	// db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusNoContent)///204
	w.Write([]byte("Tarea Eliminada Correctamente"))

}
