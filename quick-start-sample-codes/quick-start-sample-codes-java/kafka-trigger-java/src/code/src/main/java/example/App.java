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
      context.getLogger().info("Event: " + eventString);

      // A rude way to deal with the string like this ["JsonObject"]->JsonObject
      // caution: this way can only deal with one event!
      String eventJsonString = eventString.substring(2, eventString.length() - 2);
      // Deal with Escape Character
      JSONObject event = JSON.parseObject(StringEscapeUtils.unescapeJava(eventJsonString));

      // Get the data field
      JSONObject eventData = event.getJSONObject("data");

      String topic = eventData.getString("topic");
      String value = eventData.getString("value");

      context.getLogger().info("Kafka Topic: " + topic);
      context.getLogger().info("Message Value: " + value);

      outputStream.write(("Produce ok, Topic: " + topic + " Value: " + value).getBytes());
   }
        
}
