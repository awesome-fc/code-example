package example;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import java.util.List;
import java.lang.reflect.Type;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.StreamRequestHandler;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import org.apache.commons.lang.StringEscapeUtils;

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
      context.getLogger().info("Whole Event: " + eventString);

      JSONArray jsonArray = JSONArray.parseArray(eventString);

      for(int i = 0; i < jsonArray.size(); i++) {
         JSONObject event = JSONObject.parseObject(jsonArray.getString(i));

         // Get the data field
         JSONObject eventData = event.getJSONObject("data");
         String topic = eventData.getString("topic");
         String value = eventData.getString("value");

         context.getLogger().info("Message " + i + " Kafka Topic: " + topic);
         context.getLogger().info("Message " + i + " Message Value: " + value);
         outputStream.write(("Get Message " + i + ", Topic: " + topic + " Value: " + value).getBytes());
      }
   }
}
