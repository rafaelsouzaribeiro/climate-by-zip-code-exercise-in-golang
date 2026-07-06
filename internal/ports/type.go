package ports

import "github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/dto"

type CLimate interface {
	GetViaCep(cep string) (*dto.ViaCepResponse, error)
	GetTemp(city string) (*dto.TempResponse, error)
}
