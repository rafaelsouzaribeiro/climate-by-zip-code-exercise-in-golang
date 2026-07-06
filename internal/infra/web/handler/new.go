package handler

import (
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/ports"
)

type ClimateHandler struct {
	usecase ports.CLimate
}

func NewClimateHandler(climateUseCase ports.CLimate) *ClimateHandler {
	return &ClimateHandler{
		usecase: climateUseCase,
	}
}
