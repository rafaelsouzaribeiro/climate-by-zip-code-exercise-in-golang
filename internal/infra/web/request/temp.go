package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rafaelsouzaribeiro/climate-by-zip-code-exercise-in-golang/internal/dto"
	"github.com/spf13/viper"
)

func (r *Request) GetTemp(city string) (*dto.TempResponse, error) {

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", viper.GetString("KEY"), city)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in API response: %s", resp.Status)
	}

	var tempResponse dto.TempResponse
	err = json.NewDecoder(resp.Body).Decode(&tempResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return &tempResponse, nil
}
