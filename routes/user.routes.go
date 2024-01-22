package routes

import (
	"github.com/AlgorithmChopda/Expense-Manager-Api-Go/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.UserLogin)
}
