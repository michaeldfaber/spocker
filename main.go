package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	handleRequests()
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

	log.Fatal(http.ListenAndServe(":5001", router))
}
