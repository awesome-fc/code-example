'use strict';

 exports.handler = function (event, context, callback) {
    console.log("whole event: %s", event);
    
    // Parse the json
    var eventObj = JSON.parse(event.toString());

    var triggerName = eventObj['triggerName'];
    var triggerTime = eventObj['triggerTime'];
    var payload = eventObj['payload'];

    console.log("triggerName: ", triggerName);
    console.log("triggerTime: ", triggerTime);
    console.log("triggerMessgae: ", payload);

    callback(null, "Timer Payload:" + payload);
 };