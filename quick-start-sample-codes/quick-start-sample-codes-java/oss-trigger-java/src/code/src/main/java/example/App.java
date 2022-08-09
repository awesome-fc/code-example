package example;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.aliyun.fc.runtime.*;
import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.OSSObject;

import java.io.*;
import java.nio.charset.StandardCharsets;

/**
 * 本代码样例主要实现以下功能:
 * * 从 event 中解析出 OSS 事件触发相关信息
 * * 根据以上获取的信息，初始化 OSS bucket 客户端
 * * 从 OSS bucket 下载目标图片
 * * 将目标图片上传到 OSS bucket 下的 copy 目录实现图片备份
 * <p>
 * <p>
 * This sample code is mainly doing the following things:
 * * Get OSS processing related information from event
 * * Initialize OSS client with target bucket
 * * Download the target image from bucket
 * * Upload the image copy into the same bucket's copy folder to backup the image
 */
public class App implements StreamRequestHandler {

    @Override
    public void handleRequest(
            InputStream inputStream, OutputStream outputStream, Context context) throws IOException {
        FunctionComputeLogger logger = context.getLogger();

        ByteArrayOutputStream result = new ByteArrayOutputStream();
        byte[] buffer = new byte[1024];
        for (int length; (length = inputStream.read(buffer)) != -1; ) {
            result.write(buffer, 0, length);
        }
        String jsons = result.toString(StandardCharsets.UTF_8.name());

        JSONObject events = JSONArray.parseObject(jsons);
        // 取出首个事件
        // Get the first event
        JSONObject event = JSON.parseArray(events.getString("events").toString()).getJSONObject(0);

        JSONObject oss = event.getJSONObject("oss");
        String objectName = (String) oss.getJSONObject("object").getInnerMap().get("key");
        String region = event.getString("region");
        String bucketName = oss.getJSONObject("bucket").getString("name");
        // Endpoint必须填写Bucket所在地域对应的Endpoint。以华东1（杭州）为例，Endpoint填写为https://oss-cn-hangzhou.aliyuncs.com。
        // Endpoint must be the endpoint of bucket's region.For instance of hangzhou, the endpoint should be https://oss-cn-hangzhou.aliyuncs.com.
        String endpoint = "https://oss-" + region + "-internal.aliyuncs.com";

        // 获取密钥信息，执行前，确保函数所在的服务配置了角色信息，并且角色需要拥有AliyunOSSFullAccess权限
        // Obtain the key information. Before executing, make sure that the service where the function is located is configured with role information, and the role needs to have the AliyunOSSFullAccess permission.
        // 建议直接使用AliyunFCDefaultRole 角色
        // It is recommended to use AliyunFCDefaultRole directly
        Credentials creds = context.getExecutionCredentials();

        // 创建OSSClient实例。
        // Create the OSSClient instance
        OSS ossClient = new OSSClientBuilder().build(endpoint, creds.getAccessKeyId(), creds.getAccessKeySecret(), creds.getSecurityToken());

        OSSObject file = ossClient.getObject(bucketName, objectName);

        if (file == null) {
            logger.error(objectName + " doesn't exist");
            throw new FileNotFoundException();
        }
        // 依次填写Bucket名称（例如exampleBucket）和Object完整路径（例如exampleDir/exampleObject.txt）。Object完整路径中不能包含Bucket名称。
        // Fill in the Bucket name (for example, exampleBucket) and the full object path (for example, exampleDir/exampleObject.txt) in sequence. The bucket name cannot be included in the full path of Object.
        ossClient.putObject(bucketName, "copy/" + objectName, file.getObjectContent());

        // 关闭文件
        // Close the file
        file.close();
        // 关闭OSSClient
        // Close the OSSClient
        ossClient.shutdown();

        outputStream.write(new String("done").getBytes());
    }
}