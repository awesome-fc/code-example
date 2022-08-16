package example;

import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.List;
import com.aliyun.fc.runtime.*;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;

import com.aliyun.openservices.log.Client;
import com.aliyun.openservices.log.exception.LogException;
import com.aliyun.openservices.log.common.auth.DefaultCredentails;
import com.aliyun.openservices.log.request.PullLogsRequest;
import com.aliyun.openservices.log.response.PullLogsResponse;
import com.aliyun.openservices.log.common.LogGroupData;
import com.aliyun.openservices.log.common.LogItem;
import com.aliyun.openservices.log.common.LogContent;

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

        // 从 event.source 中获取日志项目名称、日志仓库名称、日志服务访问 endpoint、日志起始游标、日志终点游标以及分区 id
        // Get the name of log project, the name of log store, the endpoint of sls, begin cursor, end cursor and shardId from event.source
        JSONObject sls = event.getJSONObject("source");

        String endpoint = sls.getString("endpoint");
        String projectName = sls.getString("projectName");
        String logstoreName = sls.getString("logstoreName");
        String beginCursor = sls.getString("beginCursor");
        String endCursor = sls.getString("endCursor");
        int shardId = Integer.parseInt(sls.getString("shardId"));


        Credentials creds = context.getExecutionCredentials();

        // 初始化 sls 客户端
        // Initialize client of sls
        DefaultCredentails credsOfSLS = new DefaultCredentails(creds.getAccessKeyId(), creds.getAccessKeySecret(), creds.getSecurityToken());
        Client client = new Client(endpoint, credsOfSLS, null);

        try {
            while (true) {
                PullLogsRequest request = new PullLogsRequest(projectName, logstoreName, shardId, 100, beginCursor, endCursor);
                PullLogsResponse response = client.pullLogs(request);
                int logCount = response.getCount();
                if (logCount == 0) {
                    break;
                }
                context.getLogger().info("get " + logCount + " log group from " + logstoreName);
                List<LogGroupData> logGroups = response.getLogGroups();
                for (LogGroupData logGroup : logGroups) {

                    for (LogItem log : logGroup.GetAllLogs()) {
                        context.getLogger().info("LogTime:" + log.GetTime());
                        List<LogContent> contents = log.GetLogContents();
                        for (LogContent content : contents) {
                            context.getLogger().info(content.GetKey() + ":" + content.GetValue());
                        }
                    }
                }
                beginCursor = response.getNextCursor();;
            }
        } catch (LogException e) {
            context.getLogger().error("Pull log data failed");
            context.getLogger().error("err code: " + e.GetErrorCode());
            context.getLogger().error("err message: " + e.GetErrorMessage());
            context.getLogger().error("err requestId: " + e.GetErrorCode());
            System.exit(-1);
        }

        outputStream.write(new String("success").getBytes());
    }
}
