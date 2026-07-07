package request

import (
	"net/url"
	"testing"

	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/configs"
	"github.com/stretchr/testify/assert"
)

func TestGetTemp_Success(t *testing.T) {
	city := url.QueryEscape("São Paulo")

	request := NewRequest()
	configs.LoadConfig("../../../../cmd")
	temp, err := request.GetTemp(city)
	assert.NoError(t, err)
	assert.NotNil(t, temp)

}
