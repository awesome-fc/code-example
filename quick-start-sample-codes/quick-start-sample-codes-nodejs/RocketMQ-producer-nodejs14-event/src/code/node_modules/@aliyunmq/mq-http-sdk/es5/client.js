'use strict';

const assert = require('assert');

const debug = require('debug')('mq:client');
const httpx = require('httpx');
const kitx = require('kitx');

const {
  toXMLBuffer,
  parseXML,
  extract,
  getCanonicalizedMQHeaders,
  processMsgProperties
} = require('./helper');

/**
 * MQ的client,用于保存aliyun账号消息,以及发送http请求
 */
class MQClient {
  /**
   * MQClient构造函数
   * @param {string}  endpoint  MQ的HTTP接入地址
   * @param {string}  accessKeyId 阿里云账号的AK
   * @param {string}  accessKeySecret 阿里云账号的SK 
   * @param {string}  securityToken 阿里云RAM授权的STS TOKEN，可空
   *
   * @returns {MQClient}
   */
  constructor(endpoint, accessKeyId, accessKeySecret, securityToken) {
    assert(endpoint, '"endpoint" must be passed in');
    this.endpoint = endpoint;
    assert(accessKeyId, 'must pass in "accessKeyId"');
    this.accessKeyId = accessKeyId;
    assert(accessKeySecret, 'must pass in "accessKeySecret"');
    this.accessKeySecret = accessKeySecret;
    // security token
    this.securityToken = securityToken;
  }

  /**
   * 发送http请求
   * @param {string}  method HTTP的请求方法GET/PUT/POST/DELETE...
   * @param {string}  resource  HTTP请求URL的path
   * @param {string}  type  解析XML响应内容的元素名字
   * @param {string}  requestBody 请求的body
   * @param {object}  opts  额外请求的参数
   *
   * @returns {object}  
   * ```json
   * {
   *  code: 200,
   *  requestId: "xxxxxxxxxxxxxx",
   *  body: {A:1,B:2,C:3}
   * }
   * ```json
   */
  async request(method, resource, type, requestBody, opts = {}) {
    const url = `${this.endpoint}${resource}`;
    debug('url: %s', url);
    debug('method: %s', method);
    const headers = this.buildHeaders(method, requestBody, resource, opts.headers);
    debug('request headers: %j', headers);
    debug('request body: %s', requestBody.toString());
    const response = await httpx.request(url, Object.assign(opts, {
      method: method,
      headers: headers,
      data: requestBody
    }));

    debug('statusCode %s', response.statusCode);
    debug('response headers: %j', response.headers);
    const code = response.statusCode;

    const contentType = response.headers['content-type'] || '';
    const responseBody = await httpx.read(response, 'utf8');
    debug('response body: %s', responseBody);

    var body;
    if (responseBody && (contentType.startsWith('text/xml') || contentType.startsWith('application/xml'))) {
      const responseData = await parseXML(responseBody);

      if (responseData.Error) {
        const e = responseData.Error;
        const message = extract(e.Message);
        const requestid = extract(e.RequestId);
        //const hostid = extract(e.HostId);
        const errorcode = extract(e.Code);
        const err = new Error(`${method} ${url} failed with ${code}. ` + `RequestId: ${requestid}, ErrorCode: ${errorcode}, ErrorMsg: ${message}`);
        err.Code = errorcode;
        err.RequestId = requestid;
        throw err;
      }

      body = {};
      Object.keys(responseData[type]).forEach(key => {
        if (key !== '$') {
          body[key] = extract(responseData[type][key]);
        }
      });
    }

    return {
      code,
      requestId: response.headers['x-mq-request-id'],
      body: body
    };
  }

  /**
   * 发送HTTP GET请求
   *
   * @param {string}  resource  HTTP请求URL的path
   * @param {string}  type  解析XML响应内容的元素名字
   * @param {object}  opts  额外请求的参数
   *
   * @returns {object}  
   * ```json
   * {
   *  code: 200,
   *  requestId: "xxxxxxxxxxxxxx",
   *  body: {A:1,B:2,C:3}
   * }
   * ```
   */
  get(resource, type, opts) {
    return this.request('GET', resource, type, '', opts);
  }

