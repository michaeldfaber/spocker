package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	endpoints "spocker/endpoints"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/endpoints", endpoints.GetAll).Methods("GET")
	router.HandleFunc("/create", endpoints.Create).Methods("POST")
	router.HandleFunc("/delete", endpoints.Delete).Methods("POST")

	router.HandleFunc("/helloWorldTwo", helloWorldTwo).Methods("GET")
	log.Fatal(http.ListenAndServe(":5001", router))
}
func helloWorldTwo(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	res, err := endpoints.GetResponse("GET", "helloWorldTwo")
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(res)
}
