package main

type CreateEndpointRequest struct {
	HttpVerb             string      `json:"httpVerb"`
	Endpoint             string      `json:"endpoint"`
	ExpectedJsonResponse interface{} `json:"expectedJsonResponse"`
}

type DeleteEndpointRequest struct {
	HttpVerb string `json:"httpVerb"`
	Endpoint string `json:"endpoint"`
}

type Endpoint struct {
	HttpVerb string      `json:"httpVerb"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Response interface{} `json:"response"`
}