  /**
   * 发送HTTP POST请求
   *
   * @param {string}  resource  HTTP请求URL的path
   * @param {string}  type  解析XML响应内容的元素名字
   * @param {string}  requestBody 请求的body
   * @returns {object}  
   * ```json
   * {
   *  code: 200,
   *  requestId: "xxxxxxxxxxxxxx",
   *  body: {A:1,B:2,C:3}
   * }
   * ```
   */
  post(resource, type, body) {
    return this.request('POST', resource, type, body);
  }

  /**
   * 发送HTTP DELETE请求
   *
   * @param {string}  resource  HTTP请求URL的path
   * @param {string}  type  解析XML响应内容的元素名字
   * @param {string}  requestBody 请求的body
   * @returns {object}  
   * ```json
   * {
   *  code: 200,
   *  requestId: "xxxxxxxxxxxxxx",
   *  body: {A:1,B:2,C:3}
   * }
   * ```
   */
  delete(resource, type, body) {
    return this.request('DELETE', resource, type, body);
  }

  /**
   * 对请求的内容按照MQ的HTTP协议签名,sha1+base64
   * @param {string}  method  请求方法 
   * @param {object}  headers 请求头
   * @param {string}  resource  HTTP请求URL的path
   *
   * @returns {string} 签名
   */
  sign(method, headers, resource) {
    const canonicalizedMQHeaders = getCanonicalizedMQHeaders(headers);
    const md5 = headers['content-md5'] || '';
    const date = headers['date'];
    const type = headers['content-type'] || '';

    var toSignString = `${method}\n${md5}\n${type}\n${date}\n${canonicalizedMQHeaders}${resource}`;
    var buff = Buffer.from(toSignString, 'utf8');
    const degist = kitx.sha1(buff, this.accessKeySecret, 'binary');
    return Buffer.from(degist, 'binary').toString('base64');
  }

  /**
   * 组装请求MQ需要的请求头
   * @param {string}  method  请求方法
   * @param {string}  body  请求内容
   * @param {string}  resource  HTTP请求URL的path
   *
   * @returns {object} headers
   */
  buildHeaders(method, body, resource) {
    const date = new Date().toGMTString();

    const headers = {
      'date': date,
      'x-mq-version': '2015-06-06',
      'content-type': 'text/xml;charset=utf-8',
      'user-agent': 'mq-nodejs-sdk/1.0.4'
    };

    if (method !== 'GET' && method !== 'HEAD') {
      const digest = kitx.md5(body, 'hex');
      const md5 = Buffer.from(digest, 'utf8').toString('base64');
      Object.assign(headers, {
        'content-length': body.length,
        'content-md5': md5
      });
    }

    const signature = this.sign(method, headers, resource);

    headers['authorization'] = `MQ ${this.accessKeyId}:${signature}`;

    if (this.securityToken) {
      headers['security-token'] = this.securityToken;
    }

    return headers;
  }

  /**
   * 构造一个MQ的消费者
   * @param {string}  instanceId 实例ID
   * @param {string}  topic 主题名字
   * @param {string}  consumer  消费者名字
   * @param {string}  messageTag  消费的过滤标签，可空
   *
   * @returns {MQConsumer}
   */
  getConsumer(instanceId, topic, consumer, messageTag) {
    // eslint-disable-next-line no-use-before-define
    return new MQConsumer(this, instanceId, topic, consumer, messageTag);
  }

  /**
   * 构造一个MQ的生产者
   * @param {string}  instanceId 实例ID
   * @param {string}  topic 主题名字
   *
   * @returns {MQProducer}
   */
  getProducer(instanceId, topic) {
    // eslint-disable-next-line no-use-before-define
    return new MQProducer(this, instanceId, topic);
  }

