package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func get_endpoints(w http.ResponseWriter, r *http.Request) {
	endpointsJson, err := ioutil.ReadFile("endpoints.json")
	if err != nil {
		return
	}

	var endpoints []Endpoint
	json.Unmarshal(endpointsJson, &endpoints)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoints)
}

func create_endpoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var createEndpointRequest CreateEndpointRequest
	json.Unmarshal(body, &createEndpointRequest)

	expectedJsonResponse, err := json.Marshal(createEndpointRequest.ExpectedJsonResponse)
	if err != nil {
		return
	}

	cmd := exec.Command(
		"./scripts/create-endpoint.sh",
		createEndpointRequest.HttpVerb,
		createEndpointRequest.Endpoint,
		string(expectedJsonResponse))
	_, err = cmd.Output()
	if err != nil {
		return
	}
}

func delete_endpoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var deleteEndpointRequest DeleteEndpointRequest
	json.Unmarshal(body, &deleteEndpointRequest)

	cmd := exec.Command(
		"./scripts/delete-endpoint.sh",
		deleteEndpointRequest.HttpVerb,
		deleteEndpointRequest.Endpoint)
	_, err = cmd.Output()
	if err != nil {
		return
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", create_endpoint).Methods("POST")
	router.HandleFunc("/delete", delete_endpoint).Methods("POST")
	router.HandleFunc("/endpoints", get_endpoints).Methods("GET")

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
