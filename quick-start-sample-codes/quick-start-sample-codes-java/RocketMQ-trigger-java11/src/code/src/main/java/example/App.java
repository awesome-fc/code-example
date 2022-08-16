package example;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.aliyun.fc.runtime.*;
import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.OSSObject;

import java.io.*;
import java.nio.charset.StandardCharsets;

public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(
            InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();

        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String jsons = result.toString(StandardCharsets.UTF_8.name());

        JSONObject event = JSONArray.parseObject(jsons);
        JSONObject eventdata = event.getJSONObject("data");
        String messageBody = eventdata.getString("body");
        logger.info("message body: " + messageBody);
        outputStream.write(messageBody.getBytes());

    }
}