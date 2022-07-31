'use strict';

 exports.handler = function (event, context, callback) {
    console.log("whole event: %s", event);
    // 解析JSON格式事件
    var eventObj = JSON.parse(event.toString());
    var messageTopic = eventObj['topic']
    console.log("kafka message topic:", messageTopic)
    var message = eventObj['value']
    console.log("kafka message :", message);
    callback(null, "Kafka trigger:" + event);
 };