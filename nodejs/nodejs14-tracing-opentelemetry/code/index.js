'use strict';

const api = require('@opentelemetry/api');
var tracer

module.exports.handler = function(event, context, callback)
{
    tracer = require('./tracer')('fc-tracer',context.tracing.jaegerEndpoint);

    var spanContext = contextFromString( context.tracing.openTracingSpanContext);

    startMySpan(spanContext);

    callback(null,'success');
}

function contextFromString(value){
    const arr = value.split(`:`);
    const spanContext={
        traceId:`0000000000000000`+arr[0],
        spanId:arr[1],
        traceFlags:api.TraceFlags.SAMPLED,
        isRemote:true
    }
    return spanContext;
}

function sleep(delay) {
    var start = (new Date()).getTime();
    while ((new Date()).getTime() - start < delay) {
        continue;
    }
}

function startMySpan(spanContext){
    var FcSpan=api.trace.wrapSpanContext(spanContext);
    var ctx = api.trace.setSpan(api.ROOT_CONTEXT,FcSpan);
    tracer.startActiveSpan("fc-operation",undefined,ctx,parentSpan => {
        parentSpan.setAttribute("version","fc-v1");
        sleep(150);
        child();
        parentSpan.end()
    })
}

function child(){
    tracer.startActiveSpan("fc-operation-child",span =>{
        sleep(100);
        span.addEvent("timeout");
        span.end();
    })
}