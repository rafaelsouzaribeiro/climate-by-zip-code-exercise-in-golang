package ports

import "github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/dto"

type CLimate interface {
	GetViaCep(cep string) (*dto.ViaCepResponseOutput, error)
	GetTemp(city string) (*dto.TempResponseOutput, error)
}
