package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	mongodb "spocker/mongodb"
	types "spocker/types"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// read request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// unmarshal into createEndpointRequest
	var createEndpointRequest types.CreateEndpointRequest
	json.Unmarshal(body, &createEndpointRequest)

	// createEndpoint object
	var createEndpoint types.CreateEndpoint
	createEndpoint.HttpVerb = createEndpointRequest.HttpVerb
	createEndpoint.Name = strings.Replace(createEndpointRequest.Endpoint, "/", "_", -1)
	createEndpoint.Path = createEndpointRequest.Endpoint
	createEndpoint.Response = createEndpointRequest.ExpectedJsonResponse

	// get expectedJsonResponse as string
	expectedJsonResponse, err := json.Marshal(createEndpointRequest.ExpectedJsonResponse)
	if err != nil {
		return
	}

	// execute bash script
	cmd := exec.Command(
		"./scripts/create-endpoint.sh",
		createEndpoint.HttpVerb,
		createEndpoint.Name,
		createEndpoint.Path,
		string(expectedJsonResponse))
	_, err = cmd.Output()
	if err != nil {
		return
	}

	// save to database
	err = CreateDocument(createEndpoint)
	if err != nil {
		return
	}
}

func CreateDocument(createEndpoint types.CreateEndpoint) error {
	mongoClient, err := mongodb.New()
	if err != nil {
		return err
	}

	err = mongoClient.Create(createEndpoint)
	if err != nil {
		return err
	}

	mongoClient.Disconnect()
	return nil
}
