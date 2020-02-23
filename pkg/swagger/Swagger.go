package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"rzreversescheme/pkg/core"
	"rzreversescheme/pkg/format"
	"strings"
	"sync"
)

type SwaggerDoc struct {
	Swagger string `json:"swagger"`
	Info map[string]string `json:"info"`
	Path map[string]map[string]RequestMethod `json:"paths"`
}




type SwaggerProcessor struct {
	doc SwaggerDoc
	lock sync.RWMutex
}


func NewSwaggerProcessor() SwaggerProcessor {
	var proc = SwaggerProcessor{}
	proc.doc.Swagger = "2.0"
	proc.doc.Path = make(map[string]map[string]RequestMethod)
	proc.doc.Info = make(map[string]string)
	proc.doc.Info["title"] = "rzreversescheme - swagger scheme generator"
	proc.doc.Info["version"] = "0.0.1a"
	return proc
}


func (proc SwaggerProcessor) GetScheme() string {
	scheme, _ := json.Marshal(proc.doc)
	return string(scheme)
}


func (proc SwaggerProcessor) Process(clientReq core.ClientRequest) {
	if proc.hasSame(clientReq) {
		proc.update(clientReq)
	} else {
		proc.add(clientReq)
	}
}


func (proc SwaggerProcessor) hasSame(clientReq core.ClientRequest) bool {
	var path = proc.extractPath(clientReq)
	if !proc.hasPath(path) {
		return false
	}
	return false
}


func (proc SwaggerProcessor) hasPath(path string) bool {
	if _, ok := proc.doc.Path[path]; ok {
		return true
	}
	return false
}


func (proc SwaggerProcessor) add(clientReq core.ClientRequest) {
	var requestUri = proc.extractPath(clientReq)
	var paramList = RequestParameterList{}
	var reqMethod = NewRequestMethod()
	reqMethod.MethodName = strings.ToLower(clientReq.Request.Method)

	resp := proc.processResponse(clientReq)
	switch clientReq.Request.Method {
		case http.MethodGet:
			paramList = proc.processQueryParameters(clientReq)
		case http.MethodPost:
			paramList = proc.processPostParameters(&reqMethod, clientReq)
	}
	reqMethod.Parameters.Merge(paramList)
	reqMethod.Responses[clientReq.Response.StatusCode] = resp
	_, hasUri := proc.doc.Path[requestUri]
	if !hasUri {
		proc.doc.Path[requestUri] = make(map[string]RequestMethod)
	}
	proc.doc.Path[requestUri][reqMethod.MethodName] = reqMethod
}


func (proc SwaggerProcessor) processQueryParameters(clientReq core.ClientRequest) RequestParameterList {
	var paramList = RequestParameterList{}
	u, _ := url.Parse(clientReq.Request.URL.String())
	queryParams := u.Query()
	typeTransformer := swaggerType{}
	for name, value := range queryParams {
		var param = NewMethodParameter()
		param.Name = name
		param.In = SWAGGER_PARAM_IN_QUERY
		param.Required = true
		param.Type = typeTransformer.getName(typeTransformer.getType(typeTransformer.transformToNative(value[0])))
		paramList.Append(param)
	}
	return paramList
}


func (proc SwaggerProcessor) processPostParameters(reqMethod *RequestMethod, clientReq core.ClientRequest) RequestParameterList {
	var paramList = RequestParameterList{}
	var contentType = strings.ToLower(clientReq.Request.Header.Get("content-type"))

	if strings.HasPrefix(contentType, "multipart/") {
		contentType = "multipart/"
	 } else if strings.Contains(contentType, "form-urlencoded") {
		contentType = "form-urlencoded"
	}

	switch contentType {
		case "multipart/":
			reqMethod.Consumes[0] = "multipart/form-data"
			return proc.processPostMultipartForm(clientReq)
		case "form-urlencoded":
			return proc.processPostForm(clientReq)
		default:
			var body, _ = ioutil.ReadAll(clientReq.Request.Body)
			var bodyStr = string(body)

			if format.IsJson(bodyStr) {
				paramList = proc.processPostJsonBody(bodyStr, clientReq)
			}
	}
	return paramList
}


