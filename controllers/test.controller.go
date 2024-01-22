package controllers

import (
	"fmt"
	"net/http"
)

func Test(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "Test")
}
