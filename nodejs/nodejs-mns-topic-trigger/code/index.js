'use strict';

exports.handler = function (event, context, callback) {
  console.log("event:", event.toString());
  callback(null, event);
};