  /**
   * 构造一个MQ的事务消息生产者
   * @param {string}  instanceId 实例ID
   * @param {string}  topic 主题名字
   * @param {string}  groupId 客户端GroupID
   *
   * @returns {MQTransProducer}
   */
  getTransProducer(instanceId, topic, groupId) {
    // eslint-disable-next-line no-use-before-define
    return new MQTransProducer(this, instanceId, topic, groupId);
  }
}

/**
 * 消息属性
 */
class MessageProperties {

  constructor() {
    this.properties = {};
  }

  /**
   * 获取消息属性内部的Object
   *
   * @returns {Object}
   */
  getProperties() {
    return this.properties;
  }

  /**
   * 设置消息KEY
   * @param {string} key 消息KEY
   */
  messageKey(key) {
    if (key == null) {
      return;
    }
    this.properties["KEYS"] = key + "";
  }

  /**
   * 定时消息，单位毫秒（ms），在指定时间戳（当前时间之后）进行投递。
   * 如果被设置成当前时间戳之前的某个时刻，消息将立刻投递给消费者
   *
   * @param {Number} timeMillis 定时的绝对时间戳
   */
  startDeliverTime(timeMillis) {
    if (timeMillis == null) {
      return;
    }
    this.properties["__STARTDELIVERTIME"] = timeMillis + "";
  }

  /**
   * 在消息属性中添加第一次消息回查的最快时间，单位秒，并且表征这是一条事务消息
   *
   * @param {Number} timeSeconds 第一次消息回查时间，单位秒
   */
  transCheckImmunityTime(timeSeconds) {
    if (timeSeconds == null) {
      return;
    }
    this.properties["__TransCheckT"] = timeSeconds + "";
  }

  /**
   * 分区顺序消息中区分不同分区的关键字段，sharding key 于普通消息的 key 是完全不同的概念。
   * 全局顺序消息，该字段可以设置为任意非空字符串。
   * 
   * @param {string} key 分区键值
   */
  shardingKey(key) {
    if (key == null) {
      return;
    }
    this.properties["__SHARDINGKEY"] = key + "";
  }

  /**
   * 设置消息自定义属性
   *
   * @param {string} key 属性键,非空
   * @param {string} value 属性值,非空
   */
  putProperty(key, value) {
    if (key == null || value == null) {
      return;
    }
    const keyStr = key + "";
    const valueStr = value + "";
    if (keyStr != "" && valueStr != "") {
      this.check(keyStr);
      this.check(valueStr);
      this.properties[keyStr] = valueStr;
    }
  }

  check(key) {
    if (key.indexOf('\'') > -1 || key.indexOf('"') > -1 || key.indexOf('&') > -1 || key.indexOf('<') > -1 || key.indexOf('>') > -1 || key.indexOf('|') > -1 || key.indexOf(':') > -1) {
      throw new Error(`Property ${key} can not contains: \" \' & < > | :`);
    }
  }

}

/**
 * MQ的消息生产者
 */
class MQProducer {
  /**
   * 构造函数
   * @param {MQClient}  client  MQ的客户端
   * @param {string}  instanceId 实例ID
   * @param {string}  topic 主题名字
   *
   * @returns {MQProducer}
   */
  constructor(client, instanceId, topic) {
    assert(client, '"client" must be passed in');
    assert(topic, '"topic" must be passed in');
    this.client = client;
    this.instanceId = instanceId;
    this.topic = topic;
    if (instanceId && instanceId !== '') {
      this.path = `/topics/${topic}/messages?ns=${instanceId}`;
    } else {
      this.path = `/topics/${topic}/messages`;
    }
  }

