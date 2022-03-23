package manage_endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"

	types "spocker/types"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var deleteEndpointRequest types.DeleteEndpointRequest
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
