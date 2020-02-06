package swagger

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

type swaggerType struct {

}

const SWAGGER_TYPE_UNKNOWN = -1
const SWAGGER_TYPE_NULL = 10
const SWAGGER_TYPE_BOOL = 21
const SWAGGER_TYPE_INT = 31
const SWAGGER_TYPE_NUMBER = 32
const SWAGGER_TYPE_STRING = 41
const SWAGGER_TYPE_ARRAY = 51
const SWAGGER_TYPE_OBJECT = 61

const SWAGGER_PARAM_IN_QUERY = "query"
const SWAGGER_PARAM_IN_PATH = "path"
const SWAGGER_PARAM_IN_FORMDATA = "formData"
const SWAGGER_PARAM_IN_BODY = "body"

func (st swaggerType) getName(typeId int) string {
	var name string
	var nameList = map[int]string{
		SWAGGER_TYPE_NULL: "null",
		SWAGGER_TYPE_BOOL: "boolean",
		SWAGGER_TYPE_INT: "integer",
		SWAGGER_TYPE_NUMBER: "number",
		SWAGGER_TYPE_STRING: "string",
		SWAGGER_TYPE_ARRAY: "array",
		SWAGGER_TYPE_OBJECT: "object",
	}
	name, ok := nameList[typeId]
	if !ok {
		name = fmt.Sprintf("undefined type id: %v", typeId)
	}
	return name
}



func (st swaggerType) getType(variable interface{}) int {
	if variable == nil {
		return SWAGGER_TYPE_NULL
	}
	var typeId = SWAGGER_TYPE_UNKNOWN
	switch reflect.TypeOf(variable).Kind() {
	case reflect.Bool:
		typeId = SWAGGER_TYPE_BOOL
	case reflect.String:
		typeId = SWAGGER_TYPE_STRING
	case reflect.Array, reflect.Slice:
		typeId = SWAGGER_TYPE_ARRAY
	case reflect.Map:
		typeId = SWAGGER_TYPE_OBJECT
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			typeId = SWAGGER_TYPE_INT
	case reflect.Float32, reflect.Float64:
		typeId = SWAGGER_TYPE_NUMBER
		if variable == float64(int64(variable.(float64))) {
			typeId = SWAGGER_TYPE_INT
		}

	}
	return typeId
}


func (st swaggerType) isCompotible(mainType, cmpType int) bool {
	if mainType/10 != cmpType/10 {
		return false
	}
	if cmpType % 10 > mainType % 10 {
		return true
	}
	return false
}


func (st swaggerType) Compact(typeList []int) []int {
	var result []int
	var cache = make(map[int]int)
	sort.Ints(typeList)
	for _, v := range typeList {
		cache[int(v/10)] = v
	}
	for _, v := range typeList {
		if cache[int(v/10)] ==  v {
			result = append(result, v)
		}
	}
	return result
}

func (st swaggerType) transformToNative(value string) interface{} {
	intValue, intError := strconv.Atoi(value)
	if (intError == nil) && (strconv.Itoa(intValue) == value) {
		return intValue
	}

	floatValue, floatError := strconv.ParseFloat(value, 32)
	if (floatError == nil) && (strconv.FormatFloat(floatValue, 'f',1, 64) == value) {
		return floatValue
	}
	return value
}
