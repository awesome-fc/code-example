package example;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.PojoRequestHandler;

public class PojoHandler implements PojoRequestHandler<SimpleRequest, SimpleResponse> {

    @Override
    public SimpleResponse handleRequest(SimpleRequest request, Context context) {
        String message = "Hello, " + request.getFirstName() + " " + request.getLastName();
        System.out.println(message);
        return new SimpleResponse(message);
    }
}
