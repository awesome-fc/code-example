# MQ Nodejs HTTP SDK

Alyun MQ Documents: http://www.aliyun.com/product/ons

Aliyun MQ Console: https://ons.console.aliyun.com

## Installation

Add dependency `@aliyunmq/mq-http-sdk`, get the latest version from [npmjs](https://www.npmjs.com/)
```bash
npm install --save
```
*Note: nodejs >= 7.6.0*
## docs

[Documents](https://aliyunmq.github.io/mq-http-nodejs-sdk/)

## Note
1. Http consumer only support timer msg (less than 3 days), no matter the msg is produced from http or tcp protocol.
2. Order is only supported at special server cluster.

## Sample (github)

[Publish Message](https://github.com/aliyunmq/mq-http-samples/blob/master/nodejs/producer.js)

[Consume Message](https://github.com/aliyunmq/mq-http-samples/blob/master/nodejs/consumer.js)

[Transaction Message](https://github.com/aliyunmq/mq-http-samples/blob/master/nodejs/trans-producer.js)

[Publish Order Message](https://github.com/aliyunmq/mq-http-samples/blob/master/nodejs/order-producer.js)

[Consume Order Message](https://github.com/aliyunmq/mq-http-samples/blob/master/nodejs/order-consumer.js)

## Sample (code.aliyun.com)

[Publish Message](https://code.aliyun.com/aliware_rocketmq/mq-http-samples/blob/master/nodejs/producer.js)

[Consume Message](https://code.aliyun.com/aliware_rocketmq/mq-http-samples/blob/master/nodejs/consumer.js)

[Transaction Message](https://code.aliyun.com/aliware_rocketmq/mq-http-samples/blob/master/nodejs/trans-producer.js)

[Publish Order Message](https://code.aliyun.com/aliware_rocketmq/mq-http-samples/blob/master/nodejs/order-producer.js)

[Consume Order Message](https://code.aliyun.com/aliware_rocketmq/mq-http-samples/blob/master/nodejs/order-consumer.js)