package dto

type TempResponse struct {
	LastUpdatedEpoch int     `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
}
