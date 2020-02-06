package swagger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swaggerType_nil(t *testing.T) {
	var st = swaggerType{}

	if st.getType(nil) != SWAGGER_TYPE_NULL {
		t.Errorf("Unable convert nil to NULL")
	}
}


func Test_swaggerType_bool(t *testing.T) {
	var st = swaggerType{}
	var data bool = true
	if st.getType(data) != SWAGGER_TYPE_BOOL {
		t.Errorf("Unable detect boolean type")
	}
}


func Test_swaggerType_int64(t *testing.T) {
	var st = swaggerType{}
	var data int64 = 55
	if st.getType(data) != SWAGGER_TYPE_INT {
		t.Errorf("Unable detect int type")
	}
}


func Test_swaggerType_float(t *testing.T) {
	var st = swaggerType{}
	var data float64 = 40.5
	if st.getType(data) != SWAGGER_TYPE_NUMBER {
		t.Errorf("Unable detect number type")
	}
}


func Test_swaggerType_string(t *testing.T) {
	var st = swaggerType{}
	var data string = "my litte text"
	if st.getType(data) != SWAGGER_TYPE_STRING {
		t.Errorf("Unable detect string type")
	}
}


func Test_swaggerType_array(t *testing.T) {
	var st = swaggerType{}
	var data = []int{1,2,35,}
	if st.getType(data) != SWAGGER_TYPE_ARRAY {
		t.Errorf("Unable detect array type")
	}
}


func Test_swaggerType_object(t *testing.T) {
	var st = swaggerType{}
	var data = make(map[string]int)
	if st.getType(data) != SWAGGER_TYPE_OBJECT {
		t.Errorf("Unable detect object type")
	}
}


func Test_swaggerType_cmp_int_string(t *testing.T) {
	var st = swaggerType{}
	if st.isCompotible(SWAGGER_TYPE_INT, SWAGGER_TYPE_STRING) {
		t.Errorf("int and string must be incompotible")
	}
}


func Test_swaggerType_cmp_int_float(t *testing.T) {
	var st = swaggerType{}
	if !st.isCompotible(SWAGGER_TYPE_INT, SWAGGER_TYPE_NUMBER) {
		t.Errorf("int and float must be compotible")
	}
}


func Test_swaggerType_cmp_float_int(t *testing.T) {
	var st = swaggerType{}
	if st.isCompotible(SWAGGER_TYPE_NUMBER, SWAGGER_TYPE_INT) {
		t.Errorf("float and int must be incompotible")
	}
}


func Test_swaggerType_cmp_int_nil(t *testing.T) {
	var st = swaggerType{}
	if st.isCompotible(SWAGGER_TYPE_INT, SWAGGER_TYPE_NULL) {
		t.Errorf("int and null must be incompotible")
	}
}


func Test_swaggerType_Compact_int_float(t *testing.T) {
	var st = swaggerType{}
	var expected []int = []int{SWAGGER_TYPE_NUMBER}
	var typeList []int = []int{SWAGGER_TYPE_INT, SWAGGER_TYPE_NUMBER}

	assert.Equal(t, expected, st.Compact(typeList))
}

func Test_swaggerType_Compact_int_float_string(t *testing.T) {
	var st = swaggerType{}
	var expected []int = []int{SWAGGER_TYPE_NUMBER, SWAGGER_TYPE_STRING}
	var typeList []int = []int{SWAGGER_TYPE_INT, SWAGGER_TYPE_NUMBER, SWAGGER_TYPE_STRING}

	assert.Equal(t, expected, st.Compact(typeList))
}


func Test_swaggerType_transformToNative_int(t *testing.T) {
	var transformer = swaggerType{}
	var value = "555"

	assert.Equal(t, SWAGGER_TYPE_INT, transformer.getType(transformer.transformToNative(value)))
}


func Test_swaggerType_transformToNative_float(t *testing.T) {
	var transformer = swaggerType{}
	var value = "555.3"

	assert.Equal(t, SWAGGER_TYPE_NUMBER, transformer.getType(transformer.transformToNative(value)))
}
