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
      context.getLogger().info("Event: " + eventString);

      // Parse the json array
      JSONArray eventArray = JSONArray.parseArray(eventString);
      // Get the first json object from the json array
      JSONObject event = eventArray.getJSONObject(0);

      String topic = event.getString("topic");
      String value = event.getString("value");

      context.getLogger().info("Kafka Topic: " + topic);
      context.getLogger().info("Message Value: " + value);

      outputStream.write(("Produce ok, Topic: " + topic + " Value: " + value).getBytes());
   }
        
}
