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

func Delete(w http.ResponseWriter, r *http.Request) {
	// read request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// unmarshal into deleteEndpointRequest
	var deleteEndpointRequest types.DeleteEndpointRequest
	json.Unmarshal(body, &deleteEndpointRequest)

	name := strings.Replace(deleteEndpointRequest.Endpoint, "/", "_", -1)

	// execute bash script
	cmd := exec.Command(
		"./scripts/delete-endpoint.sh",
		deleteEndpointRequest.HttpVerb,
		name,
		deleteEndpointRequest.Endpoint)
	_, err = cmd.Output()
	if err != nil {
		return
	}

	// delete from database
	err = DeleteDocument(deleteEndpointRequest.Id)
	if err != nil {
		return
	}
}

func DeleteDocument(id string) error {
	mongoClient, err := mongodb.New()
	if err != nil {
		return err
	}

	err = mongoClient.Delete(id)
	if err != nil {
		return err
	}

	mongoClient.Disconnect()
	return nil
}
