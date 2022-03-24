package endpoints

import (
	"net/http"

	mongodb "spocker/mongodb"
)

func GetAll(w http.ResponseWriter, r *http.Request) {}

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
