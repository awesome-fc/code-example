package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.ByteArrayOutputStream;
import java.time.*;
import java.time.format.*;
import java.util.Base64;

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
        context.getLogger().info("Input data: " + res);
        MnsQueueMessageFromEB mnsMessage = new MnsQueueMessageFromEB();

        try {
            ObjectMapper objectMapper = new ObjectMapper();
            mnsMessage = objectMapper.readValue(result.toByteArray(), MnsQueueMessageFromEB.class);
            logger.info("mns message: "+mnsMessage);

            // 事件触发耗时
            DateTimeFormatter dtf = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'");
            LocalDateTime publishTime = LocalDateTime.parse(mnsMessage.getAliyunpublishtime(), dtf);
            Duration d = Duration.between(publishTime, LocalDateTime.now());
            context.getLogger().info("publish time:"+publishTime);
            context.getLogger().info("message trigger time cost:"+d);

            // 在此处添加消息处理逻辑

            // 事件中的消息默认是base64编码的，需要进行解码
            // 若希望自动解码，可在 s.yaml 中设置 IsBase64Decode: true
            outputStream.write(mnsMessage.getData().getMessageBody().getBytes());
        } catch (Exception ex) {
            context.getLogger().error("mns message is not in json format.");
            outputStream.write(res.getBytes());
            throw new RuntimeException(ex.getMessage());
        }
    }
}
