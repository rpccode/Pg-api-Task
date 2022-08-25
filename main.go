package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rpccode/pg-api-task/db"

	"github.com/rpccode/pg-api-task/router"
)


func main() {
	
	db.DBConnect()
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
