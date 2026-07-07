package request

import (
	"testing"

	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/configs"
	"github.com/stretchr/testify/assert"
)

func TestViaCep_Success(t *testing.T) {

	request := NewRequest()
	configs.LoadConfig("../../../../cmd")
	temp, err := request.GetViaCep("18199899")
	assert.NoError(t, err)
	assert.NotNil(t, temp)
	assert.Equal(t, "18199-899", temp.Cep)

}

func TestViaCep_NotFound(t *testing.T) {
	request := NewRequest()
	configs.LoadConfig("../../../../cmd")
	temp, err := request.GetViaCep("")
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid zipcode")
	assert.Nil(t, temp)

}
