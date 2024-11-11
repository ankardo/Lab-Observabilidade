package main

import (
	"net/http"
	"os"

	"github.com/ankardo/Lab-Observabilidade/configs"
	"github.com/ankardo/Lab-Observabilidade/service-b/internal/handlers"
	"github.com/ankardo/Lab-Observabilidade/service-b/internal/repositories"
	"github.com/ankardo/Lab-Observabilidade/service-b/internal/usecases"
	"github.com/ankardo/Lab-Observabilidade/tracing"
	"go.opentelemetry.io/otel"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("SERVICE_B_PORT")
	if port == "" {
		port = cfg.ServiceBPort
	}
	cleanUptracing := tracing.InitTracer("service-b")
	defer cleanUptracing()
	tracer := otel.Tracer("service-b-tracer")
	locationRepo := repositories.NewLocationRepository(http.DefaultClient)
	weatherRepo := repositories.NewWeatherRepository(http.DefaultClient)
	weatherUsecase := usecases.NewGetWeatherUseCase(locationRepo, weatherRepo)
	weatherHandler := handlers.NewWeatherHandler(weatherUsecase, tracer)
	http.HandleFunc("/weather", weatherHandler.GetWeather)
	http.ListenAndServe(":"+port, nil)
}
