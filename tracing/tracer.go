package tracing

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer(serviceName string) func() {
	exporterURL := "otel-collector:4318"
	ctx := context.Background()
	exporter, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(exporterURL),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create the OTLP trace exporter: %v", err)
	}

	bsp := trace.NewBatchSpanProcessor(exporter)
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	)

	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(bsp),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shut down tracer provider: %v", err)
		}
	}
}
