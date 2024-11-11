package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/domain"
)

var viaCEPBaseURL = "https://viacep.com.br/ws/"

func FetchLocation(zipcode string, client *http.Client) (*domain.Location, error) {
	resp, err := client.Get(viaCEPBaseURL + zipcode + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location domain.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.ZipCode == "" {
		err := errors.New("invalid zipcode")
		return nil, err
	}
	return &location, nil
}
