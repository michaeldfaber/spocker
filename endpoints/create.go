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

	// create endpoint object
	var endpoint types.Endpoint
	endpoint.HttpVerb = createEndpointRequest.HttpVerb
	endpoint.Name = strings.Replace(createEndpointRequest.Endpoint, "/", "_", -1)
	endpoint.Path = createEndpointRequest.Endpoint
	endpoint.Response = createEndpointRequest.ExpectedJsonResponse

	// get expectedJsonResponse as string
	expectedJsonResponse, err := json.Marshal(createEndpointRequest.ExpectedJsonResponse)
	if err != nil {
		return
	}

	// execute bash script
	cmd := exec.Command(
		"./scripts/create-endpoint.sh",
		endpoint.HttpVerb,
		endpoint.Name,
		endpoint.Path,
		string(expectedJsonResponse))
	_, err = cmd.Output()
	if err != nil {
		return
	}

	// save to database
	err = CreateDocument(endpoint)
	if err != nil {
		return
	}
}

func CreateDocument(endpoint types.Endpoint) error {
	mongoClient, err := mongodb.New()
	if err != nil {
		return err
	}

	err = mongoClient.Create(endpoint)
	if err != nil {
		return err
	}

	mongoClient.Disconnect()
	return nil
}
