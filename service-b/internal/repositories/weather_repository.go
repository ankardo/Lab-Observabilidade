package repositories

import (
	"net/http"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/api"
	"github.com/ankardo/Lab-Observabilidade/service-b/internal/domain"
)

type WeatherRepository interface {
	GetWeatherByLocation(location *domain.Location) (*domain.Weather, error)
}

type weatherRepo struct {
	client *http.Client
}

func NewWeatherRepository(client *http.Client) WeatherRepository {
	return &weatherRepo{client: client}
}

func (r *weatherRepo) GetWeatherByLocation(location *domain.Location) (*domain.Weather, error) {
	weather, err := api.FetchWeather(location, r.client)
	return weather, err
}
