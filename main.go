package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rpccode/pg-api-task/db"
	"github.com/rpccode/pg-api-task/models"
	"github.com/rpccode/pg-api-task/routes"
)

func main() {

	db.DBConnect()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
