package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", test).Methods("POST")
	router.HandleFunc("/Testing", testing).Methods("GET")
	log.Fatal(http.ListenAndServe(":5001", router))
}

func testing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{ "Testing": "Hello World!" }`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{}`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
