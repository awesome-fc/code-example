package example;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.aliyun.fc.runtime.*;
import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.OSSObject;
import com.aliyun.oss.*;
import com.aliyun.oss.common.utils.BinaryUtil;
import com.aliyun.oss.common.utils.IOUtils;
import com.aliyun.oss.model.GenericResult;
import com.aliyun.oss.model.ProcessObjectRequest;

import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.Formatter;

/**
 * 本代码样例主要实现以下功能:
 * * 从 event 中解析出 OSS 事件触发相关信息
 * * 根据以上获取的信息，初始化 OSS bucket 客户端
 * * 将源图片 resize 后持久化到OSS bucket 下指定的目标图片路径，从而实现图片备份
 * <p>
 * <p>
 * This sample code is mainly doing the following things:
 * * Get OSS processing related information from event
 * * Initialize OSS client with target bucket
 * * Resize the source image and then store the processed image into the same bucket's copy folder to backup the image
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
        // String endpoint = "https://oss-" + region + "-internal.aliyuncs.com";
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
        String targetImage = objectName.replace("source/", "processed/");

        try {
          // 将图片缩放为固定宽高128 px。
          StringBuilder sbStyle = new StringBuilder();
          Formatter styleFormatter = new Formatter(sbStyle);
          String styleType = "image/resize,m_fixed,w_128,h_128";
          
          styleFormatter.format("%s|sys/saveas,o_%s,b_%s", styleType,
                  BinaryUtil.toBase64String(targetImage.getBytes()),
                  BinaryUtil.toBase64String(bucketName.getBytes()));
          System.out.println(sbStyle.toString());
          ProcessObjectRequest request = new ProcessObjectRequest(bucketName, objectName, sbStyle.toString());
          // 将源图片 resize 后再存储到目标图片路径
          GenericResult processResult = ossClient.processObject(request);
          String json = IOUtils.readStreamAsString(processResult.getResponse().getContent(), "UTF-8");
          processResult.getResponse().getContent().close();
          System.out.println(json);
      } catch (OSSException oe) {
          System.out.println("Caught an OSSException, which means your request made it to OSS, "
                  + "but was rejected with an error response for some reason.");
          System.out.println("Error Message:" + oe.getErrorMessage());
          System.out.println("Error Code:" + oe.getErrorCode());
          System.out.println("Request ID:" + oe.getRequestId());
          System.out.println("Host ID:" + oe.getHostId());
      } catch (ClientException ce) {
          System.out.println("Caught an ClientException, which means the client encountered "
                  + "a serious internal problem while trying to communicate with OSS, "
                  + "such as not being able to access the network.");
          System.out.println("Error Message:" + ce.getMessage());
      } finally {
          if (ossClient != null) {
              ossClient.shutdown();
          }
          outputStream.write(new String("done").getBytes());
      }
    }

}