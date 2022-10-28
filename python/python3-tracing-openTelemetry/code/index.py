# -*- coding: utf-8 -*-
import time
from opentelemetry import trace
from opentelemetry.exporter.jaeger.thrift import JaegerExporter
from opentelemetry.sdk.resources import SERVICE_NAME, Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import SimpleSpanProcessor
from opentelemetry.trace import NonRecordingSpan

trace.set_tracer_provider(
    TracerProvider(
        resource=Resource.create({SERVICE_NAME: "my-helloworld-service"})
    )
)
tracer = trace.get_tracer(__name__)


def handler(event, context):
    init_tracer(context.tracing.jaeger_endpoint)
    span_context = get_fc_span(context.tracing.span_context)
    start_my_span(trace.set_span_in_context(NonRecordingSpan(span_context)))
    return 'hello world'


def init_tracer(endpoint):
    # create a JaegerExporter
    jaeger_exporter = JaegerExporter(
        collector_endpoint=endpoint
    )

    # Create a SimpleSpanProcessor and add the exporter to it
    span_processor = SimpleSpanProcessor(jaeger_exporter)

    # add to the tracer
    trace.get_tracer_provider().add_span_processor(span_processor)


def get_fc_span(jaeger_span_context):
    jaeger_span_context_arr = jaeger_span_context.split(":")
    tid = int(jaeger_span_context_arr[0], 16)
    sid = int(jaeger_span_context_arr[1], 16)

    span_context = trace.SpanContext(
        trace_id=tid,
        span_id=sid,
        is_remote=True,
        trace_flags=trace.TraceFlags(0x01),
    )
    return span_context


def start_my_span(context):
    with tracer.start_as_current_span(name="fc-operation", context=context):
        time.sleep(0.15)
        with tracer.start_as_current_span("child"):
            time.sleep(0.1)
