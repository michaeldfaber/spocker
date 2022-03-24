package types

type Endpoint struct {
	Id       string      `json:"id"`
	HttpVerb string      `json:"httpVerb"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Response interface{} `json:"response"`
}

type CreateEndpoint struct {
	HttpVerb string      `json:"httpVerb"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Response interface{} `json:"response"`
}
