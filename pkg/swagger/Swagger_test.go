package swagger

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"rzreversescheme/pkg/core"
	"testing"
)



func Test_Swagger_CreateScheme_Object_IntValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"id":{"type":"integer"}}}`
	var proc = NewSwaggerProcessor()
	var clientReq = createGetClientReq("/get/request")
	clientReq.Response.Body = `{"id": 40}`

	scheme := proc.createSchemeFromResponse(clientReq)
	schemeStr, _ := json.Marshal(scheme)


	assert.JSONEq(t, expected, string(schemeStr))
}


func Test_Swagger_hasPath_Exists(t *testing.T) {

	var proc = NewSwaggerProcessor()
	var requestUri string = "/my/test/path"

	var reqMethod = NewRequestMethod()
	var methodName = http.MethodGet
	proc.doc.Path[requestUri] = make(map[string]RequestMethod)
	proc.doc.Path[requestUri][methodName] = reqMethod


	if !proc.hasPath(requestUri) {
		t.Errorf("Path must be exists")
	}
}


func Test_Swagger_hasPath_NotExists(t *testing.T) {
	var proc = NewSwaggerProcessor()
	var requestUri string = "/my/test/path"

	var reqMethod = NewRequestMethod()
	var methodName = "GET"
	proc.doc.Path[requestUri] = make(map[string]RequestMethod)
	proc.doc.Path[requestUri][methodName] = reqMethod

	if proc.hasPath("other/path") {
		t.Errorf("Path must be not exists")
	}
}


func createGetClientReq(url string) core.ClientRequest {
	var clientReq = core.ClientRequest{}
	req, _ := http.NewRequest("GET", url, nil)

	req.RequestURI = "/get/one"
	req.Method = http.MethodGet
	clientReq.Request = *req
	clientReq.Response.Body = `{"id": 40, "username": "Vasya"}`
	clientReq.Response.StatusCode = http.StatusOK
	return clientReq
}


func createPostClientReq(reqUrl string, params map[string]string) core.ClientRequest {

	formData := url.Values{}
	for name, value := range params {
		formData.Set(name, value)
	}
	var clientReq = core.ClientRequest{}

	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBufferString(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.RequestURI = "/get/one"
	req.Method = http.MethodPost
	clientReq.Request = *req
	clientReq.Response.Body = `{"id": 40, "username": "Vasya"}`
	clientReq.Response.StatusCode = http.StatusOK
	return clientReq
}


func createPostClientReqJson(reqUrl string, body string) core.ClientRequest {

	var clientReq = core.ClientRequest{}

	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBufferString(body))
	urlObj, _ := url.Parse(reqUrl)
	req.RequestURI = urlObj.RequestURI()
	req.Method = http.MethodPost
	clientReq.Request = *req
	clientReq.Response.Body = `{"id": 40, "username": "Vasya"}`
	clientReq.Response.StatusCode = http.StatusOK
	return clientReq
}

