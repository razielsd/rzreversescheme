package swagger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_JsonSchema_CreateScheme_Object_IntValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"id":{"type":"integer"}}}`
	var proc = JsonSchema{}
	var jsonStr = `{"id": 40}`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Object_FloatValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"id":{"type":"number"}}}`
	var proc = JsonSchema{}
	var jsonStr = `{"id": 40.1}`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Object_StringValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"username":{"type":"string"}}}`
	var proc = JsonSchema{}
	var jsonStr = `{"username": "Vasya"}`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Object_BoolValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"flag":{"type":"boolean"}}}`
	var proc = JsonSchema{}
	var jsonStr = `{"flag": true}`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Object_NullValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"nullValue":{"type":"null"}}}`
	var proc = JsonSchema{}
	var jsonStr = `{"nullValue": null}`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Object_ObjectValue(t *testing.T) {
	var expected = `{"type":"object","properties":{"info":{"type":"object","properties":{"username":{"type":"string"}}}}}`
	var proc = JsonSchema{}
	var jsonStr = `{"info": {"username":"Vasya"}}`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Array_IntValue(t *testing.T) {
	var expected = `{"type":"array","items":{"type":"integer"}}`
	var proc = JsonSchema{}
	var jsonStr = `[1,3,5]`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Array_FloatValue(t *testing.T) {
	var expected = `{"type":"array","items":{"type":"number"}}`
	var proc = JsonSchema{}
	var jsonStr = `[1.1,3.4,5.7]`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Array_mixedFloatIntValue(t *testing.T) {
	var expected = `{"type":"array","items":{"type":"number"}}`
	var proc = JsonSchema{}
	var jsonStr = `[1,3.4,5]`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_Array_Object(t *testing.T) {
	var expected = `{"type":"array","items":{"type":"object","properties":{"id":{"type":"integer"}}}}`
	var proc = JsonSchema{}
	var jsonStr = `[{"id":1},{"id":4},{"id":7}]`
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_Swagger_CreateScheme_JsonRpc(t *testing.T) {
	var expected = `{"type":"object","properties":{"id":{"type":"string"},"jsonrpc":{"type":"string"},"method":{"type":"string"},"params":{"type":"object","properties":{"activation_code":{"type":"string"},"certificate_code":{"type":"string"},"customer_id":{"type":"integer"},"email":{"type":"string"},"phone":{"type":"string"}}}}}`
	var proc = JsonSchema{}
	var jsonStr = `{
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
	prop := proc.CreateFromString(jsonStr)
	scheme := proc.toString(prop)

	assert.JSONEq(t, expected, scheme, "Bad scheme for: " + jsonStr)
}


func Test_JsonSchemaProperty_AddType_int_float(t *testing.T) {
	var prop = NewJsonSchemaProperty()
	prop.addType(SWAGGER_TYPE_INT)
	prop.addType(SWAGGER_TYPE_NUMBER)
	assert.Equal(t,"number", prop.Type, "int and number must be cast to number")
}


func Test_JsonSchemaProperty_AddType_int_float_string(t *testing.T) {
	var expected = []string{"number", "string"}
	var prop = NewJsonSchemaProperty()
	prop.addType(SWAGGER_TYPE_INT)
	prop.addType(SWAGGER_TYPE_NUMBER)
	prop.addType(SWAGGER_TYPE_STRING)
	assert.Equal(t,expected, prop.Type, "int and number must be cast to number")
}


func Test_JsonSchemaProperty_Merge_Two_DifferentStruct(t *testing.T) {
	var proc = JsonSchema{}
	var expected = `{
	   "type":"object",
	   "properties":{
		  "id":{
			 "type":"integer"
		  },
		  "username":{
			 "type":"string"
		  }
	   }
	}`
	var jsonStr = `{"username": "Vasya"}`
	var injectStr = `{"username": "Vasya", "id": 45}`
	prop := proc.CreateFromString(jsonStr)
	inject := proc.CreateFromString(injectStr)
	prop.Merge(inject)
	assert.JSONEq(t, expected, proc.toString(prop))
}
