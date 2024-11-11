package usecases

import (
	"context"
	"errors"
	"regexp"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/domain"
	"github.com/ankardo/Lab-Observabilidade/service-b/internal/repositories"
)

type GetWeatherUseCase struct {
	LocationRepo repositories.LocationRepository
	WeatherRepo  repositories.WeatherRepository
}

func NewGetWeatherUseCase(locRepo repositories.LocationRepository, weathRepo repositories.WeatherRepository) *GetWeatherUseCase {
	return &GetWeatherUseCase{LocationRepo: locRepo, WeatherRepo: weathRepo}
}

func (uc *GetWeatherUseCase) Execute(cep string, ctx context.Context) (*domain.Weather, error) {
	if err := ValidateZipcode(cep); err != nil {
		return nil, err
	}

	location, err := uc.LocationRepo.GetLocationByCEP(cep)
	if err != nil {
		return nil, errors.New("cannot find zipcode")
	}

	weather, err := uc.WeatherRepo.GetWeatherByLocation(location)
	if err != nil {
		return nil, errors.New("could not retrieve weather information")
	}

	return domain.NewWeather(location.City, weather.TempC), nil
}

func ValidateZipcode(cep string) error {
	matched, err := regexp.MatchString("^[0-9]{8}$", cep)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("invalid zipcode")
	}
	return nil
}
