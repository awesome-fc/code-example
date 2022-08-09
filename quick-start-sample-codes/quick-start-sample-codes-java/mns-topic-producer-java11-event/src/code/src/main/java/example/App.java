package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.PreStopHandler;
import com.aliyun.fc.runtime.StreamRequestHandler;
import com.aliyun.mns.client.CloudAccount;
import com.aliyun.mns.client.CloudTopic;
import com.aliyun.mns.client.MNSClient;
import com.aliyun.mns.model.Base64TopicMessage;
import com.aliyun.mns.model.TopicMessage;

public class App implements StreamRequestHandler, FunctionInitializer, PreStopHandler {

    MNSClient mnsClient = null;
    CloudTopic topic = null;
    String initErrorMessage = null;

    @Override
    public void initialize(Context context) throws IOException {
        String MNS_ENDPOINT = System.getenv("MNS_ENDPOINT");
        String MNS_TOPIC = System.getenv("MNS_TOPIC_NAME");
        String accessKeyId = context.getExecutionCredentials().getAccessKeyId();
        String accessKeySecret = context.getExecutionCredentials().getAccessKeySecret();
        if (accessKeyId.length() == 0 || accessKeySecret.length() == 0) {
            initErrorMessage = "service role is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        if (MNS_ENDPOINT.length() == 0) {
            initErrorMessage = "mns endpoint is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        if (MNS_TOPIC.length() == 0) {
            initErrorMessage = "mns topic is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }

        long start = System.currentTimeMillis();
        CloudAccount account = new CloudAccount(
            accessKeyId, 
            accessKeySecret, 
            MNS_ENDPOINT);
        mnsClient = account.getMNSClient(); // 在程序中，CloudAccount以及MNSClient单例实现即可，多线程安全。
        topic = mnsClient.getTopicRef(MNS_TOPIC);
        context.getLogger().info("init mns client time cost: " + (System.currentTimeMillis() - start) + "ms");
    }

    @Override
    public void preStop(Context context) throws IOException {
        context.getLogger().info("preStop start");
        if (mnsClient != null) {
            mnsClient.close();
            mnsClient = null;
        }
        context.getLogger().info("preStop end");
    }

    @Override
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        if (mnsClient == null && initErrorMessage != null) {
            outputStream.write(initErrorMessage.getBytes());
            return;
        }
        String result = "";
        try {
            TopicMessage msg = new Base64TopicMessage(); //可以使用TopicMessage结构，选择不进行Base64加密。
            msg.setMessageBody("hello world!");
            //msg.setMessageTag("filterTag"); //设置该条发布消息的filterTag。
            msg = topic.publishMessage(msg);
            result = "publish message succ, message id:"+msg.getMessageId()
                +", message body md5:"+msg.getMessageBodyMD5();
            context.getLogger().info(result);
        } catch (Exception e) {
            context.getLogger().info("publish message error: " + e.getMessage());
            e.printStackTrace();
            throw new RuntimeException(e.getMessage());
        } 
        outputStream.write(result.getBytes());
    }
}
