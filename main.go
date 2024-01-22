package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlgorithmChopda/Expense-Manager-Api-Go/routes"
)

func main() {
	router := routes.Route()

	fmt.Println("Expense Manger Server started")
	err := http.ListenAndServe(":8000", router)
	log.Fatal("something went wrong:", err)
}
