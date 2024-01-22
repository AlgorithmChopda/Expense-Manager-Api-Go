package controllers

import (
	"fmt"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "Welcome to Expense Manager API made in Golang")
}
