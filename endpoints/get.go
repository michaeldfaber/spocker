package endpoints

import (
	"encoding/json"
	"net/http"

	mongodb "github.com/michaeldfaber/spocker/mongodb"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	mongoClient, err := mongodb.New()
	if err != nil {
		mongoClient.Disconnect()
		return
	}

	res, err := mongoClient.GetAll()
	if err != nil {
		mongoClient.Disconnect()
		return
	}

	mongoClient.Disconnect()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetResponse(id string) (interface{}, error) {
	mongoClient, err := mongodb.New()
	if err != nil {
		mongoClient.Disconnect()
		return nil, err
	}

	response, err := mongoClient.GetResponse(id)
	if err != nil {
		mongoClient.Disconnect()
		return nil, err
	}

	mongoClient.Disconnect()
	return response, nil
}
