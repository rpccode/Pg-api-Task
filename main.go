package main

import (
	"net/http"

	"github.com/rpccode/pg-api-task/db"

	"github.com/gorilla/mux"
	"github.com/rpccode/pg-api-task/models"
	"github.com/rpccode/pg-api-task/routes"
)

func main() {

	db.DBConnect()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	//User router
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("Put")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("Post")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Task router
	r.HandleFunc("/task", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.UpdateTaskHandler).Methods("Put")
	r.HandleFunc("/task", routes.PostTaskHandler).Methods("Post")
	r.HandleFunc("/task/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
