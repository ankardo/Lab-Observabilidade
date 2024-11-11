package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/domain"
)

func FetchWeather(location *domain.Location, client *http.Client) (*domain.Weather, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		err := fmt.Errorf("missing API key")
		return nil, err
	}

	requestURL := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, url.QueryEscape(location.City))
	resp, err := client.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to fetch weather data, status code: %d", resp.StatusCode)
		return nil, err
	}

	var weatherResponse struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}

	return domain.NewWeather(location.City, weatherResponse.Current.TempC), nil
}
