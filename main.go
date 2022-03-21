package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	handleRequests()
}

func create_endpoint(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./scripts/create-endpoint.sh", "GET", "create_test", "{}")
	_, err := cmd.Output()
	fmt.Printf("%v", err) // TODO handle err
}

func delete_endpoint(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./scripts/delete-endpoint.sh", "GET", "create_test")
	_, err := cmd.Output()
	fmt.Printf("%v", err) // TODO handle err
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", create_endpoint).Methods("POST")
	router.HandleFunc("/delete", delete_endpoint).Methods("POST")

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
