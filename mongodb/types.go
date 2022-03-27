package mongodb

type Document struct {
	Id       string                 `bson:"_id" json:"id,omitempty"`
	HttpVerb string                 `json:"httpVerb"`
	Name     string                 `json:"name"`
	Path     string                 `json:"path"`
	Response map[string]interface{} `json:"response"`
}
