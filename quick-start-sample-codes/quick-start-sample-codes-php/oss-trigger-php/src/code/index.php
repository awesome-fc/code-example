/*本代码样例主要实现以下功能:
*   1. 从 request 中解析出 endpoint,bucket,object
*   2. 根据以上获取的信息，初始化 OSS 客户端
*   3. 将源图片 resize 后持久化到OSS bucket 下指定的目标图片路径，从而实现图片备份
*
*This code sample mainly implements the following functions:
* 1. Parse out endpoint, bucket, object from request
* 2. According to the information obtained above, initialize the OSS client
* 3. Resize the source image and then store the processed image into the same bucket's copy folder to backup the image
*/
<?php

use RingCentral\Psr7\Response;
use OSS\OssClient;
use OSS\Core\OssException;

function base64url_encode($data)
{
    return rtrim(strtr(base64_encode($data), '+/', '-_'), '=');
}

function handler($event, $context) {
  $event           = json_decode($event, $assoc = true);
  /*
    阿里云账号AccessKey拥有所有API的访问权限，建议您使用RAM用户进行API访问或日常运维。
    建议不要把AccessKey ID和AccessKey Secret保存到工程代码里，否则可能导致AccessKey泄露，威胁您账号下所有资源的安全。
    本示例以从上下文中获取AccessKey/AccessSecretKey为例。
  */
  $accessKeyId     = $context["credentials"]["accessKeyId"];
  $accessKeySecret = $context["credentials"]["accessKeySecret"];
  $securityToken   = $context["credentials"]["securityToken"];
  $evt        = $event['events']{0};
  $bucketName = $evt['oss']['bucket']['name'];
  $endpoint   = 'oss-' . $evt['region'] . '-internal.aliyuncs.com';
  $objectName = $evt['oss']['object']['key'];
  $targetObject = str_replace("source/", "processed/", $objectName);

    try {
        // 连接OSS
        // Connect to OSS
        $ossClient = new OssClient($accessKeyId, $accessKeySecret, $endpoint, false, $securityToken);
        // 将图片缩放为固定宽高128 px
        $style = "image/resize,m_fixed,w_128,h_128";
        $process = $style.
           '|sys/saveas'.
           ',o_'.base64url_encode($targetObject).
           ',b_'.base64url_encode($bucketName);
        // 将图片 Resize 后保存到目标文件中
        $result = $ossClient->processObject($bucketName, $objectName, $process);
        // 打印处理结果。
        print($result);
    } catch (OssException $e) {
        print_r(__FUNCTION__ . ": FAILED\n");
        printf($e->getMessage() . "\n");
    }

    print(__FUNCTION__ . ": OK" . "\n");

    return $targetObject;
}