package example;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.StreamRequestHandler;
import io.opentelemetry.api.trace.*;
import io.opentelemetry.api.GlobalOpenTelemetry;
import org.apache.thrift.transport.TTransportException;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.util.concurrent.TimeUnit;

public class App implements StreamRequestHandler {

    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        String endpoint = context.getTracing().getJaegerEndpoint();
        try {
            ExampleConfiguration.initOpenTelemetry(endpoint);
        } catch (TTransportException e) {
            throw new RuntimeException(e);
        }

        SpanContext spanContext = contextFromString(context.getTracing().getSpanContext());

        startMySpan(io.opentelemetry.context.Context.current().with(Span.wrap(spanContext)));
        }
    SpanContext contextFromString(String value) throws IOException {
        {
            if (value != null && !value.equals("")) {
                String[] parts = value.split(":");
                if (parts.length != 4) {
                    throw new RuntimeException(value);
                } else {
                    String traceId = parts[0];
                    if (traceId.length() <= 32 && traceId.length() >= 1) {
                        return SpanContext.createFromRemoteParent("0000000000000000"+parts[0], parts[1], TraceFlags.getSampled(), TraceState.getDefault());
                    } else {
                        throw new RuntimeException("Trace id [" + traceId + "] length is not withing 1 and 32");
                    }
                }
            } else {
                throw new RuntimeException();
            }
        }
    }

    void startMySpan(io.opentelemetry.context.Context ctx){
        Tracer tracer = GlobalOpenTelemetry.getTracer("fc-Trace");
        Span parentSpan = tracer.spanBuilder("fc-operation").setParent(ctx).startSpan();
        parentSpan.setAttribute("version","fc-v1");
        try {
            TimeUnit.MILLISECONDS.sleep(150);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        child(parentSpan.storeInContext(ctx));
        parentSpan.end();
    }

    void child(io.opentelemetry.context.Context ctx){
        Tracer tracer = GlobalOpenTelemetry.getTracer("fc-Trace");
        Span childSpan = tracer.spanBuilder("fc-operation-child").setParent(ctx).startSpan();
        childSpan.addEvent("timeout");
        try {
            Thread.sleep(100);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        childSpan.end();
    }
}


