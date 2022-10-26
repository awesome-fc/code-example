var initTracer = require('jaeger-client').initTracer;
var spanContext = require('jaeger-client').SpanContext;
var tracer;

module.exports.handler = function(event, context, callback)
{
    tracer=newTracer(context);

    var invokeSpanContext = spanContext.fromString(context.tracing.openTracingSpanContext);

    startMySpan(invokeSpanContext);

    callback(null,'success')
}

function newTracer(context){
    var config = {
        serviceName: 'fc-tracer',
        reporter: {
            // spans over HTTP
            collectorEndpoint: context.tracing.jaegerEndpoint,
            flushIntervalMs: 10,
        },
        sampler: {
            type: "const",
            param: 1
        },
    };
    var options = {
        tags: {
            'version': 'fc-v1',
        },
    };
    var tracer = initTracer(config, options);
    return tracer
}

function sleep(delay) {
    var start = (new Date()).getTime();
    while ((new Date()).getTime() - start < delay) {
        continue;
    }
}

function startMySpan(spanContext){
    var parentSpan = tracer.startSpan("fc-operation", {
        childOf: spanContext
    });
    sleep(150);
    child(parentSpan.context())
    parentSpan.finish();
}

function child(spanContext){
    var childSpan = tracer.startSpan("fc-operation-child", {
        childOf: spanContext
    });
    childSpan.log({event:"timeout"});
    sleep(100);
    childSpan.finish();
}
