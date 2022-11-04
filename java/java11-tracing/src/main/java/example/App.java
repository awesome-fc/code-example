package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.math.BigInteger;
import java.util.concurrent.TimeUnit;
import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.StreamRequestHandler;
import io.jaegertracing.internal.JaegerSpanContext;
import io.jaegertracing.internal.exceptions.EmptyTracerStateStringException;
import io.jaegertracing.internal.exceptions.MalformedTracerStateStringException;
import io.jaegertracing.internal.exceptions.TraceIdOutOfBoundException;
import io.opentracing.Span;
import io.opentracing.SpanContext;
import io.opentracing.Tracer;
import io.opentracing.util.GlobalTracer;


public class App implements StreamRequestHandler {

    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {

        registerTracer(context);

        JaegerSpanContext spanContext = contextFromString(context.getTracing().getSpanContext());

        startMySpan(spanContext);
    }
    static JaegerSpanContext contextFromString(String value) throws MalformedTracerStateStringException, EmptyTracerStateStringException {
        if (value != null && !value.equals("")) {
            String[] parts = value.split(":");
            if (parts.length != 4) {
                throw new MalformedTracerStateStringException(value);
            } else {
                String traceId = parts[0];
                if (traceId.length() <= 32 && traceId.length() >= 1) {
                    return new JaegerSpanContext(0L, (new BigInteger(traceId, 16)).longValue(), (new BigInteger(parts[1], 16)).longValue(), (new BigInteger(parts[2], 16)).longValue(), (new BigInteger(parts[3], 16)).byteValue());
                } else {
                    throw new TraceIdOutOfBoundException("Trace id [" + traceId + "] length is not withing 1 and 32");
                }
            }
        } else {
            throw new EmptyTracerStateStringException();
        }
    }

    void registerTracer(Context context){
        io.jaegertracing.Configuration config = new io.jaegertracing.Configuration("FCTracer");
        io.jaegertracing.Configuration.SenderConfiguration sender = new io.jaegertracing.Configuration.SenderConfiguration();
        sender.withEndpoint(context.getTracing().getJaegerEndpoint());
        config.withSampler(new io.jaegertracing.Configuration.SamplerConfiguration().withType("const").withParam(1));
        config.withReporter(new io.jaegertracing.Configuration.ReporterConfiguration().withSender(sender).withMaxQueueSize(10000));
        GlobalTracer.register(config.getTracer());
    }

    void startMySpan(SpanContext spanContext){
        Tracer tracer = GlobalTracer.get();
        Tracer.SpanBuilder spanBuilder = tracer.buildSpan("fc-operation").withTag("version", "fc-v1").asChildOf(spanContext);
        Span parentSpan = spanBuilder.start();
        try {
            TimeUnit.MILLISECONDS.sleep(150);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        child(parentSpan.context());
        parentSpan.finish();
    }

    void child(SpanContext spanContext){
        Tracer tracer = GlobalTracer.get();
        Tracer.SpanBuilder spanBuilder = tracer.buildSpan("fc-operation-child").asChildOf(spanContext);
        Span childSpan = spanBuilder.start();
        try {
            TimeUnit.MILLISECONDS.sleep(100);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        childSpan.log("timeout");
        childSpan.finish();
    }
}
