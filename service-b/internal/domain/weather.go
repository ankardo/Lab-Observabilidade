package domain

import "math"

type Weather struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeather(city string, tempC float64) *Weather {
	return &Weather{
		City:  city,
		TempC: truncateFloat(tempC, 1),
		TempF: truncateFloat((tempC*1.8 + 32), 1),
		TempK: truncateFloat((tempC + 273.15), 1),
	}
}

func truncateFloat(value float64, precision int) float64 {
	pow := math.Pow(10, float64(precision))
	return float64(int(value*pow)) / pow
}
