package swagger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Swagger_Empty_Data(t *testing.T) {
	var proc = NewSwaggerProcessor()

	var expected = `{
		  "swagger": "2.0",
		  "info": {
			"title": "rzreversescheme - swagger scheme generator",
			"version": "0.0.1a"
		  },
		  "paths": {}
		}`
	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request")
}


func Test_Swagger_Add_Request_Get_Without_Parameters(t *testing.T) {
	var clientReq = createGetClientReq("/get/one")
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
		  "swagger": "2.0",
		  "info": {
			"title": "rzreversescheme - swagger scheme generator",
			"version": "0.0.1a"
		  },
		  "paths": {
			"/get/one": {
			  "get": {
				"responses": {
				  "200": {
					"description": "Status 200 OK",
					"schema": {
					  "type": "object",
					  "properties": {
						"id": {
						  "type": "integer"
						},
						"username": {
						  "type": "string"
						}
					  }
					}
				  }
				}
			  }
			}
		  }
		}`

	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request")
}


func Test_Swagger_Add_Request_Get_With_QueryParameters_String(t *testing.T) {
	var clientReq = createGetClientReq("/get/query?var1=str1")
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/get/query": {
		  "get": {
			"parameters": [
			  {
				"name": "var1",
				"in": "query",
				"required": true,
				"type": "string"
			  }
			  ],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "integer"
					},
					"username": {
					  "type": "string"
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	}`

	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request with query string parameters")
}


func Test_Swagger_Add_Request_Get_With_QueryParameters_Int(t *testing.T) {
	var clientReq = createGetClientReq("/get/query?var1=12")
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/get/query": {
		  "get": {
			"parameters": [
			  {
				"name": "var1",
				"in": "query",
				"required": true,
				"type": "integer"
			  }
			  ],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "integer"
					},
					"username": {
					  "type": "string"
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	}`
	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request with query int parameters")
}


func Test_Swagger_Add_Request_Get_With_QueryParameters_float(t *testing.T) {
	var clientReq = createGetClientReq("/get/query?var1=55.1")
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/get/query": {
		  "get": {
			"parameters": [
			  {
				"name": "var1",
				"in": "query",
				"required": true,
				"type": "number"
			  }
			  ],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "integer"
					},
					"username": {
					  "type": "string"
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	}`
	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request with query float(number) parameters")
}


func Test_Swagger_Add_Request_Post_Without_Parameters(t *testing.T) {
	var clientReq = createPostClientReq("/post/one", make(map[string]string))
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
		  "swagger": "2.0",
		  "info": {
			"title": "rzreversescheme - swagger scheme generator",
			"version": "0.0.1a"
		  },
		  "paths": {
			"/post/one": {
			  "post": {
				"responses": {
				  "200": {
					"description": "Status 200 OK",
					"schema": {
					  "type": "object",
					  "properties": {
						"id": {
						  "type": "integer"
						},
						"username": {
						  "type": "string"
						}
					  }
					}
				  }
				}
			  }
			}
		  }
		}`

	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for post request")
}


func Test_Swagger_Add_Request_Post_With_FormParameters_String(t *testing.T) {
	var params = make(map[string]string)
	params["var1"] = "str1"
	var clientReq = createPostClientReq("/post/one", params)
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/post/one": {
		  "post": {
			"parameters": [
			  {
				"name": "var1",
				"in": "formData",
				"required": true,
				"type": "string"
			  }
			  ],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "integer"
					},
					"username": {
					  "type": "string"
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	}`

	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request with query string parameters")
}


func Test_Swagger_Add_Request_Post_With_FormParameters_Int(t *testing.T) {
	var params = make(map[string]string)
	params["var1"] = "57"
	var clientReq = createPostClientReq("/post/one", params)
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/post/one": {
		  "post": {
			"parameters": [
			  {
				"name": "var1",
				"in": "formData",
				"required": true,
				"type": "integer"
			  }
			  ],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "integer"
					},
					"username": {
					  "type": "string"
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	}`

	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request with query string parameters")
}


func Test_Swagger_Add_Request_Post_With_FormParameters_float(t *testing.T) {
	var params = make(map[string]string)
	params["var1"] = "126.3"
	var clientReq = createPostClientReq("/post/one", params)
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/post/one": {
		  "post": {
			"parameters": [
			  {
				"name": "var1",
				"in": "formData",
				"required": true,
				"type": "number"
			  }
			  ],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "integer"
					},
					"username": {
					  "type": "string"
					}
				  }
				}
			  }
			}
		  }
		}
	  }
	}`

	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for get request with query string parameters")
}



