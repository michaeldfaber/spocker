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

	router.HandleFunc("/helloWorldFive", helloWorldFive).Methods("GET")
	router.HandleFunc("/helloWorldFour", helloWorldFour).Methods("GET")
	router.HandleFunc("/helloWorld", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":5001", router))
}
func helloWorld(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{"Hello":"World"}`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
func helloWorldFour(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{"Hello":"World`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
func helloWorldFive(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{"Hello":"World`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
