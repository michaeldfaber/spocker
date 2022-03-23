package types

type Endpoint struct {
	HttpVerb string      `json:"httpVerb"`
	Name     string      `json:"name"`
	Path     string      `json:"path"`
	Response interface{} `json:"response"`
}
