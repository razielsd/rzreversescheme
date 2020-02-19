package swagger

import (
	"encoding/json"
)

type JsonSchemaProperty struct {
	Type interface{}                         `json:"type"`
	TypeList map[int]int                     `json:"-"`
	Items interface{}                        `json:"items,,omitempty"`
	Properties map[string]JsonSchemaProperty `json:"properties,,omitempty"`
}


func NewJsonSchemaProperty() JsonSchemaProperty {
	var schema = JsonSchemaProperty{}
	schema.Type = "object"
	schema.Properties = make(map[string]JsonSchemaProperty)
	schema.TypeList = make(map[int]int)
	return schema
}


type JsonSchema []string
//https://cswr.github.io/JsonSchema/spec/multiple_types/

func (proc JsonSchema) CreateFromString(value string) JsonSchemaProperty{
	var data  interface{}
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		panic("SwaggerProcessor.createSchemeFromResponse required valid json (check on upper level)")
	}
	return proc.Create(data)
}


func (proc JsonSchema) Create(value interface{}) JsonSchemaProperty {
	var property JsonSchemaProperty
	typeTransformer := swaggerType{}
	switch typeTransformer.getType(value) {
	case SWAGGER_TYPE_OBJECT:
		property = proc.createObjectFromValue(value)
	case SWAGGER_TYPE_ARRAY:
		property = proc.createArrayFromValue(value)
	default:
		property = proc.createSchemePropertyScalar(value)
	}
	return property
}



func (proc JsonSchema) createArrayFromValue(value interface{}) JsonSchemaProperty {
	var typeId = SWAGGER_TYPE_NULL
	prop := NewJsonSchemaProperty()
	typeTransformer := swaggerType{}
	prop.Type = typeTransformer.getName(SWAGGER_TYPE_ARRAY)
	var items = NewJsonSchemaProperty()
	ir := value.([]interface{})
	for _, v := range ir {
		tid := typeTransformer.getType(v)
		if tid > typeId {
			typeId = tid
		}
		switch typeId {
		case SWAGGER_TYPE_OBJECT:
			items = proc.createObjectFromValue(v)
		case SWAGGER_TYPE_ARRAY:
			items = proc.createArrayFromValue(v)
		default:
			items.addType(typeId)
		}
	}
	prop.Items = items
	return prop
}


func (proc JsonSchema) createObjectFromValue(value interface{}) JsonSchemaProperty {
	var property = NewJsonSchemaProperty()
	ir := value.(map[string]interface{})
	for k, v := range ir {
		property.Properties[k] = proc.Create(v)
	}
	return property
}


func (proc JsonSchema) createSchemePropertyScalar(value interface{}) JsonSchemaProperty {
	property := NewJsonSchemaProperty()
	typeTransformer := swaggerType{}
	typeId := typeTransformer.getType(value)
	property.Type = typeTransformer.getName(typeId)
	return property
}


func (proc JsonSchema) toString(schema interface{}) string {
	result, _ := json.Marshal(schema)
	return string(result)
}


func (prop *JsonSchemaProperty) addType(typeId int) {
	_, exist := (*prop).TypeList[typeId]
	if !exist {
		typeTransformer := swaggerType{}
		(*prop).TypeList[typeId] = typeId
		var ids []int
		for _, v := range (*prop).TypeList {
			ids = append(ids, v)
		}
		var compact = typeTransformer.Compact(ids)

		if len(compact) == 1 {
			(*prop).Type = typeTransformer.getName(compact[0])
		} else {
			var typeNameList []string
			for _, id := range compact {
				typeNameList = append(typeNameList, typeTransformer.getName(id))
			}
			(*prop).Type = typeNameList
		}
	}
}


