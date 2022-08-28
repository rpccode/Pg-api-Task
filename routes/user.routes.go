package routes

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/rpccode/pg-api-task/db"
	"github.com/rpccode/pg-api-task/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	// getUsers :=
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	db.DB.Find(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte("Usuario No Encontrado"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createUser := db.DB.Create(&user)
	err := createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getUsersHnadler"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.Find(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte("Usuario No Encontrado"))
		return
	}
	/// no elimina el dato solo lo cambia de estado
	db.DB.Delete(&user)

	///elimina el dato de la base de datos
	// db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuario Eliminado Correctamente"))

}
