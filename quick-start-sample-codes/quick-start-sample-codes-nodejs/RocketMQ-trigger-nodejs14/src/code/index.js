'use strict';

exports.handler = function (event, context, callback) {
  console.log("Receive mns topic whole message:", event.toString());
  // Parse the json
  try{
    var eventJson=JSON.parse(event.toString())
  } catch (err){
    callback(null, "the event format is STREAM and topic message content is:"+event);
  }
  if(eventJson.hasOwnProperty("Message")){
    callback(null, "the event format is JSON and topic message content is:"+eventJson.Message);
  }
  callback(null, "the event format is STREAM and topic message content is:"+event);
};