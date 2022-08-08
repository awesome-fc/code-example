/*本代码样例主要实现以下功能:
*   1. 从 request 中解析出 endpoint,bucket,object
*   2. 根据以上获取的信息，初始化 OSS 客户端
*   3. 从OSS中获取object的内容并将其上传到copy目录下实现备份
*
*This code sample mainly implements the following functions:
* 1. Parse out endpoint, bucket, object from request
* 2. According to the information obtained above, initialize the OSS client
* 3. Obtain the content of the object from OSS and upload it to the copy directory for backup
*/
<?php

use RingCentral\Psr7\Response;
use OSS\OssClient;
use OSS\Core\OssException;

function handler($request, $context): Response
{
    // 获取requestBody并将其解析为json
    // Get requestBody and parse it as json
    $requestBody = $request->getBody()->getContents();
    $jObj = json_decode($requestBody);
    // 从context中获取credentials
    // Get credentials from context
    $credentials = $context['credentials'];
    $accessKeyId = $credentials['accessKeyId'];
    $accessKeySecret = $credentials['accessKeySecret'];
    $securityToken = $credentials['securityToken'];

    $endpoint = $jObj->endpoint;
    $bucket = $jObj->bucket;
    $object = $jObj->object;

    try {
        // 连接OSS
        // Connect to OSS
        $ossClient = new OssClient($accessKeyId, $accessKeySecret, $endpoint, false, $securityToken);
        // 获取文件内容
        // Get the file content
        $content = $ossClient->getObject($bucket, $object);
        // 上传文件到copy目录下实现文件备份
        // Upload files to the copy directory to achieve file backup
        $ossClient->putObject($bucket, "copy/" . $object, $content);
    } catch (OssException $e) {
        print_r(__FUNCTION__ . ": FAILED\n");
        printf($e->getMessage() . "\n");
    }

    print(__FUNCTION__ . ": OK" . "\n");

    return new Response(
        200,
        array(
            'custom_header1' => 'v1',
            'custom_header2' => ['v2', 'v3'],
            'Content-Type' => 'text/plain',
        ),
        'done'
    );
}