  /**
   * 向主题发送一条消息
   * @param {string}  body  发送的内容
   * @param {string}  tag   发送消息的标签
   * @param {MessageProperties} msgProps 发送消息的属性
   *
   * @returns {object} 
   * ```json
   * {
   *  // http请求状态码，发送成功就是201，如果发送失败则抛异常
   *  code: 201,
   *  // 请求ID
   *  requestId: "xxxxxxxxxxxxxx",
   *  // 发送消息的响应内容
   *  body: {
   *    // 消息ID
   *    MessageId: "",
   *    // 消息体内容的MD5值
   *    MessageBodyMD5: ""
   *  }
   * }
   * ```
   * @throws {exception} err  MQ服务端返回的错误或者其它网络异常
   * ```json
   * {
   *  // MQ服务端返回的错误Code，like: TopicNotExist 
   *  Code:"",
   *  // 请求ID
   *  RequestId:""
   * }
   * ```
   */
  async publishMessage(body, tag, msgProps) {
    var xmlBody;
    var params = { MessageBody: body };
    if (tag && tag != '') {
      params.MessageTag = tag;
    }
    if (msgProps) {
      var props = msgProps.getProperties();
      var propKeys = Object.keys(props);
      if (propKeys.length > 0) {
        params.Properties = propKeys.map(key => `${key}:${props[key]}`).join('|');
      }
    }

    var response = this.client.post(this.path, 'Message', toXMLBuffer('Message', params));
    return response;
  }
}

/**
 * MQ的消息生产者，支持事务 
 */
class MQTransProducer {
  /**
   * 构造函数
   * @param {MQClient}  client  MQ的客户端
   * @param {string}  instanceId 实例ID
   * @param {string}  topic 主题名字
   * @param {string}  groupId 客户端GroupID
   *
   * @returns {MQTransProducer}
   */
  constructor(client, instanceId, topic, groupId) {
    assert(client, '"client" must be passed in');
    assert(topic, '"topic" must be passed in');
    assert(groupId, '"groupId" must be passed in');
    this.client = client;
    this.instanceId = instanceId;
    this.topic = topic;
    this.groupId = groupId;
    if (instanceId && instanceId !== '') {
      this.path = `/topics/${topic}/messages?ns=${instanceId}`;
      this.transPopPath = `/topics/${topic}/messages?consumer=${groupId}&ns=${instanceId}&trans=pop`;
      this.transOprPath = `/topics/${topic}/messages?consumer=${groupId}&ns=${instanceId}`;
    } else {
      this.path = `/topics/${topic}/messages`;
      this.transPopPath = `/topics/${topic}/messages?consumer=${groupId}&trans=pop`;
      this.transOprPath = `/topics/${topic}/messages?consumer=${groupId}`;
    }
  }

  /**
   * 向主题发送一条消息
   * @param {string}  body  发送的内容
   * @param {string}  tag   发送消息的标签
   * @param {MessageProperties} msgProps 发送消息的属性
   *
   * @returns {object} 
   * ```json
   * {
   *  // http请求状态码，发送成功就是201，如果发送失败则抛异常
   *  code: 201,
   *  // 请求ID
   *  requestId: "xxxxxxxxxxxxxx",
   *  // 发送消息的响应内容
   *  body: {
   *    // 消息ID
   *    MessageId: "",
   *    // 消息体内容的MD5值
   *    MessageBodyMD5: ""
   *    // 消息句柄，仅事务消息存在
   *    ReceiptHandle: ""
   *  }
   * }
   * ```
   * @throws {exception} err  MQ服务端返回的错误或者其它网络异常
   * ```json
   * {
   *  // MQ服务端返回的错误Code，like: TopicNotExist 
   *  Code:"",
   *  // 请求ID
   *  RequestId:""
   * }
   * ```
   */
  async publishMessage(body, tag, msgProps) {
    var xmlBody;
    var params = { MessageBody: body };
    if (tag && tag != '') {
      params.MessageTag = tag;
    }
    if (msgProps) {
      var props = msgProps.getProperties();
      var propKeys = Object.keys(props);
      if (propKeys.length > 0) {
        params.Properties = propKeys.map(key => `${key}:${props[key]}`).join('|');
      }
    }

    var response = this.client.post(this.path, 'Message', toXMLBuffer('Message', params));
    return response;
  }

