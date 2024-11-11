package repositories

import (
	"net/http"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/api"
	"github.com/ankardo/Lab-Observabilidade/service-b/internal/domain"
)

type LocationRepository interface {
	GetLocationByCEP(zipcode string) (*domain.Location, error)
}

type locationRepo struct {
	client *http.Client
}

func NewLocationRepository(client *http.Client) LocationRepository {
	return &locationRepo{client: client}
}

func (r *locationRepo) GetLocationByCEP(zipcode string) (*domain.Location, error) {
	location, err := api.FetchLocation(zipcode, r.client)
	return location, err
}
