package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	manage_endpoints "spocker/manage_endpoints"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/endpoints", manage_endpoints.GetAll).Methods("GET")
	router.HandleFunc("/create", manage_endpoints.Create).Methods("POST")
	router.HandleFunc("/delete", manage_endpoints.Delete).Methods("POST")

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
