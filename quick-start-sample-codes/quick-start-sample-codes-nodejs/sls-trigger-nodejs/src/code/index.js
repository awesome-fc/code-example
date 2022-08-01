/*
* 本代码样例主要实现以下功能:
*   1. 从 event 中解析出 SLS 事件触发相关信息
*   2. 根据以上获取的信息，初始化 SLS 客户端
*   3. 从源日志仓库中获取实时日志数据
*
*
* This sample code is mainly doing the following things:
*   1. Get SLS processing related information from event
*   2. Initiate SLS client
*   3. Pull logs from source log store
* */

'use strict'

const ALY = require('aliyun-sdk');
var util = require('util');

exports.handler = (event, context, callback) => {
    // 可以通过 context.credentials 获取密钥信息
    // Access keys can be fetched through context.credentials
    console.log("The content in context entity is: \n");
    console.dir(context);
    const { accessKeyId, accessKeySecret, securityToken } = context.credentials;

    // 解析 event 参数至 object 格式
    // parse event in object
    const eventObj = JSON.parse(event.toString());
    console.log("The content in event entity is: \n");
    console.dir(eventObj);

    // 从 event 中获取 cursorTime，该字段表示本次函数调用包括的数据中，最后一条日志到达日志服务的服务器端的 unix_timestamp
    // Get cursorTime from event, where cursorTime indicates that in the data of the invocation, the unix timestamp of the last log arrived at log store
    const { cursorTime } = eventObj

    // 从 event.source 中获取日志项目名称、日志仓库名称以及日志服务访问 endpoint
    // Get the name of log project, the name of log store and the endpoint of sls from event.source
    const { projectName, logstoreName, endpoint } = eventObj.source;

    // 初始化 sls 客户端
    // Initialize client of sls
    const client = new ALY.SLS({
        accessKeyId: accessKeyId,
        secretAccessKey: accessKeySecret,
        securityToken: securityToken,
        endpoint: endpoint,
        apiVersion: '2015-06-01'                         //SDK版本号，固定值。
    });

    // 从环境变量中获取触发时间间隔，该环境变量可在 s.yml 中配置
    // Get interval of trigger from environment variables, which was configured via s.yaml
    const triggerInterval = process.env['triggerInterval'];

    // 从源数据库中拉取日志数据
    // Get data from source logstore
    readLogs(client, projectName, logstoreName, triggerInterval, cursorTime).then((data) => {
        callback(null, data);
    }).catch((err) => {
        callback(err, '');
    });

}

// readLogs 用于从指定日志库中读取日志内容并返回，日志的时间区间为
// writeLogs is used to write contents into the specified log store
function readLogs(client, logProject, logStore, triggerInterval, cursorTime) {
    // 获取触发时间间隔内的日志数据
    // Get logs within the trigger interval
    const to = cursorTime;
    const from = to - triggerInterval;

    const param = {
        projectName: logProject,                           // 必选，Project名称。
        logStoreName: logStore,                            // 必选，Logstore名称。
        from: from,                                       // 必选，开始时间，精度为秒。
        to: to,                                           // 必选，结束时间，精度为秒
        topic: "",
        query: ""
    };
    return new Promise((resolve, reject) => {
        client.getLogs(param, function (err, data) {
            if(err) {
                console.error(util.format('Get logs from logProject: %s, logStore: %s failed due to error: \n', logProject, logStore));
                console.error(err);
                reject(err);
            } else {
                console.log('Log data from source logstore: ', data);
                resolve(data);
            }
        })
    });
}