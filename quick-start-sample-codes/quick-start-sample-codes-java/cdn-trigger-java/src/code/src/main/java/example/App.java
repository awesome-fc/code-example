package example;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.FunctionComputeLogger;
import com.aliyun.fc.runtime.StreamRequestHandler;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;

import org.json.JSONArray;
import org.json.JSONObject;

/**
 * 本代码样例主要实现以下功能:
 * * 打印 event 信息
 * <p>
 * <p>
 * This sample code is mainly doing the following things:
 * * Print event
 */

// 各 event 示例见文档：https://help.aliyun.com/document_detail/75123.html，event结构如下所示：
//
// {  "events": [
//       {
//          "eventName": "***",
//          "eventVersion": "***",
//          "eventSource": "***",
//          "region": "***",
//          "eventTime": "***",
//          "traceId": "***",
//          "resource": {
//               "domain": "***"
//          },
//          "eventParameter": {
//
//          },
//          "userIdentity": {
//               "aliUid": "***"
//          }
//       }
//    ]
// }
public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(InputStream input, OutputStream output, Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();
        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = input.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String sEvt = result.toString(StandardCharsets.UTF_8.name());
        logger.info(sEvt);
        
        JSONObject obj = new JSONObject(sEvt);
        JSONObject event = obj.getJSONArray("events").getJSONObject(0);
        String eventName = event.getString("eventName");
        JSONObject eventParam = event.getJSONObject("eventParameter");
        String domain = eventParam.getString("domain");
        String info = "";
        if(eventName.equals("CachedObjectsRefreshed") || eventName.equals("CachedObjectsPushed") || eventName.equals("CachedObjectsBlocked")){
            JSONArray objPathList = eventParam.getJSONArray("objectPath");
            for(int i = 0 ; i < objPathList.length(); i ++){
                info += objPathList.getString(i) + ", ";
            }
        }else if(eventName.equals("LogFileCreated")){
            info = eventParam.getString("filePath");
        }else if(eventName.equals("CdnDomainStarted") || eventName.equals("CdnDomainStopped")){
            // 对应业务逻辑...
        }else if (eventName.equals("CdnDomainAdded") || eventName.equals("CdnDomainDeleted")) {
            // 对应业务逻辑...
        }
        String res = String.format("eventName:%s, domain:%s, info:%s\n", eventName, domain, info);
        output.write(res.getBytes());
    }
}