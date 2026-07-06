package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClimateMock struct {
	mock.Mock
}

func (m *ClimateMock) GetViaCep(cep string) (*dto.ViaCepResponseOutput, error) {
	args := m.Called(cep)
	return args.Get(0).(*dto.ViaCepResponseOutput), args.Error(1)
}

func (m *ClimateMock) GetTemp(city string) (*dto.TempResponseOutput, error) {
	args := m.Called(city)
	return args.Get(0).(*dto.TempResponseOutput), args.Error(1)
}

func TestGetClimateByZipCode_Success(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &ClimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "01001000").Return(&dto.ViaCepResponseOutput{
		Cep:        "01001000",
		Localidade: "São Paulo",
	}, nil).Once()

	mockUC.On("GetTemp", "S%C3%A3o+Paulo").Return(&dto.TempResponseOutput{
		Currents: dto.Current{
			TempC: 25.0,
		},
	}, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/01001000", nil)
	req.SetPathValue("cep", "01001000")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var body map[string]float64
	err := json.NewDecoder(rr.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, 25.0, body["temp_C"])
	assert.Equal(t, 77.0, body["temp_F"])
	assert.InDelta(t, 298.15, body["temp_K"], 0.0001)

	mockUC.AssertExpectations(t)
}

func TestGetClimateByZipCode_InvalidZipCode_Returns422(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &ClimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "123").Return(&dto.ViaCepResponseOutput{}, errors.New("invalid zipcode")).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/123", nil)
	req.SetPathValue("cep", "123")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid zipcode")

	mockUC.AssertExpectations(t)
}

func TestGetClimateByZipCode_ZipNotFound_Returns404(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &ClimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "00000000").Return(&dto.ViaCepResponseOutput{
		Cep: "",
	}, nil).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/00000000", nil)
	req.SetPathValue("cep", "00000000")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), ErrZipCodeNotFound.Error())

	mockUC.AssertExpectations(t)
}

func TestGetClimateByZipCode_TempNotFound_Returns404(t *testing.T) {
	mockUC := new(ClimateMock)
	h := &ClimateHandler{usecase: mockUC}

	mockUC.On("GetViaCep", "01001000").Return(&dto.ViaCepResponseOutput{
		Cep:        "01001000",
		Localidade: "São Paulo",
	}, nil).Once()

	mockUC.On("GetTemp", "S%C3%A3o+Paulo").Return(&dto.TempResponseOutput{}, errors.New("temp api error")).Once()

	req := httptest.NewRequest(http.MethodGet, "/climate/01001000", nil)
	req.SetPathValue("cep", "01001000")
	rr := httptest.NewRecorder()

	h.GetClimateByZipCode(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), ErrTempNotFound.Error())

	mockUC.AssertExpectations(t)
}
