package swagger

type SwaggerMethodResponse struct {
	Description string	`json:"description"`
	Schema JsonSchemaProperty `json:"schema"`
}


func NewSwaggerMethodResponse() SwaggerMethodResponse {
	var resp = SwaggerMethodResponse{}
	return resp
}


func (resp *SwaggerMethodResponse) Merge(element SwaggerMethodResponse) {
	resp.Schema.Merge(element.Schema)
}

