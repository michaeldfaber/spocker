package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"

	mongodb "spocker/mongodb"
	types "spocker/types"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// read request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// unmarshal into deleteEndpointRequest
	var deleteEndpointRequest types.DeleteEndpointRequest
	json.Unmarshal(body, &deleteEndpointRequest)

	// delete from database
	endpoint, err := DeleteDocument(deleteEndpointRequest.Id)
	if err != nil {
		return
	}

	// execute bash script
	cmd := exec.Command(
		"./scripts/delete-endpoint.sh",
		endpoint.HttpVerb,
		endpoint.Name,
		endpoint.Path)
	_, err = cmd.Output()
	if err != nil {
		return
	}
}

func DeleteDocument(id string) (types.Endpoint, error) {
	var endpoint types.Endpoint
	mongoClient, err := mongodb.New()
	if err != nil {
		return endpoint, err
	}

	endpoint, err = mongoClient.Delete(id)
	if err != nil {
		return endpoint, err
	}

	mongoClient.Disconnect()
	return endpoint, nil
}
