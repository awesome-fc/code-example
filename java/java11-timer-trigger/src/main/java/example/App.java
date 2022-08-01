package example;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.StreamRequestHandler;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

public class App implements StreamRequestHandler{

    @Override
      public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
         // Read the inputStream
         ByteArrayOutputStream result = new ByteArrayOutputStream();
         byte[] buffer = new byte[1024];
         for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
         }
         String eventString = result.toString(StandardCharsets.UTF_8.name());

         // Parse the json object
         JSONObject event = JSONObject.parseObject(eventString);

         String triggerTime = event.getString("triggerTime");
         String triggerName = event.getString("triggerName");
         String payload = event.getString("payload");

         context.getLogger().info("triggerTime: " + triggerTime);
         context.getLogger().info("triggerName: " + triggerName);
         context.getLogger().info("payload: " + payload);

         outputStream.write(("Timer Payload: " + payload).getBytes());
   }
        
}