func Test_Swagger_Add_Request_Post_With_JsonBody(t *testing.T) {
	var body = `{
       "jsonrpc":"2.0",
       "id":"550e8400-e29b-41d4-a716-446655440000",
       "method":"certs.activate",
       "params":{
          "certificate_code": "GIFT190333-999666",
          "phone": "+70000000001",
          "activation_code": "112233",
          "customer_id": 777,
          "email": "customer@lamoda.ru"
       }
   }`
	var clientReq = createPostClientReqJson("/post/body/json", body)
	var proc = NewSwaggerProcessor()
	proc.Process(clientReq)

	var expected = `{
	  "swagger": "2.0",
	  "info": {
		"title": "rzreversescheme - swagger scheme generator",
		"version": "0.0.1a"
	  },
	  "paths": {
		"/post/body/json": {
		  "post": {
			"parameters": [
			  {
				"in": "body",
				"name": "body",
				"required": true,
				"schema": {
				  "type": "object",
				  "properties": {
					"id": {
					  "type": "string"
					},
					"jsonrpc": {
					  "type": "string"
					},
					"method": {
					  "type": "string"
					},
					"params": {
					  "type": "object",
					  "properties": {
						"activation_code": {
						  "type": "string"
						},
						"certificate_code": {
						  "type": "string"
						},
						"customer_id": {
						  "type": "integer"
						},
						"email": {
						  "type": "string"
						},
						"phone": {
						  "type": "string"
						}
					  }
					}
				  }
				}
			  }
			],
			"responses": {
			  "200": {
				"description": "Status 200 OK",
				"schema": {
					"type": "object",
					"properties": {
					  "id": {
						"type": "integer"
					  },
					  "username": {
						"type": "string"
					  }
					
				  }
				}
			  }
			}
		  }
		}
	  }
	}`
	assert.JSONEq(t, expected, proc.GetScheme(), "Bad scheme for json body")
}



//func Test_Swagger_Add_Request_GetPost_OnePath_Without_Parameters(t *testing.T) {
//	var params = make(map[string]string)
//	var clientReq = createPostClientReq("/get/one", params)
//	clientReq.Response.Body = `{"id":21}`
//	var proc = NewSwaggerProcessor()
//	proc.Process(clientReq)
//
//	clientReq = createGetClientReq("/get/one")
//	clientReq.Response.Body = `{"id":22}`
//	proc.Process(clientReq)
//	var expScheme = `{"paths":{"/get/one":{"get":{"responses":{"200":{"content":{"application/json":{"type":"object","properties":{"id":{"type":"integer"}}}}}}},"post":{"responses":{"200":{"content":{"application/json":{"type":"object","properties":{"id":{"type":"integer"}}}}}}}}}}`
//	assert.JSONEq(t, expScheme, proc.GetScheme(), "Bad scheme for 2 request with one path: get&post")
//}
//
//
//
//
//
//
//func Test_Swagger_Add_Request_Get_With_QueryParameter_Array(t *testing.T) {
//
//	var clientReq = createGetClientReq("/get/query?var[0]=123&var[2]=param2&var[3]=345")
//
//
//	//u, _ := url.Parse(clientReq.Request.URL.String())
//	//queryParams := u.Query()
//	//for name, val := range queryParams {
//	//	fmt.Printf("Params: %v => %v\n", name, val[0])
//	//}
//
//
//	var proc = NewSwaggerProcessor()
//	proc.Process(clientReq)
//
//	//var expScheme = `{"paths":{"/get/one":{"get":{"parameters":{},"responses":{"200":{}}}}}}`
//	//var actualScheme = proc.GetScheme()
//	//if expScheme != actualScheme {
//	//	t.Errorf("Bad scheme struct: " + actualScheme)
//	//}
//}
//
//
////https://swagger.io/docs/specification/basic-structure/
//
