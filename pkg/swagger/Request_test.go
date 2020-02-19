package swagger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_Request_Parameter_append_single(t *testing.T) {
	var paramList = RequestParameterList{}
	var param = NewMethodParameter()
	param.Name = "test01"
	paramList.Append(param)

	assert.Equal(t, 1, len(paramList), "Unable Append RequestParameterList")
	assert.Equal(t, param.Name, paramList[0].Name, "Unable Append RequestParameterList - unknown value added")
}


func Test_Request_Parameter_append_multiple(t *testing.T) {
	var paramList = RequestParameterList{}
	var testList []MethodGetParameter
	for i := 0; i < 10; i++ {
		var param = NewMethodParameter()
		param.Name = fmt.Sprintf("test%d", i)
		testList = append(testList, param)
		paramList.Append(param)
	}

	assert.Equal(t, len(testList), len(paramList), "After add elements not all added")
	for index, param := range testList {
		assert.Equal(t, param.Name, paramList[index].Name, "Unable Append RequestParameterList - unknown value added")
	}
}


func Test_Request_Parameter_IndexOf_Found(t *testing.T) {
	var paramList = RequestParameterList{}
	var testList []MethodGetParameter
	for i := 0; i < 10; i++ {
		var param = NewMethodParameter()
		param.Name = fmt.Sprintf("test%d", i)
		param.In = SWAGGER_PARAM_IN_QUERY
		testList = append(testList, param)
		paramList.Append(param)
	}

	assert.Equal(t, len(testList), len(paramList), "After add elements not all added")
	for index, param := range testList {
		assert.Equal(t, index, paramList.IndexOf(param.Name, param.In), "Unable find parameter using RequestParameterList.IndexOf")
	}
}


func Test_Request_Parameter_IndexOf_NotFound(t *testing.T) {
	var paramList = RequestParameterList{}
	var testList []MethodGetParameter
	for i := 0; i < 10; i++ {
		var param = NewMethodParameter()
		param.Name = fmt.Sprintf("test%d", i)
		param.In = SWAGGER_PARAM_IN_QUERY
		testList = append(testList, param)
		paramList.Append(param)
	}

	assert.Equal(t, len(testList), len(paramList), "After add elements not all added")
	assert.Equal(t, -1, paramList.IndexOf("noname", "not_in"), "Element found, but not exists")
}


func Test_Request_Parameter_merge_into_empty(t *testing.T) {
	var paramList = RequestParameterList{}
	var inject = RequestParameterList{}
	inject = paramListAddParam(inject, "test", SWAGGER_PARAM_IN_QUERY)

	paramList.Merge(inject)
	assert.GreaterOrEqual(t, 0, paramList.IndexOf("test", SWAGGER_PARAM_IN_QUERY), "Parameter not merged")
}


func Test_Request_Parameter_merge_new_into_not_empty(t *testing.T) {
	var paramList = RequestParameterList{}
	var inject = RequestParameterList{}
	inject = paramListAddParam(inject, "test", SWAGGER_PARAM_IN_QUERY)
	paramList = paramListAddParam(paramList, "exists", SWAGGER_PARAM_IN_QUERY)

	paramList.Merge(inject)
	assert.Equal(t, 2, len(paramList), "No new parameters added after merge")
	assert.GreaterOrEqual(t, paramList.IndexOf("test", SWAGGER_PARAM_IN_QUERY),0, "Parameter not merged")
}


func Test_Request_Parameter_merge_exists_parameter(t *testing.T) {
	var paramList = RequestParameterList{}
	var inject = RequestParameterList{}
	inject = paramListAddParam(inject, "test", SWAGGER_PARAM_IN_QUERY)
	paramList = paramListAddParam(paramList, "other!value", SWAGGER_PARAM_IN_QUERY)
	paramList = paramListAddParam(paramList, "test", SWAGGER_PARAM_IN_QUERY)

	paramList.Merge(inject)
	assert.GreaterOrEqual(t, paramList.IndexOf("test", SWAGGER_PARAM_IN_QUERY),0,"Parameter lost after merge")
	assert.Equal(t, 2, len(paramList), "Parameter duplicated")
}


func Test_Request_Parameter_merge_not_required_parameter(t *testing.T) {
	var paramList = RequestParameterList{}
	var inject = RequestParameterList{}
	inject = paramListAddParam(inject, "test", SWAGGER_PARAM_IN_QUERY)
	paramList = paramListAddParam(paramList, "not_required", SWAGGER_PARAM_IN_QUERY)
	paramList = paramListAddParam(paramList, "test", SWAGGER_PARAM_IN_QUERY)

	paramList.Merge(inject)
	var index = paramList.IndexOf("not_required", SWAGGER_PARAM_IN_QUERY)
	assert.GreaterOrEqual(t, index,0,"Parameter lost after merge")
	assert.Equal(t, false, paramList[index].Required,"When parameter not in new request - parameter mark required=false")
}


func paramListAddParam(paramList RequestParameterList, name, in string) RequestParameterList {
	var param = NewMethodParameter()
	param.Name = name
	param.In = in
	paramList.Append(param)
	return paramList
}
