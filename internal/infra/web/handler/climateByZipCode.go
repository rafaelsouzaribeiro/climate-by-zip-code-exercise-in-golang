package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrZipCodeNotFound = errors.New("can not find zipcode")
	ErrTempNotFound    = errors.New("can not find temp")
)

func (h *ClimateHandler) GetClimateByZipCode(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")

	viaCep, err := h.usecase.GetViaCep(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if viaCep.Cep == "" {
		http.Error(w, ErrZipCodeNotFound.Error(), http.StatusNotFound)
		return
	}

	temp, err := h.usecase.GetTemp(viaCep.Localidade)
	if err != nil {
		http.Error(w, ErrTempNotFound.Error(), http.StatusNotFound)
		return
	}

	tempF := (temp.Current.TempC * 1.8) + 32
	tempK := temp.Current.TempC + 273.15

	response := map[string]float64{
		"temp_C": temp.Current.TempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}
