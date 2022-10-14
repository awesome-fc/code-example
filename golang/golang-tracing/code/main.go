package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	fmt.Println("hello world")
	endpoint := fctx.Tracing.JaegerEndpoint
	OpenTracingSpanContext := fctx.Tracing.OpenTracingSpanContext
	if len(endpoint) == 0 {
		return "", fmt.Errorf("invalid jaeger endpoint")
	}
	// New tracer
	tracer, closer := NewJaegerTracer("FCTracer", endpoint)
	defer closer.Close()

	// retrieve spanContext
	spanContext, err := jaeger.ContextFromString(OpenTracingSpanContext)
	if err != nil {
		return "", fmt.Errorf("OpenTracingSpanContext: %s, error: %v", fctx.Tracing.OpenTracingSpanContext, err)
	}

	// span start/finish
	startMySpan(spanContext, tracer)

	return fmt.Sprintf("hello world! 你好，%s!", event.Name), nil
}

func NewJaegerTracer(service, endpoint string) (opentracing.Tracer, io.Closer) {
	sender := transport.NewHTTPTransport(endpoint)
	tracer, closer := jaeger.NewTracer(service,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender))
	return tracer, closer
}

func startMySpan(context jaeger.SpanContext, tracer opentracing.Tracer) {
	parentSpan := tracer.StartSpan("MyFCSpan", opentracing.ChildOf(context))
	defer parentSpan.Finish()
	parentSpan.SetOperationName("fc-operation")
	parentSpan.SetTag("version", "fc-v1")
	time.Sleep(150 * time.Millisecond)
	// child span start/finish
	childSpan := tracer.StartSpan("fc-operation-child", opentracing.ChildOf(parentSpan.Context()))
	defer childSpan.Finish()
	time.Sleep(100 * time.Millisecond)
	childSpan.LogFields(
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 100))
}

func main() {
	fc.Start(HandleRequest)
}
