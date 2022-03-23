package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	types "spocker/types"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	endpointsJson, err := ioutil.ReadFile("endpoints.json")
	if err != nil {
		return
	}

	var endpoints []types.Endpoint
	json.Unmarshal(endpointsJson, &endpoints)
	endpoints = endpoints[0 : len(endpoints)-1]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoints)
}
