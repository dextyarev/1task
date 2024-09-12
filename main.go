package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string

type requestBody struct {
	Message string `json:"message"`
}

func SetMessage(w http.ResponseWriter, r *http.Request) {
	var req requestBody
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message = req.Message
	fmt.Fprintf(w, "message: %s", message)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s", message)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sendMessage", SetMessage).Methods("POST")
	r.HandleFunc("/getMessage", GetMessage).Methods("GET")
	http.ListenAndServe(":8080", r)
}