  /**
   * 消费检查事务半消息,默认如果该条消息没有被 {commit} 或者 {rollback} 在NextConsumeTime时会再次消费到该条消息
   *
   * @param {int} numOfMessages 每次从服务端消费条消息
   * @param {int} waitSeconds 长轮询的等待时间（可空），如果服务端没有消息请求会在该时间之后返回等于请求阻塞在服务端，如果期间有消息立刻返回
   *
   * @returns {object}
   * ```json
   * {
   *  code: 200,
   *  requestId: "",
   *  body: [
   *    {
   *      // 消息ID
   *      MessageId: "",
   *      // 消息体MD5
   *      MessageBodyMD5: "",
   *      // 发送消息的时间戳，毫秒
   *      PublishTime: {long},
   *      // 下次重试消费的时间，前提是这次不调用{commit} 或者 {rollback}，毫秒
   *      NextConsumeTime: {long},
   *      // 第一次消费的时间，毫秒，对于顺序消费无意义
   *      FirstConsumeTime: {long},
   *      // 消费的次数
   *      ConsumedTimes: {long},
   *      // 消息句柄，调用 {commit} 或者 {rollback} 需要将消息句柄传入，用于提交或者回滚该条事务消息
   *      ReceiptHandle: "",
   *      // 消息内容
   *      MessageBody: "",
   *      // 消息标签
   *      MessageTag: ""
   *    }
   *  ]
   * }
   *
   * ```
   * @throws {exception} err  MQ服务端返回的错误或者其它网络异常
   * ```json
   *  {
   *    // MQ服务端返回的错误Code，其中MessageNotExist是正常现象，表示没有可消费的消息
   *    Code: "",
   *    // 请求ID
   *    RequestId: ""
   *  }
   * ```json
   */
  async consumeHalfMessage(numOfMessages, waitSeconds) {
    var url = this.transPopPath + `&numOfMessages=${numOfMessages}`;
    if (waitSeconds) {
      url += `&waitseconds=${waitSeconds}`;
    }

    const subType = 'Message';
    var response = await this.client.get(url, 'Messages', { timeout: 33000 });
    response.body = response.body[subType];
    response.body.forEach(msg => {
      processMsgProperties(msg);
    });
    return response;
  }

  /**
   * 提交事务消息
   *
   * @param {string} receiptHandle consumeHalfMessage返回的单条消息句柄或者是发送事务消息返回的句柄
   *
   * @returns {object}
   * ```json
   * {
   *  // 请求成功
   *  code:204,
   *  // 请求ID
   *  requestId:""
   * }
   * ```
   *
   * @throws {exception}  err 请求失败或者其它网络异常
   * ```json
   * {
   *  // MQ服务端返回的错误Code，如ReceiptHandleError，表示消息句柄非法，MessageNotExist如果超过了TransCheckImmunityTime（针对发送事务消息的句柄）或者超过NextCnosumeTime
   *  Code: ""
   *  // 请求ID
   *  RequestId: ""
   * }
   * ```
   */
  async commit(receiptHandle) {
    const body = toXMLBuffer('ReceiptHandles', [receiptHandle], 'ReceiptHandle');
    const response = await this.client.delete(this.transOprPath + '&trans=commit', 'Errors', body);
    // 3种情况，普通失败，部分失败，全部成功
    if (response.body) {
      const subType = 'Error';
      // 部分失败
      response.body = response.body[subType];
    }
    return response;
  }

  /**
   * 回滚事务消息
   *
   * @param {string} receiptHandle consumeHalfMessage返回的单条消息句柄或者是发送事务消息返回的句柄
   *
   * @returns {object}
   * ```json
   * {
   *  // 请求成功
   *  code:204,
   *  // 请求ID
   *  requestId:""
   * }
   * ```
   *
   * @throws {exception}  err 请求失败或者其它网络异常
   * ```json
   * {
   *  // MQ服务端返回的错误Code，如ReceiptHandleError，表示消息句柄非法，MessageNotExist如果超过了TransCheckImmunityTime（针对发送事务消息的句柄）或者超过NextCnosumeTime
   *  Code: ""
   *  // 请求ID
   *  RequestId: ""
   * }
   * ```
   */
  async rollback(receiptHandle) {
    const body = toXMLBuffer('ReceiptHandles', [receiptHandle], 'ReceiptHandle');
    const response = await this.client.delete(this.transOprPath + '&trans=rollback', 'Errors', body);
    // 3种情况，普通失败，部分失败，全部成功
    if (response.body) {
      const subType = 'Error';
      // 部分失败
      response.body = response.body[subType];
    }
    return response;
  }
}

