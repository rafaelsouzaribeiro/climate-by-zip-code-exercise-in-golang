package usecase

import "github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/ports"

type ClimateUseCase struct {
	ports ports.CLimate
}

func NewClimateUseCase(ports ports.CLimate) *ClimateUseCase {
	return &ClimateUseCase{
		ports: ports,
	}
}
