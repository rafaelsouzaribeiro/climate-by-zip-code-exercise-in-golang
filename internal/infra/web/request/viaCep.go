package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/dto"
)

var ErrInvalidZipCode = errors.New("invalid zipcode")

var cepRegex = regexp.MustCompile(`^\d{8}$`)

func (r *Request) GetViaCep(cep string) (*dto.ViaCepResponse, error) {
	if !cepRegex.MatchString(cep) {
		return nil, ErrInvalidZipCode
	}

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in API response: %s", resp.Status)
	}

	var viaCepResponse dto.ViaCepResponse
	err = json.NewDecoder(resp.Body).Decode(&viaCepResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return &viaCepResponse, nil
}
