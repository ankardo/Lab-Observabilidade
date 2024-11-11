package usecases

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/ankardo/Lab-Observabilidade/service-a/internal/services"
)

type SendZipcode struct {
	zipcodeService *services.ZipcodeService
}

func NewSendZipcode(zipcodeService *services.ZipcodeService) *SendZipcode {
	return &SendZipcode{
		zipcodeService: zipcodeService,
	}
}

func (s *SendZipcode) Execute(cep string) (string, error) {
	if err := ValidateZipcode(cep); err != nil {
		return "", err
	}

	response, err := s.zipcodeService.Send(cep)
	if err != nil {
		return "", fmt.Errorf("failed to make request to Service B: %v", err)
	}
	return response, nil
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
