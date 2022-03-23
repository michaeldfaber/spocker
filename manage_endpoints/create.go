package manage_endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	mongodb "spocker/mongodb"
	types "spocker/types"
)

func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var createEndpointRequest types.CreateEndpointRequest
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

	err = CreateDocument(createEndpointRequest)
	if err != nil {
		return
	}
}

func CreateDocument(createEndpointRequest types.CreateEndpointRequest) error {
	mongoClient, err := mongodb.New()
	if err != nil {
		fmt.Printf("4 %v\n", err)
		return err
	}

	err = mongoClient.Create(createEndpointRequest)
	if err != nil {
		fmt.Printf("5 %v\n", err)
		return err
	}

	mongoClient.Disconnect()
	return nil
}
