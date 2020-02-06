package core

import "net/http"

type Response struct {
	StatusCode int
	Header http.Header
	Body string
}

type ClientRequest struct {
	Request http.Request
	Response Response
}

