package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type WeatherHandler struct {
	GetWeatherUseCase *usecases.GetWeatherUseCase
	Tracer            trace.Tracer
}

func NewWeatherHandler(uc *usecases.GetWeatherUseCase, tracer trace.Tracer) *WeatherHandler {
	return &WeatherHandler{GetWeatherUseCase: uc, Tracer: tracer}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), carrier)
	cep := r.URL.Query().Get("zipcode")
	weather, err := h.GetWeatherUseCase.Execute(cep, h.Tracer, ctx)
	if err != nil {
		switch err.Error() {
		case "invalid zipcode":
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		case "cannot find zipcode":
			http.Error(w, err.Error(), http.StatusNotFound)
		case "could not retrieve weather information":
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
