'use strict';

exports.handler = function (event, context, callback) {
   console.log("whole event: %s", event);
   // Parse the json array
   var eventObj = JSON.parse(event.toString());

   for (let i = 0; i < eventObj.length; ++i) {
      var evt = JSON.parse(eventObj[i]);
      var messageTopic = evt['data']['topic']
      var message = evt['data']['value']

      console.log("kafka message topic:", messageTopic)
      console.log("kafka message :", message);
   }
   
   callback(null, "Kafka Trigger Event:" + eventObj);
};