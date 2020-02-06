package processor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Processor_extractHostWithoutPort(t *testing.T) {
	var host = "mysite"
	assert.Equal(t, host, extractHost(host))
}

func Test_Processor_extractHostWithPort(t *testing.T) {
	var host = "mysite:8080"
	assert.Equal(t, "mysite", extractHost(host))

}
