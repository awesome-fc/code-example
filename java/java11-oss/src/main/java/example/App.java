package example;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;

import com.aliyun.fc.runtime.Context;
import com.aliyun.fc.runtime.Credentials;
import com.aliyun.fc.runtime.StreamRequestHandler;
import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;

public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(InputStream inputStream, OutputStream outputStream, Context context) throws IOException {

        // Bucket 名称, 需要预先创建
        String bucketName = "my-bucket";
        // Endpoint必须填写Bucket所在地域对应的Endpoint，推荐使用内网访问地址。以华东1（杭州）为例，内网访问Endpoint为https://oss-cn-hangzhou-internal.aliyuncs.com。
        String endpoint = "https://oss-cn-hangzhou-internal.aliyuncs.com";

        // 获取密钥信息，执行前，确保函数所在的服务配置了角色信息，并且角色需要拥有AliyunOSSFullAccess权限
        // 建议直接使用AliyunFCDefaultRole 角色
        Credentials creds = context.getExecutionCredentials();

        // 创建OSSClient实例。
        OSS ossClient = new OSSClientBuilder().build(endpoint, creds.getAccessKeyId(), creds.getAccessKeySecret(), creds.getSecurityToken());

        // 填写Byte数组。
        byte[] content = "Hello FC".getBytes();
        // 依次填写Bucket名称（例如examplebucket）和Object完整路径（例如exampledir/exampleobject.txt）。Object完整路径中不能包含Bucket名称。
        ossClient.putObject(bucketName, "exampledir/exampleobject.txt", new ByteArrayInputStream(content));

        // 关闭OSSClient
        ossClient.shutdown();
    
        outputStream.write(new String("done").getBytes());
    }
}