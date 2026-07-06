package usecase

import (
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/dto"
)

func (t *ClimateUseCase) GetTemp(city string) (*dto.TempResponseOutput, error) {
	temp, err := t.ports.GetTemp(city)

	if err != nil {
		return nil, err
	}

	return temp, nil
}
