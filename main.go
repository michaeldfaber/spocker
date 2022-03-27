package main

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	endpoints "spocker/endpoints"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/heartbeat", heartbeat).Methods("GET")
	router.HandleFunc("/endpoints", endpoints.GetAll).Methods("GET")
	router.HandleFunc("/create", endpoints.Create).Methods("POST")
	router.HandleFunc("/delete", endpoints.Delete).Methods("POST")

	router.HandleFunc("/helloWorldTwo", helloWorldTwo).Methods("GET")
	router.HandleFunc("/helloWorld", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":5001", handlers.CORS(credentials, methods, origins)(router)))
}
func heartbeat(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	json.NewEncoder(w).Encode(res)
}
func helloWorld(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	res, err := endpoints.GetResponse("GET", "helloWorld")
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(res)
}
func helloWorldTwo(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	res, err := endpoints.GetResponse("GET", "helloWorldTwo")
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(res)
}
