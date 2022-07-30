'use strict';

 exports.handler = function (event, context, callback) {
    console.log("whole event: %s", event);
    // 解析JSON格式事件
    var eventObj = JSON.parse(event.toString());
    var triggerName = eventObj['triggerName']
    console.log("triggerName: ", triggerName)
    var triggerTime = eventObj['triggerTime']
    console.log("triggerTime: ", triggerTime);
    var message = eventObj['payload']
    console.log("triggerMessgae: ", message);
    callback(null, "timer trigger:" + message);
 };