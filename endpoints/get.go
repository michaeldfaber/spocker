package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	mongodb "spocker/mongodb"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	mongoClient, err := mongodb.New()
	if err != nil {
		mongoClient.Disconnect()
		return
	}

	res, err := mongoClient.GetAll()
	fmt.Printf("%v\n", res)
	if err != nil {
		mongoClient.Disconnect()
		return
	}

	mongoClient.Disconnect()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetResponse(httpVerb string, name string) (interface{}, error) {
	mongoClient, err := mongodb.New()
	if err != nil {
		mongoClient.Disconnect()
		return nil, err
	}

	response, err := mongoClient.GetResponse(httpVerb, name)
	if err != nil {
		mongoClient.Disconnect()
		return nil, err
	}

	mongoClient.Disconnect()
	return response, nil
}
