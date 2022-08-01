package example;

import java.io.*;

import com.aliyun.fc.runtime.*;
import com.aliyun.odps.Instance;
import com.aliyun.odps.Odps;
import com.aliyun.odps.OdpsException;
import com.aliyun.odps.account.Account;
import com.aliyun.odps.account.AliyunAccount;
import com.aliyun.odps.task.SQLTask;

/**
* 本代码样例主要实现以下功能:
* 1. 从 event 中解析事件列表。
* 2. 将解析的数据插入到 ODPS 数据表中。
*
* This sample code is mainly doing the following things:
* 1. Parse the event list from event param.
* 2. Insert payload into the ODPS table.
*/

public class App implements StreamRequestHandler {
    @Override
    public void handleRequest(
            InputStream input, OutputStream output, Context context) throws IOException {
        try {
            // 解析请求的 payload 数据
            // Parse request payload data
            byte[] buffer = new byte[1024];
            ByteArrayOutputStream payload = new ByteArrayOutputStream();
            int cnt;
            while((cnt = input.read(buffer)) != -1) {
                payload.write(buffer, 0, cnt);
            }
            String data = payload.toString();

            // 初始化 odps 客户端
            // Initialize odps client
            Account account = new AliyunAccount(System.getenv("accountAccessKeyID"), System.getenv("accountAccessKeySecret"));
            Odps odpsClient = new Odps(account);
            odpsClient.setEndpoint(System.getenv("odpsEndpoint"));
            odpsClient.setDefaultProject(System.getenv("odpsProject"));
            String odpsTableName = System.getenv("odpsTableName");

            // 将 payload 数据通过 sql 的方式插入到 odps 数据表中
            // Insert the payload data into the odps table
            String sql = " INSERT INTO " + odpsTableName + " VALUES (" + data.substring(1, data.length()-1) + ");";
            context.getLogger().info("run sql is: " + sql);
            Instance t = SQLTask.run(odpsClient, sql);
            t.waitForSuccess();
        } catch (OdpsException e) {
            context.getLogger().error("write data to odps failed " + e.toString());
            System.exit(-1);
        }
        output.write(new String("success").getBytes());
    }
}