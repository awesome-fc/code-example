package example;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.aliyun.fc.runtime.*;
import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.OSSObject;
import com.aliyun.mq.http.MQClient;
import com.aliyun.mq.http.MQProducer;
import com.aliyun.mq.http.model.TopicMessage;

import java.util.Date;

import java.io.*;
import java.nio.charset.StandardCharsets;

public class App implements StreamRequestHandler, FunctionInitializer {
    MQClient mqClient = null;
    String topic, instanceID;
    String initErrorMessage = null;

    @Override
    public void initialize(Context context) throws IOException {
        String accessKeyId = context.getExecutionCredentials().getAccessKeyId();
        String accessKeySecret = context.getExecutionCredentials().getAccessKeySecret();
        String securityToken = context.getExecutionCredentials().getSecurityToken();
        String endpoint = System.getenv("ROCKETMQ_ENDPOINT");
        topic = System.getenv("TOPIC");
        instanceID = System.getenv("INSTANCEID");
        if (accessKeyId.length() == 0 || accessKeySecret.length() == 0) {
            initErrorMessage = "service role is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        if (endpoint.length() == 0) {
            initErrorMessage = "RocketMQ endpoint is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        if (topic.length() == 0) {
            initErrorMessage = "topic is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        long start = System.currentTimeMillis();
        mqClient = new MQClient(endpoint, accessKeyId, accessKeySecret, securityToken);
    }

    @Override
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        if (mqClient == null && initErrorMessage != null) {
            outputStream.write(initErrorMessage.getBytes());
            return;
        }
        MQProducer producer = mqClient.getProducer(instanceID, topic);
        String result = "";
        try {
            TopicMessage pubMsg;
            pubMsg = new TopicMessage(
                    // 消息内容。
                    "hello mq!".getBytes(),
                    // 消息标签。
                    "A"
            );
            // 设置消息的自定义属性。
            pubMsg.getProperties().put("a", "1");
            // 设置消息的Key。
            pubMsg.setMessageKey("MessageKey");
            // 同步发送消息，只要不抛异常就是成功。
            TopicMessage pubResultMsg = producer.publishMessage(pubMsg);
            result = "publish messgae succ ,message id: " + pubResultMsg.getMessageId() + ", bodyMD5 is: " + pubResultMsg.getMessageBodyMD5();
            context.getLogger().info(result);
        } catch (Exception e) {
            context.getLogger().info("publish message error: "+e.getMessage());
            e.printStackTrace();
            throw new RuntimeException(e.getMessage());
        }
        outputStream.write(result.getBytes());
    }
}