func (proc SwaggerProcessor) processPostMultipartForm(clientReq core.ClientRequest) RequestParameterList {
	var paramList = RequestParameterList{}
	err := clientReq.Request.ParseMultipartForm(10e6)//will be get from config
	if err != nil {
		return paramList //can't and to nothing, may bad way, but better no idea
	}
	//@todo type "file" not supported
	typeTransformer := swaggerType{}
	for name, value := range clientReq.Request.Form {
		var param = NewMethodParameter()
		param.Name = name
		param.In = SWAGGER_PARAM_IN_FORMDATA
		param.Required = true
		param.Type = typeTransformer.getName(typeTransformer.getType(typeTransformer.transformToNative(value[0])))
		paramList.Append(param)
	}
	return paramList
}


func (proc SwaggerProcessor) processPostForm(clientReq core.ClientRequest) RequestParameterList {
	var paramList = RequestParameterList{}
	err := clientReq.Request.ParseForm()
	if err != nil {
		return paramList //can't and to nothing, may be bad way, but better no idea
	}
	typeTransformer := swaggerType{}
	for name, value := range clientReq.Request.Form {
		var param = NewMethodParameter()
		param.Name = name
		param.In = SWAGGER_PARAM_IN_FORMDATA
		param.Required = true
		param.Type = typeTransformer.getName(typeTransformer.getType(typeTransformer.transformToNative(value[0])))
		paramList.Append(param)
	}
	return paramList
}


func (proc SwaggerProcessor) processPostJsonBody(body string, clientReq core.ClientRequest) RequestParameterList {
	var paramList = RequestParameterList{}
	var schemeBuilder = JsonSchema{}
	prop := schemeBuilder.CreateFromString(body)

	var param = NewMethodParameter()
	param.In = SWAGGER_PARAM_IN_BODY
	param.Name = SWAGGER_PARAM_IN_BODY
	param.Required = true
	param.Schema = prop
	paramList.Append(param)

	return paramList
}


func (proc SwaggerProcessor) processResponse(clientReq core.ClientRequest) SwaggerMethodResponse{
	methodResp := NewSwaggerMethodResponse()
	methodResp.Description = fmt.Sprintf(
		"Status %d %s",
		clientReq.Response.StatusCode,
		http.StatusText(clientReq.Response.StatusCode))
	methodResp.Schema = proc.createSchemeFromResponse(clientReq)
	return methodResp
}


func (proc SwaggerProcessor) createSchemeFromResponse(clientReq core.ClientRequest) JsonSchemaProperty {
	var resp  interface{}
	var data = []byte(clientReq.Response.Body)
	err := json.Unmarshal(data, &resp)
	if err != nil {
		panic("SwaggerProcessor.createSchemeFromResponse required valid json (check on upper level)")
	}
	var schemeBuilder = JsonSchema{}
	schema := schemeBuilder.Create(resp)
	return schema
}


func (proc SwaggerProcessor) update(clientReq core.ClientRequest) {
	var requestUri = proc.extractPath(clientReq)
	var paramList = RequestParameterList{}

	var methodName = strings.ToUpper(clientReq.Request.Method)
	var reqMethod = proc.doc.Path[requestUri][methodName]
	reqMethod.MethodName = methodName

	if proc.isSuccessStatus(clientReq.Response.StatusCode) {
		switch methodName {
		case http.MethodGet:
			paramList = proc.processQueryParameters(clientReq)
		case http.MethodPost:
			paramList = proc.processPostParameters(&reqMethod, clientReq)
		}
		reqMethod.Parameters.Merge(paramList)
	}

	resp := proc.processResponse(clientReq)
	_, hasResponse := reqMethod.Responses[clientReq.Response.StatusCode]
	if (!hasResponse) {
		reqMethod.Responses[clientReq.Response.StatusCode] = resp
	} else {
		proc.lock.Lock()
		respContainer := reqMethod.Responses[clientReq.Response.StatusCode]
		respContainer.Merge(resp)
		reqMethod.Responses[clientReq.Response.StatusCode] = respContainer
		proc.lock.Unlock()
	}

	proc.doc.Path[requestUri][methodName] = reqMethod
}


func (proc SwaggerProcessor) extractPath(clientReq core.ClientRequest) string {
	return clientReq.Request.URL.Path
}


func (proc SwaggerProcessor) isSuccessStatus(statusCode int) bool {
	return statusCode < 300
}
