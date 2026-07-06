package handler

import (
	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/usecase"
)

type ClimateHandler struct {
	usecase usecase.ClimateUseCase
}

func NewClimateHandler(climateUseCase usecase.ClimateUseCase) *ClimateHandler {
	return &ClimateHandler{
		usecase: climateUseCase,
	}
}
