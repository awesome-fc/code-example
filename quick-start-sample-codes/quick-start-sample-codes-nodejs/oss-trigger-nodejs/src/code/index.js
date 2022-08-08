'use strict'
/*
本代码样例主要实现以下功能:
*   1. 从 event 中解析出 OSS 事件触发相关信息
*   2. 根据以上获取的信息，初始化 OSS 客户端
*   3. 从目标OSS bucket中获取事件触发的object
*   4. 将获取得到的object进行备份到OSS bucket中


This code sample mainly implements the following functions:
*1. Parse the OSS event trigger related information from the event
* 2. According to the above information, initialize the OSS client
* 3. Get the object triggered by the event from the target OSS bucket
* 4. Back up the acquired objects to the OSS bucket
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

    // 从OSS中获取文件buffer
    // Get file buffer from OSS
    let objectBuffer = await getBuffer(client, objectName);
    // 将文件buffer进行备份到OSS中
    await putBuffer(client, objectBuffer, 'copy/' + objectName)

    callback(null, "done");
}

// 从OSS中下载文件
// Download files from OSS
async function putBuffer(client, objectBuffer, objectName) {
    try {
        console.log("上传文件备份:" + objectName);
        const result = await client.put(objectName, objectBuffer.content);
        console.log(result);
    } catch (e) {
        console.log(e);
    }
}

// 上传文件到OSS中
// Upload files to OSS
async function getBuffer(client, objectName) {
    try {
        console.log("下载:" + objectName);
        const objectBuffer = await client.get(objectName);
        console.log(objectBuffer.content);
        return objectBuffer
    } catch (e) {
        console.log(e);
    }
}