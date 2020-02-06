package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Format_IsJson_valid_Json(t *testing.T) {
	var data = "{}"
	assert.True(t, IsJson(data), "Unable detect json: " + data)
}


func Test_Format_IsJson_invalid_Json(t *testing.T) {
	var data = "{"
	assert.False(t, IsJson(data), "Detected json, but data broken " + data)
}


func Test_Format_IsJson_empty_data(t *testing.T) {
	var data = ""
	assert.False(t, IsJson(data), "Detected json, but data empty")
}
