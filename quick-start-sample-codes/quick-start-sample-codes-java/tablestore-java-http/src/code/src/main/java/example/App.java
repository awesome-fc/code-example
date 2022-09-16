package example;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.alicloud.openservices.tablestore.SyncClient;
import com.alicloud.openservices.tablestore.model.GetRowResponse;
import com.alicloud.openservices.tablestore.model.GetRowRequest;
import com.alicloud.openservices.tablestore.model.PrimaryKey;
import com.alicloud.openservices.tablestore.model.PrimaryKeyBuilder;
import com.alicloud.openservices.tablestore.model.PrimaryKeyValue;
import com.alicloud.openservices.tablestore.model.Row;
import com.alicloud.openservices.tablestore.model.SingleRowQueryCriteria;
import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.Credentials;
import com.aliyun.fc.runtime.FunctionInitializer;
import com.aliyun.fc.runtime.HttpRequestHandler;
import com.aliyun.fc.runtime.PreStopHandler;

public class App implements HttpRequestHandler, FunctionInitializer, PreStopHandler {

    private SyncClient client = null;

    @Override
    public void initialize(Context context) {
        // 在initialize回调中创建客户端，可以实现在整个函数实例生命周期内复用该客户端
        Credentials creds = context.getExecutionCredentials();
        String endpoint = System.getenv("ENDPOINT"), instanceName = System.getenv("INSTANCE_NAME");
        String accessKeyId = creds.getAccessKeyId(), accessKeySecret = creds.getAccessKeySecret(), stsToken = creds.getSecurityToken();
        client = new SyncClient(endpoint, accessKeyId, accessKeySecret, instanceName, stsToken);
    }

    @Override
    public void preStop(Context context) {
        if(client != null) client.shutdown();
    }

    @Override
    public void handleRequest(HttpServletRequest request, HttpServletResponse response, Context context) throws IOException{
        // 本示例所用表格存储的主键包含两个主键列：region 和 id
        PrimaryKey primaryKey = PrimaryKeyBuilder.createPrimaryKeyBuilder()
        .addPrimaryKeyColumn("region", PrimaryKeyValue.fromString(request.getParameter("region")))
        .addPrimaryKeyColumn("id", PrimaryKeyValue.fromLong(Long.parseLong(request.getParameter("id"))))
        .build();

        // 读取一行数据，设置数据表名称。
        SingleRowQueryCriteria criteria = new SingleRowQueryCriteria(System.getenv("TABLE_NAME"), primaryKey);
        // 设置读取最新版本
        criteria.setMaxVersions(1);
        GetRowResponse getRowResponse = client.getRow(new GetRowRequest(criteria));
        Row row = getRowResponse.getRow();
        response.setStatus(200);
        response.setHeader("Content-Type", "application/json");
        OutputStream out = response.getOutputStream();
        out.write(row.toString().getBytes());
        out.flush();
        out.close();
    }
}
