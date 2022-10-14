package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get FcContext")
	}
	fmt.Println("hello world")

	spanCtx, endpoint, err := getFcTracingInfo(fctx)
	if err != nil {
		return "", fmt.Errorf("failed to getFcTracingInfo, error: %v", err)
	}

	tp, err := NewTracerProvider(endpoint)
	if err != nil {
		return "", fmt.Errorf("OpenTracingJaegerEndpoint: %s, error: %v", fctx.Tracing.JaegerEndpoint, err)
	}
	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	if err != nil {
		return "", fmt.Errorf("failed to getFcSpanCtx, error: %v", err)
	}

	startMySpan(trace.ContextWithSpanContext(ctx, spanCtx))
	return fmt.Sprintf("hello world! 你好，%s!", event.Name), nil
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.

func NewTracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithSyncer(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("FCTracer"),
			attribute.String("environment", "production"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

func getFcTracingInfo(fctx *fccontext.FcContext) (trace.SpanContext, string, error) {
	spanContext := trace.SpanContext{}
	endpoint := fctx.Tracing.JaegerEndpoint
	OpenTracingSpanContext := fctx.Tracing.OpenTracingSpanContext
	if len(endpoint) == 0 {
		return spanContext, endpoint, fmt.Errorf("invalid jaeger endpoint")
	}
	spanContextSlice := strings.Split(OpenTracingSpanContext, ":")

	//Fill the high bits of the TraceID
	tid, err := trace.TraceIDFromHex("0000000000000000" + spanContextSlice[0])
	if err != nil {
		return spanContext, endpoint, err
	}
	fid := trace.FlagsSampled
	spanContext = spanContext.WithTraceID(tid).WithTraceFlags(fid).WithRemote(true)
	return spanContext, endpoint, nil
}

func startMySpan(ctx context.Context) {
	tr := otel.Tracer("fc-Trace")
	ctx, parentSpan := tr.Start(ctx, "fc-operation")
	defer parentSpan.End()
	parentSpan.SetAttributes(attribute.Key("version").String("fc-v1"))
	time.Sleep(150 * time.Millisecond)
	child(ctx)
}

func child(ctx context.Context) {
	// Use the global TracerProvider.
	tr := otel.Tracer("fc-Trace")
	_, childSpan := tr.Start(ctx, "fc-operation-child")
	defer childSpan.End()
	time.Sleep(100 * time.Millisecond)
	childSpan.AddEvent("timeout")
}

func main() {
	fc.Start(HandleRequest)
}
