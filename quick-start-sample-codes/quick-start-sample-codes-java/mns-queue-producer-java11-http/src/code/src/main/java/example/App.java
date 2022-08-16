package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.PreStopHandler;
import com.aliyun.fc.runtime.HttpRequestHandler;
import com.aliyun.mns.client.CloudAccount;
import com.aliyun.mns.client.CloudQueue;
import com.aliyun.mns.client.MNSClient;
import com.aliyun.mns.common.ClientException;
import com.aliyun.mns.common.ServiceException;
import com.aliyun.mns.model.Message;

public class App implements HttpRequestHandler, FunctionInitializer, PreStopHandler {

    MNSClient mnsClient = null;
    CloudQueue queue = null;
    String initErrorMessage = null;

    @Override
    public void initialize(Context context) throws IOException {
        String MNS_ENDPOINT = System.getenv("MNS_ENDPOINT");
        String MNS_QUEUE = System.getenv("MNS_QUEUE_NAME");
        String accessKeyId = context.getExecutionCredentials().getAccessKeyId();
        String accessKeySecret = context.getExecutionCredentials().getAccessKeySecret();
        if (accessKeyId.length() == 0 || accessKeySecret.length() == 0) {
            initErrorMessage = "service role is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        if (MNS_ENDPOINT == null || MNS_ENDPOINT.length() == 0) {
            initErrorMessage = "mns endpoint is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }
        if (MNS_QUEUE == null || MNS_QUEUE.length() == 0) {
            initErrorMessage = "mns queue is not set";
            context.getLogger().error(initErrorMessage);
            return;
        }

        long start = System.currentTimeMillis();
        CloudAccount account = new CloudAccount(
            accessKeyId, 
            accessKeySecret, 
            MNS_ENDPOINT);
        mnsClient = account.getMNSClient();
        queue = mnsClient.getQueueRef(MNS_QUEUE);
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
    public void handleRequest(HttpServletRequest request, HttpServletResponse response, Context context) 
            throws IOException, ServletException {
        OutputStream outputStream = response.getOutputStream();
        if (mnsClient == null && initErrorMessage != null) {
            response.setStatus(500);
            outputStream.write(initErrorMessage.getBytes());
            return;
        }
        String result = "";
        try {
            Message message = new Message();
            message.setMessageBody("demo_message_body");
            Message putMsg = queue.putMessage(message);
            context.getLogger().info("Send message id is: " + putMsg.getMessageId());
            result = "Send message succ, message id:"+putMsg.getMessageId()
                +", message body:"+message.getMessageBody();
            context.getLogger().info(result);
        } catch (ClientException ce) {
            context.getLogger().error("Something wrong with the network connection between client and MNS service."
            + "Please check your network and DNS availability.");
            ce.printStackTrace();
            throw new RuntimeException(ce.getMessage());
        } catch (ServiceException se) {
            if (se.getErrorCode().equals("QueueNotExist")) {
                context.getLogger().error("Queue is not exist. Please create before use");
            } else if (se.getErrorCode().equals("TimeExpired")) {
                context.getLogger().error("The request is time expired. Please check your local machine timeclock");
            }
            se.printStackTrace();
            throw new RuntimeException(se.getMessage());
        } catch (Exception e) {
            context.getLogger().error("Unknown exception happened!");
            e.printStackTrace();
            throw new RuntimeException(e.getMessage());
        }
        response.setStatus(200);
        outputStream.write(result.getBytes());
    }
}
