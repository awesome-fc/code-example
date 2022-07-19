package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.ByteArrayOutputStream;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.PreStopHandler;
import com.aliyun.fc.runtime.FunctionComputeLogger;
import com.aliyun.fc.runtime.StreamRequestHandler;
import java.nio.charset.StandardCharsets;

import com.fasterxml.jackson.databind.ObjectMapper;

public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();
        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String res = result.toString(StandardCharsets.UTF_8.name());
        // Event 格式为 STREAM 时，传入的 Event 只包括消息信息
        // Event 格式为 JSON 时，传入的 Event 还包含 TopicName 等元数据。
        logger.info("Input data: " + res);
        MnsMessage mnsMessage = new MnsMessage();

        // 若选择 STREAM 的 Event 格式，在代码中则不需要将 Event 解析成json
        try{
            ObjectMapper objectMapper = new ObjectMapper();
            mnsMessage = objectMapper.readValue(result.toByteArray(), MnsMessage.class);
            logger.info("mns message: "+mnsMessage);
            outputStream.write(mnsMessage.getMessage().getBytes());
        } catch (Exception ex) {
            logger.info("mns message is not in json format.");
            outputStream.write(res.getBytes());
        }
    }
}