/**
 * MQ的消息消费者
 */
class MQConsumer {

  /**
   * 消费者构造函数
   * @param {MQClient}  client MQ客户端
   * @param {string}  instanceId 实例ID
   * @param {string}  topic  主题名字
   * @param {string}  consumer 消费者名字(CID)a
   * @param {string}  messageTag  消费消息的过滤标签，可空 
   *
   * @returns {MQConsumer}
   */
  constructor(client, instanceId, topic, consumer, messageTag) {
    assert(client, '"client" must be passed in');
    assert(topic, '"topic" must be passed in');
    assert(consumer, '"consumer" must be passed in');
    this.client = client;
    this.instanceId = instanceId;
    this.topic = topic;
    this.consumer = consumer;
    if (messageTag) {
      this.messageTag = encodeURIComponent(messageTag);
    }
    if (instanceId && instanceId !== '') {
      if (this.messageTag) {
        this.path = `/topics/${topic}/messages?consumer=${consumer}&ns=${instanceId}&tag=${this.messageTag}`;
      } else {
        this.path = `/topics/${topic}/messages?consumer=${consumer}&ns=${instanceId}`;
      }
      this.ackPath = `/topics/${topic}/messages?consumer=${consumer}&ns=${instanceId}`;
    } else {
      if (this.messageTag) {
        this.path = `/topics/${topic}/messages?consumer=${consumer}&tag=${this.messageTag}`;
      } else {
        this.path = `/topics/${topic}/messages?consumer=${consumer}`;
      }
      this.ackPath = `/topics/${topic}/messages?consumer=${consumer}`;
    }
  }

  /**
   * 消费消息,默认如果该条消息没有被 {ackMessage} 确认消费成功，即在NextConsumeTime时会再次消费到该条消息
   *
   * @param {int} numOfMessages 每次从服务端消费条消息
   * @param {int} waitSeconds 长轮询的等待时间（可空），如果服务端没有消息请求会在该时间之后返回等于请求阻塞在服务端，如果期间有消息立刻返回
   *
   * @returns {object}
   * ```json
   * {
   *  code: 200,
   *  requestId: "",
   *  body: [
   *    {
   *      // 消息ID
   *      MessageId: "",
   *      // 消息体MD5
   *      MessageBodyMD5: "",
   *      // 发送消息的时间戳，毫秒
   *      PublishTime: {long},
   *      // 下次重试消费的时间，前提是这次不调用{ackMessage} 确认消费消费成功，毫秒
   *      NextConsumeTime: {long},
   *      // 第一次消费的时间，毫秒
   *      FirstConsumeTime: {long},
   *      // 消费的次数
   *      ConsumedTimes: {long},
   *      // 消息句柄，调用 {ackMessage} 需要将消息句柄传入，用于确认该条消息消费成功
   *      ReceiptHandle: "",
   *      // 消息内容
   *      MessageBody: "",
   *      // 消息标签
   *      MessageTag: ""
   *    }
   *  ]
   * }
   *
   * ```
   * @throws {exception} err  MQ服务端返回的错误或者其它网络异常
   * ```json
   *  {
   *    // MQ服务端返回的错误Code，其中MessageNotExist是正常现象，表示没有可消费的消息
   *    Code: "",
   *    // 请求ID
   *    RequestId: ""
   *  }
   * ```json
   */
  async consumeMessage(numOfMessages, waitSeconds) {
    var url = this.path + `&numOfMessages=${numOfMessages}`;
    if (waitSeconds) {
      url += `&waitseconds=${waitSeconds}`;
    }

    const subType = 'Message';
    var response = await this.client.get(url, 'Messages', { timeout: 33000 });
    response.body = response.body[subType];
    response.body.forEach(msg => {
      processMsgProperties(msg);
    });
    return response;
  }

