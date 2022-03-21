package main

type CreateEndpointRequest struct {
	HttpVerb             string      `json:"httpVerb`
	Endpoint             string      `json:"endpoint"`
	ExpectedJsonResponse interface{} `json:"expectedJsonResponse"`
}

type DeleteEndpointRequest struct {
	HttpVerb string `json:"httpVerb`
	Endpoint string `json:"endpoint"`
}
