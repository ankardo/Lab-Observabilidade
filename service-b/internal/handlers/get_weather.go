package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ankardo/Lab-Observabilidade/service-b/internal/usecases"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

	_, locationSpan := h.Tracer.Start(ctx, "get-location-info")
	defer locationSpan.End()

	cep := r.URL.Query().Get("zipcode")
	locationSpan.SetAttributes(attribute.String("zipcode", cep))

	_, weatherSpan := h.Tracer.Start(ctx, "get-weather-info")
	defer weatherSpan.End()

	weather, err := h.GetWeatherUseCase.Execute(cep, ctx)
	if err != nil {
		weatherSpan.RecordError(err)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
