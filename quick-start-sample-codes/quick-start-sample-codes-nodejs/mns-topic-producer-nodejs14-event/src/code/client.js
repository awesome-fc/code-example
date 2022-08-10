'use strict';
const assert = require('assert');
const httpx = require('httpx');
const kitx = require('kitx');
const xml2js = require('xml2js');
/**
 * MNS的client,用于保存aliyun账号消息,以及发送http请求
 */
class MNSClient {
  /**
   * MNSClient构造函数
   * @param {string}  endpoint  MNSClient的HTTP接入地址
   * @param {string}  accessKeyId 阿里云账号的AK
   * @param {string}  accessKeySecret 阿里云账号的SK 
   * @param {string}  securityToken 阿里云RAM授权的STS TOKEN，可空
   *
   * @returns {MNSClient}
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
    console.log('url: %s', url);
    console.log('method: %s', method);
    const headers = this.buildHeaders(method, requestBody, resource, opts.headers);
    console.log('request headers: %j', headers);
    console.log('request body: %s', requestBody.toString());
    const response = await httpx.request(url, Object.assign(opts, {
      method: method,
      headers: headers,
      data: requestBody
    }));

    console.log('statusCode %s', response.statusCode);
    console.log('response headers: %j', response.headers);
    const code = response.statusCode;

    const contentType = response.headers['content-type'] || '';
    const responseBody = await httpx.read(response, 'utf8');
    console.log('response body: %s', responseBody);

    var body;
    if (responseBody && (contentType.startsWith('text/xml') || contentType.startsWith('application/xml'))) {
      const responseData = await parseXML(responseBody);

      if (responseData.Error) {
        const e = responseData.Error;
        const message = extract(e.Message);
        const requestid = extract(e.RequestId);
        const errorcode = extract(e.Code);
        const err = new Error(`${method} ${url} failed with ${code}. ` +
          `RequestId: ${requestid}, ErrorCode: ${errorcode}, ErrorMsg: ${message}`);
        err.Code = errorcode;
        err.RequestId = requestid;
        throw err;
      }

      body = {};
      Object.keys(responseData[type]).forEach((key) => {
        if (key !== '$') {
          body[key] = extract(responseData[type][key]);
        }
      });
    }

    return {
      code,
      requestId: response.headers['x-mns-request-id'],
      body: body
    };
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
  sign(method, headers, resource) {
    const canonicalizedMNSHeaders = getCanonicalizedMNSHeaders(headers);
    const md5 = headers['content-md5'] || '';
    const date = headers['date'];
    const type = headers['content-type'] || '';

    var toSignString = `${method}\n${md5}\n${type}\n${date}\n${canonicalizedMNSHeaders}${resource}`;
    var buff = Buffer.from(toSignString, 'utf8');
    const degist = kitx.sha1(buff, this.accessKeySecret, 'binary');
    return Buffer.from(degist, 'binary').toString('base64');
  }

  /**
   * 组装请求MNS需要的请求头
   * @param {string}  method  请求方法
   * @param {string}  body  请求内容
   * @param {string}  resource  HTTP请求URL的path
   *
   * @returns {object} headers
   */
  buildHeaders(method, body, resource) {
    const date = new Date().toGMTString();
    //const date = "Wed, 27 Jul 2022 02:38:05 GMT" //固定一个，方便调试
    const headers = {
      'date': date,
      'x-mns-version': '2015-06-06',
      'content-type': 'application/xml;charset=utf-8',
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

    headers['authorization'] = `MNS ${this.accessKeyId}:${signature}`;

    if (this.securityToken) {
      headers['security-token'] = this.securityToken;
    }

    return headers;
  }
  /**
   * 构造一个MNS topic 
   * @param {string}  topic 主题名字
   *
   * @returns {MNSTopic}
   */
  getMNSTopic(topic) {
    return new MNSTopic(this, topic);
  }
  /**
   * 构造一个MNS queue
   * @param {string}  queue 队列名字
   *
   * @returns {MNSTopic}
   */
  getMNSQueue(queue) {
    return new MNSQueue(this, queue);
  }


}
/**
 * MNS的Topic
 */
class MNSQueue {
  /**
   * 构造函数
   * @param {MNSClient}  client  MNS的客户端
   * @param {string}  queue 队列名字
   *
   * @returns {MNSQueue}
   */
  constructor(client, queue) {
    assert(client, '"client" must be passed in');
    assert(queue, '"queue" must be passed in');
    this.client = client;
    this.queue = queue;
    this.path = `/queues/${queue}/messages`; //队列发送消息路由
  }

  /**
   * 向主题发送一条消息
   * @param {string}  body  发送的内容
   * @param {string}  timer  消息定时时间  
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
   * @throws {exception} err  MNS服务端返回的错误或者其它网络异常
   * ```json
   * {
   *  Code:"",
   *  // 请求ID
   *  RequestId:""
   * }
   * ```
   */
  async sendMessage(body, timer) {
    var params = { MessageBody: body };
    if (timer && timer != '') {  
      params.DelaySeconds = timer;
    }

    var response = this.client.post(this.path, 'Message', toXMLBuffer('Message', params));
    return response;
  }
}
/**
 * MNS的Topic
 */
class MNSTopic {
  /**
   * 构造函数
   * @param {MNSClient}  client  MNS的客户端
   * @param {string}  topic 主题名字
   *
   * @returns {MNSTopic}
   */
  constructor(client, topic) {
    assert(client, '"client" must be passed in');
    assert(topic, '"topic" must be passed in');
    this.client = client;
    this.topic = topic;
    this.path = `/topics/${topic}/messages`; //主题发送消息路由
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
   * @throws {exception} err  MNS服务端返回的错误或者其它网络异常
   * ```json
   * {
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
        params.Properties = propKeys
          .map((key) => `${key}:${props[key]}`).join('|');
      }
    }

    var response = this.client.post(this.path, 'Message', toXMLBuffer('Message', params));
    return response;
  }
}
function parseXML(input) {
  return new Promise((resolve, reject) => {
    xml2js.parseString(input, (err, obj) => {
      if (err) {
        return reject(err);
      }
      resolve(obj);
    });
  });
};

function extract(arr) {
  if (arr && arr.length === 1 && typeof arr[0] === 'string') {
    return arr[0];
  }

  arr.forEach((item) => {
    Object.keys(item).forEach((key) => {
      item[key] = extract(item[key]);
    });
  });

  return arr;
};

function format(params) {
  if (typeof params === 'string') {
    return params;
  }

  var xml = '';
  Object.keys(params).forEach((key) => {
    const value = params[key];
    if (typeof value === 'object') {
      xml += `<${key}>${format(value)}</${key}>`;
    } else {
      xml += `<${key}>${value}</${key}>`;
    }
  });
  return xml;
}

function toXMLBuffer(entityType, params, subType) {
  var xml = '<?xml version="1.0" encoding="UTF-8"?>';
  xml += `<${entityType} xmlns="http://mns.aliyuncs.com/doc/v1/">`;
  if (Array.isArray(params)) {
    params.forEach((item) => {
      xml += `<${subType}>`;
      xml += format(item);
      xml += `</${subType}>`;
    });
  } else {
    xml += format(params);
  }
  xml += `</${entityType}>`;
  return Buffer.from(xml, 'utf8');
};

function getCanonicalizedMNSHeaders(headers) {
  return Object.keys(headers)
    .filter((key) => key.startsWith('x-mns-'))
    .sort()
    .map((key) => `${key}:${headers[key]}\n`)
    .join('');
};

module.exports = {
  MNSClient,
  MNSTopic,
  MNSQueue
};
