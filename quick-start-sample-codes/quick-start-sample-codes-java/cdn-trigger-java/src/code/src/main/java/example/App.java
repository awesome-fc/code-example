package example;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionComputeLogger;
import com.aliyun.fc.runtime.StreamRequestHandler;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;

/**
 * 本代码样例主要实现以下功能:
 * * 打印 event 信息
 * <p>
 * <p>
 * This sample code is mainly doing the following things:
 * * Print event
 */

public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(InputStream input, OutputStream output, Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();
        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = input.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String res = result.toString(StandardCharsets.UTF_8.name());
        logger.info(res);
        output.write(res.getBytes());
    }
}