package cmdserver

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"rzreversescheme/pkg/processor"
	"testing"
)


func Test_handlerConfigureHost(t *testing.T) {
	r := LoadRouter()
	processor.Init()
	var body = `{
		  "host": "localtest",
		  "hostfilter": "localhost"
		}`
	req := createPostRequest("/configure/host", body)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		assert.Equal(t, http.StatusOK, w.Code)
		var expected = `{
			"message": "rule added"
		}`
		json, _ := ioutil.ReadAll(w.Body)
		assert.JSONEq(t, expected, string(json))
		assert.Equal(t, 1, processor.HostMatcherService.Count())
		return true
	})
}


func Test_handlerConfigureHost_EmptyFilter(t *testing.T) {
	r := LoadRouter()
	var body = `{
		  "host": "localtest"
		}`
	req := createPostRequest("/configure/host", body)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var expected = `{
			"error": "No active filter"
		}`
		json, _ := ioutil.ReadAll(w.Body)
		assert.JSONEq(t, expected, string(json))
		return true
	})
}


func Test_handlerConfigureHost_EmptyHost(t *testing.T) {
	r := LoadRouter()
	var body = `{
		  "hostfilter": "localhost"
		}`
	req := createPostRequest("/configure/host", body)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var expected = `{
			"error": "Empty field: host"
		}`
		json, _ := ioutil.ReadAll(w.Body)
		assert.JSONEq(t, expected, string(json))
		return true
	})
}


func createPostRequest(reqUrl string, body string) *http.Request {
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBufferString(body))
	urlObj, _ := url.Parse(reqUrl)
	req.RequestURI = urlObj.RequestURI()
	req.Method = http.MethodPost
	return req
}


func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	f(w)
}