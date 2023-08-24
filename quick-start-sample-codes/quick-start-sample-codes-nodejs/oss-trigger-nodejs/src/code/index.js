'use strict'
/*
本代码样例主要实现以下功能:
*   1. 从 event 中解析出 OSS 事件触发相关信息
*   2. 根据以上获取的信息，初始化 OSS 客户端
*   3. 将源图片 resize 后持久化到OSS bucket 下指定的目标图片路径，从而实现图片备份


This code sample mainly implements the following functions:
*1. Parse the OSS event trigger related information from the event
* 2. According to the above information, initialize the OSS client
* 3. Resize the source image and then store the processed image into the same bucket's copy folder to backup the image
*/

const OSS = require('ali-oss');

exports.handler = async function (event, context, callback) {

    console.log("The content in context entity is: \n");
    console.dir(context);

    const {accessKeyId, accessKeySecret, securityToken} = context.credentials;

    const events = JSON.parse(event.toString()).events;
    console.log("The content in event entity is: \n");
    console.dir(events);

    let objectName = events[0].oss.object.key;
    let region = events[0].region;
    let bucketName = events[0].oss.bucket.name;

    // 连接目标OSS
    // Connect to the target OSS
    const client = new OSS({
        region: region,
        accessKeyId: accessKeyId,
        accessKeySecret: accessKeySecret,
        stsToken: securityToken,
        bucket: bucketName,
        endpoint: "https://oss-" + region + "-internal.aliyuncs.com"
    });

    console.log("The client entity is: \n");
    console.dir(events);
    
    const targetImage = objectName.replace("source/", "processed/")
    // 将图片缩放为固定宽高128 px。
    const processStr = "image/resize,m_fixed,w_128,h_128"
    // 将源图片 resize 后再存储到目标图片路径
    const result = await client.processObjectSave(
      objectName,
      targetImage,
      processStr,
      bucketName
    );
    console.log(result.res.status);

    callback(null, "done");
}