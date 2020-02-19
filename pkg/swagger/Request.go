package swagger

type RequestMethod struct {
	MethodName string                        `json:"-"`
	Consumes   []string                      `json:"consumes,omitempty"`
	Parameters RequestParameterList          `json:"parameters,omitempty"`
	Responses  map[int]SwaggerMethodResponse `json:"responses"`
}

type RequestParameterList []MethodGetParameter


type MethodGetParameter struct {
	In string `json:"in"`
	Name string `json:"name"`
	Required bool `json:"required"`
	Schema interface{} `json:"schema,omitempty"`
	Type string `json:"type,omitempty"`
}


func NewRequestMethod() RequestMethod {
	var method = RequestMethod{}
	method.Responses = make(map[int]SwaggerMethodResponse)
	return method
}


func NewMethodParameter() MethodGetParameter {
	var param = MethodGetParameter{}
	param.Required = true
	param.Schema = nil
	return param
}


func (req *RequestParameterList) Merge(paramList RequestParameterList) {
	if len(*req) == 0 {
		for _, param := range paramList {
			req.Append(param)
		}
		return
	}
	for _, param := range paramList {
		if req.IndexOf(param.Name, param.In) == -1 {
			req.Append(param)
			continue
		}
	}
	for index := range *req {
		var param = (*req)[index]
		if paramList.IndexOf(param.Name, param.In) == -1 {
			(*req)[index].Required = false
		}
	}
}


func (req *RequestParameterList) Append(param MethodGetParameter) {
	*req = append(*req, param)
}


func (req *RequestParameterList) IndexOf (paramName, in string) int {
	for index, v := range *req {
		if (paramName == v.Name) && (in == v.In) {
			return index;
		}
	}
	return -1
}



