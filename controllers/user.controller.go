package controllers

import (
	"fmt"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "User Login")
}
