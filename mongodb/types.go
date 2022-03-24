package mongodb

type Filter struct {
	httpVerb string
	name     string
}

type Document struct {
	HttpVerb string                 `json:"httpVerb"`
	Name     string                 `json:"name"`
	Path     string                 `json:"path"`
	Response map[string]interface{} `json:"response"`
}
