package domain

type Location struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	IBGE         string `json:"ibge"`
	GIA          string `json:"gia"`
	DDD          string `json:"ddd"`
	SIAFI        string `json:"siafi"`
}

func NewLocation(
	zipCode, street, complement, neighborhood, city, state, ibge, gia, areaCode, siafi string,
) *Location {
	location := &Location{
		ZipCode:      zipCode,
		Street:       street,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		IBGE:         ibge,
		GIA:          gia,
		SIAFI:        siafi,
	}
	return location
}
