/*
* 本代码样例主要实现以下功能:
*   1. 打印 event 信息

* This sample code is mainly doing the following things:
*  1. print event
* */

exports.handler = (event, context, callback) => {
    console.log(event.toString())
    var eventObj = JSON.parse(event.toString());
    callback(null, eventObj);
};