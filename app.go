package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//handler
func home(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(responseWriter, request)
		return
	}
	responseWriter.Write([]byte("Hello SE 08"))
}
func snippetCreate(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		responseWriter.Header().Set("Allow", http.MethodPost)
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		responseWriter.Write([]byte("Method not Allowed"))
		return
	}
	responseWriter.Write([]byte("Default snippet"))
}
func snippet(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(responseWriter, request)
		return
	}
	fmt.Fprintf(responseWriter, "Display a secific snippet with Id %d...", id)
}