  /**
   * 顺序消费消息,拿到的消息可能是多个分区的（对于分区顺序）一个分区的内的消息一定是顺序的
   * 对于顺序消费，如果一个分区内的消息只要有没有被确认消费 {ackMessage} 成功，则对于这个分区在NextConsumeTime后还会消费到相同的消息
   * 对于一个分区，只有所有消息确认消费成功才能消费下一批消息
   *
   * @param {int} numOfMessages 每次从服务端消费条消息
   * @param {int} waitSeconds 长轮询的等待时间（可空），如果服务端没有消息请求会在该时间之后返回等于请求阻塞在服务端，如果期间有消息立刻返回
   *
   * @returns {object}
   * ```json
   * {
   *  code: 200,
   *  requestId: "",
   *  body: [
   *    {
   *      // 消息ID
   *      MessageId: "",
   *      // 消息体MD5
   *      MessageBodyMD5: "",
   *      // 发送消息的时间戳，毫秒
   *      PublishTime: {long},
   *      // 下次重试消费的时间，前提是这次不调用{ackMessage} 确认消费消费成功，毫秒
   *      NextConsumeTime: {long},
   *      // 第一次消费的时间，毫秒，顺序消费无意义
   *      FirstConsumeTime: {long},
   *      // 消费的次数
   *      ConsumedTimes: {long},
   *      // 消息句柄，调用 {ackMessage} 需要将消息句柄传入，用于确认该条消息消费成功
   *      ReceiptHandle: "",
   *      // 消息内容
   *      MessageBody: "",
   *      // 消息标签
   *      MessageTag: ""
   *    }
   *  ]
   * }
   *
   * ```
   * @throws {exception} err  MQ服务端返回的错误或者其它网络异常
   * ```json
   *  {
   *    // MQ服务端返回的错误Code，其中MessageNotExist是正常现象，表示没有可消费的消息
   *    Code: "",
   *    // 请求ID
   *    RequestId: ""
   *  }
   * ```json
   */
  async consumeMessageOrderly(numOfMessages, waitSeconds) {
    var url = this.path + `&numOfMessages=${numOfMessages}`;
    if (waitSeconds) {
      url += `&waitseconds=${waitSeconds}`;
    }
    url += `&trans=order`;

    const subType = 'Message';
    var response = await this.client.get(url, 'Messages', { timeout: 33000 });
    response.body = response.body[subType];
    response.body.forEach(msg => {
      processMsgProperties(msg);
    });
    return response;
  }

  /**
   * 确认消息消费成功，消费成功后需要调用该接口否则会重复消费消息
   *
   * @param {array} receiptHandles 消息句柄数组
   *
   * @returns {object}
   * ```json
   * {
   *  // 请求成功
   *  code:204,
   *  // 请求ID
   *  requestId:""
   * }
   * ```
   *
   * @throws {exception}  err 请求失败或者其它网络异常
   * ```json
   * {
   *  // MQ服务端返回的错误Code，如ReceiptHandleError，表示消息句柄非法，MessageNotExist表示超过了ack的时间，即NextConsumeTime
   *  Code: ""
   *  // 请求ID
   *  RequestId: ""
   * }
   * ```
   */
  async ackMessage(receiptHandles) {
    const body = toXMLBuffer('ReceiptHandles', receiptHandles, 'ReceiptHandle');
    const response = await this.client.delete(this.ackPath, 'Errors', body);
    // 3种情况，普通失败，部分失败，全部成功
    if (response.body) {
      const subType = 'Error';
      // 部分失败
      response.body = response.body[subType];
    }
    return response;
  }

}

module.exports = {
  MQClient,
  MQConsumer,
  MQProducer,
  MessageProperties
};