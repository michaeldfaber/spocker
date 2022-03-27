package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	mongodb "github.com/michaeldfaber/spocker/mongodb"
	types "github.com/michaeldfaber/spocker/types"
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

	// save to database
	id, err := CreateDocument(createEndpoint)
	if err != nil {
		return
	}

	// execute bash script
	cmd := exec.Command(
		"./scripts/create-endpoint.sh",
		createEndpoint.HttpVerb,
		createEndpoint.Name,
		createEndpoint.Path,
		id)
	_, err = cmd.Output()
	if err != nil {
		return
	}
}

func CreateDocument(createEndpoint types.CreateEndpoint) (string, error) {
	var id string
	mongoClient, err := mongodb.New()
	if err != nil {
		return id, err
	}

	id, err = mongoClient.Create(createEndpoint)
	if err != nil {
		return id, err
	}

	mongoClient.Disconnect()
	return id, nil
}
