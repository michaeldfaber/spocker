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
	router.HandleFunc("/hello", hello).Methods("POST")
	router.HandleFunc("/test", test).Methods("GET")
	log.Fatal(http.ListenAndServe(":5001", router))
}
func test(w http.ResponseWriter, r *http.Request) { // GET
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{}`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
func hello(w http.ResponseWriter, r *http.Request) { // POST
	w.Header().Set("Content-Type", "application/json")
	var res interface{}
	resJson := `{}`
	json.Unmarshal([]byte(resJson), &res)
	json.NewEncoder(w).Encode(res)
}
