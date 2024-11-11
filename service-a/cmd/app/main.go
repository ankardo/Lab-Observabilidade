package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ankardo/Lab-Observabilidade/configs"
	"github.com/ankardo/Lab-Observabilidade/service-a/internal/handlers"
	"github.com/ankardo/Lab-Observabilidade/service-a/internal/services"
	"github.com/ankardo/Lab-Observabilidade/service-a/internal/usecases"
	"github.com/ankardo/Lab-Observabilidade/tracing"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("SERVICE_A_PORT")
	if port == "" {
		port = cfg.ServiceAPort
	}
	cleanUptracing := tracing.InitTracer("service-a")
	defer cleanUptracing()

	zipcodeService := services.NewZipcodeService(cfg.ServiceBURL, http.DefaultClient)
	zipcodeUsecase := usecases.NewSendZipcode(zipcodeService)
	zipcodeHandler := handlers.NewZipcodeHandler(zipcodeUsecase)

	http.HandleFunc("/", zipcodeHandler.SendZipcode)

	log.Printf("Service A is running on port %s...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
