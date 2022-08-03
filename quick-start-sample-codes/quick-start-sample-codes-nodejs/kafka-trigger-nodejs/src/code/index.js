'use strict';

exports.handler = function (event, context, callback) {
   console.log("whole event: %s", event);
   // Parse the json
   var eventObj = JSON.parse(event.toString());
   var eventObjStr = eventObj.toString();
   var evt = JSON.parse(eventObjStr);

   var messageTopic = evt['data']['topic']
   var message = evt['data']['value']

   console.log("kafka message topic:", messageTopic)
   console.log("kafka message :", message);
   
   callback(null, "Kafka trigger data value:" + message);
};