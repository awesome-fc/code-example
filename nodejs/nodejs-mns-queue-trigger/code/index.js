'use strict';

exports.handler = function (event, context, callback) {
  // 解析JSON格式事件
  var eventObj = JSON.parse(event.toString());
  // 计算事件触发耗时
  var aliyunpublishtime = eventObj['aliyunpublishtime']
  console.log("message publish time:", aliyunpublishtime)
  console.log("message trigger time cost:", (new Date() - new Date(aliyunpublishtime)) / 1000, 's');
  // 返回消息内容
  // 事件中的消息默认是base64编码的，需要进行解码
  // 若希望自动解码，可在 s.yaml 中设置 IsBase64Decode: true
  console.log("messageBody:", eventObj['data']['messageBody'])
  callback(null, "succ");
};