package example;

import java.io.*;
import java.nio.charset.StandardCharsets;
import com.aliyun.fc.runtime.*;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.aliyun.openservices.log.Client;
import com.aliyun.openservices.log.response.GetLogsResponse;
import com.aliyun.openservices.log.common.QueriedLog;
import com.aliyun.openservices.log.common.LogItem;
import com.aliyun.openservices.log.exception.LogException;
import com.aliyun.openservices.log.common.auth.DefaultCredentails;

import java.util.Vector;

/**
 * 本代码样例主要实现以下功能:
 * * 从 event 中解析出 SLS 事件触发相关信息
 * * 根据以上获取的信息，初始化 SLS 客户端
 * * 从源 log store 获取实时日志数据
 * <p>
 * <p>
 * This sample code is mainly doing the following things:
 * * Get SLS processing related information from event
 * * Initiate SLS client
 * * Pull logs from source log store
 */

public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(
            InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String jsons = result.toString(StandardCharsets.UTF_8.name());

        JSONObject event = JSONArray.parseObject(jsons);

        // 从 event 中获取 cursorTime，该字段表示本次函数调用包括的数据中，最后一条日志到达日志服务的服务器端的 unix_timestamp
        // Get cursorTime from event, where cursorTime indicates that in the data of the invocation, the unix timestamp of the last log arrived at log store
        int cursorTime = Integer.parseInt(event.getString("cursorTime"));

        // 从 event.source 中获取日志项目名称、日志仓库名称以及日志服务访问 endpoint
        // Get the name of log project, the name of log store and the endpoint of sls from event.source
        JSONObject sls = event.getJSONObject("source");

        String endpoint = sls.getString("endpoint");
        String projectName = sls.getString("projectName");
        String logstoreName = sls.getString("logstoreName");

        // 从环境变量中获取目标日志仓库名称以及触发时间间隔，该环境变量可在 s.yml 中配置
        // Get interval of trigger from environment variables, which was configured via s.yaml
        int triggerInterval = Integer.parseInt(System.getenv("triggerInterval"));

        // 获取密钥信息，执行前，确保函数所在的服务配置了角色信息，并且角色需要拥有AliyunSLSFullAccess权限
        Credentials creds = context.getExecutionCredentials();

        // 初始化 sls 客户端
        // Initialize client of sls
        DefaultCredentails credsOfSLS = new DefaultCredentails(creds.getAccessKeyId(), creds.getAccessKeySecret(), creds.getSecurityToken());
        Client client = new Client(endpoint, credsOfSLS, null);


        Vector<LogItem> logItems = new Vector<LogItem>();
        try {
            // 从源日志库中读取日志
            // Read data from source logstore
            int toTime = cursorTime;
            int fromTime = toTime - triggerInterval;
            GetLogsResponse getLogsResponse = client.GetLogs(projectName, logstoreName, fromTime, toTime, "", "");
            context.getLogger().info("Read log data count:" + getLogsResponse.GetCount());
            context.getLogger().info("from time is: " + fromTime);
            context.getLogger().info("to time is: " + toTime);
            for (QueriedLog log : getLogsResponse.GetLogs()) {
                LogItem item = log.GetLogItem();
                context.getLogger().info("log time: " + item.mLogTime);
                context.getLogger().info("Jsonstring: " + item.ToJsonString());
            }
        } catch (LogException e) {
            context.getLogger().error("Read log data failed");
            context.getLogger().error("err code: " + e.GetErrorCode());
            context.getLogger().error("err message: " + e.GetErrorMessage());
            context.getLogger().error("err requestId: " + e.GetErrorCode());
            System.exit(-1);
        }


        outputStream.write(new String("success").getBytes());
    }
}
