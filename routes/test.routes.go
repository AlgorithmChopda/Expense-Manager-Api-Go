package routes

import (
	"github.com/AlgorithmChopda/Expense-Manager-Api-Go/controllers"
	"github.com/gorilla/mux"
)

func TestRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.Test).Methods("GET")
}
