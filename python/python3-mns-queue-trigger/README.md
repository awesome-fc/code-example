# Python3 消息服务 MNS 队列触发器示例

本示例为您展示了 python runtime 的 [消息服务MNS](https://help.aliyun.com/document_detail/27414.html) 队列模型触发器示例。
本示例使用了 MNS 的队列模型作为示例，与示例  python3-mns-queue-producer 一起实现了消息服务的生产者-消费者模型。
MNS的配置在函数的环境变量配置中（参考s.yaml)。

## 准备开始
- 一个可用的mns队列，可参考MNS官方文档[队列模型快速入门-创建队列](https://help.aliyun.com/document_detail/34417.html) 创建。
- [开通事件总线EventBridge并授权](https://help.aliyun.com/document_detail/182246.html)
- 有 MNS 权限的 RAM 用户
  - 建议直接使用函数计算默认的角色 AliyunFCDefaultRole
  - 也可参考MNS官方文档[开通消息服务MNS并授权](https://help.aliyun.com/document_detail/27423.html)，函数计算需要该RAM密钥访问MNS队列。
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始

### 方式一、使用控制台创建

#### 1. 编译打包

```shell
# 打包文件
cd code && zip -r python3-mns-queue-trigger.zip *
```

#### 2. 创建函数
选择服务（或创建服务）后，单击创建函数
- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 Python 3.6/3.9
- 选择函数触发方式：通过事件请求触发
- 其他设置使用默认

> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)

#### 3. 使用模拟数据测试
- 3.1 Json类型的测试事件
```json
{
  "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",
  "source":"MNS-Function-mnstrigger",
  "specversion":"1.0",
  "type":"mns:Queue:SendMessage",
  "datacontenttype":"application/json; charset=utf-8",
  "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",
  "time":"2021-04-08T06:28:17.093Z",
  "aliyunaccountid":"16490154********",
  "aliyunpublishtime":"2021-10-15T07:06:34.028Z",
  "aliyunoriginalaccountid":"164901546557****",
  "aliyuneventbusname":"MNS-Function-mnstrigger",
  "aliyunregionid":"cn-chengdu",
  "aliyunpublishaddr":"42.120.XX.XX",
  "data":{
      "requestId":"606EA3074344430D4C81****",
      "messageId":"C6DB60D1574661357FA227277445****",
      "messageBody":"TEST"
  }
}
```
测试返回结果如下所示：
```bash
mns_queue trigger event = b'{\n  "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",\n  "source":"MNS-Function-mnstrigger",\n  "specversion":"1.0",\n  "type":"mns:Queue:SendMessage",\n  "datacontenttype":"application/json; charset=utf-8",\n  "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",\n  "time":"2021-04-08T06:28:17.093Z",\n  "aliyunaccountid":"16490154********",\n  "aliyunpublishtime":"2021-10-15T07:06:34.028Z",\n  "aliyunoriginalaccountid":"164901546557****",\n  "aliyuneventbusname":"MNS-Function-mnstrigger",\n  "aliyunregionid":"cn-chengdu",\n  "aliyunpublishaddr":"42.120.XX.XX",\n  "data":{\n      "requestId":"606EA3074344430D4C81****",\n      "messageId":"C6DB60D1574661357FA227277445****",\n      "messageBody":"TEST"\n  }\n}'
```

返回日志如下所示
```bash
2022-07-21 19:03:18FunctionCompute python3 runtime inited.
2022-07-21 19:03:18FC Invoke Start RequestId: d39d5414-fb38-44cf-a371-de73********
2022-07-21 19:03:182022-07-21 19:03:18 d39d5414-fb38-44cf-a371-de73******** [INFO] b'{\n  "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",\n  "source":"MNS-Function-mnstrigger",\n  "specversion":"1.0",\n  "type":"mns:Queue:SendMessage",\n  "datacontenttype":"application/json; charset=utf-8",\n  "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",\n  "time":"2021-04-08T06:28:17.093Z",\n  "aliyunaccountid":"16490154********",\n  "aliyunpublishtime":"2021-10-15T07:06:34.028Z",\n  "aliyunoriginalaccountid":"164901546557****",\n  "aliyuneventbusname":"MNS-Function-mnstrigger",\n  "aliyunregionid":"cn-chengdu",\n  "aliyunpublishaddr":"42.120.XX.XX",\n  "data":{\n      "requestId":"606EA3074344430D4C81****",\n      "messageId":"C6DB60D1574661357FA227277445****",\n      "messageBody":"TEST"\n  }\n}'
2022-07-21 19:03:18FC Invoke End RequestId: d39d5414-fb38-44cf-a371-de73********
```

- 3.2 Stream类型的测试事件
```bash
mock mns message
```

测试返回结果如下所示：
```bash
mns_queue trigger event = b'mock mns message'
```

返回日志如下所示
```bash
2022-07-21 19:07:42FC Invoke Start RequestId: b89fcd08-e684-4fa2-bb80-8305********
2022-07-21 19:07:422022-07-21 19:07:42 b89fcd08-e684-4fa2-bb80-8305******** [INFO] b'mock mns message'
2022-07-21 19:07:42FC Invoke End RequestId: b89fcd08-e684-4fa2-bb80-8305********
```

#### 4. 配置MNS触发器
选择 queue 模型 MNS 触发器

<img src="assets/20220720102639.jpg" alt="img_1.png" style="zoom: 40%;" />

#### 5. 通过MNS控制台触发测试函数

![img_2.png](assets/20220720104405.jpg)

测试返回结果如下所示：
```bash
mns_topic trigger event = b'{"id":"0FFF111D2A6B444B7FE46AE2xxxxxxxx","source":"MNS-python3-mns-queue-trigger-trigger-5h3jxxxx","specversion":"1.0","type":"mns:Queue:SendMessage","datacontenttype":"application/json;charset=utf-8","subject":"acs:mns:cn-shenzhen:15812231xxxxxxxx:queues/fc-example","time":"2022-07-21T10:07:31.525Z","aliyunaccountid":"15812231xxxxxxxx","aliyunpublishtime":"2022-07-21T10:07:32.019Z","aliyunoriginalaccountid":"15812231xxxxxxxx","aliyuneventbusname":"MNS-python3-mns-queue-trigger-trigger-5h3jxxxx","aliyunregionid":"cn-shenzhen","aliyunpublishaddr":"10.58.xx.xx","data":{"requestId":"62D92563354133CAxxxxxxxx","messageId":"0FFF111D2A6B444B7FE46AE2xxxxxxxx","messageBody":"bWVlc2FnZSBmcm9tIE1OUyBjb25zb2xl"}}'
```

在函数计算控制台查看请求日志，如下所示：
```bash
2022-07-21 18:07:32FunctionCompute python3 runtime inited.
2022-07-21 18:07:32FC Invoke Start RequestId: 0FFF111D2A6B444B7FE46AE2xxxxxxxx
2022-07-21 18:07:322022-07-21 18:07:32 0FFF111D2A6B444B7FE46AE2xxxxxxxx [INFO] b'{"id":"0FFF111D2A6B444B7FE46AE2xxxxxxxx","source":"MNS-python3-mns-queue-trigger-trigger-5h3jxxxx","specversion":"1.0","type":"mns:Queue:SendMessage","datacontenttype":"application/json;charset=utf-8","subject":"acs:mns:cn-shenzhen:15812231xxxxxxxx:queues/fc-example","time":"2022-07-21T10:07:31.525Z","aliyunaccountid":"15812231xxxxxxxx","aliyunpublishtime":"2022-07-21T10:07:32.019Z","aliyunoriginalaccountid":"15812231xxxxxxxx","aliyuneventbusname":"MNS-python3-mns-queue-trigger-trigger-5h3jxxxx","aliyunregionid":"cn-shenzhen","aliyunpublishaddr":"10.58.xx.xx","data":{"requestId":"62D92563354133CAxxxxxxxx","messageId":"0FFF111D2A6B444B7FE46AE2xxxxxxxx","messageBody":"bWVlc2FnZSBmcm9tIE1OUyBjb25zb2xl"}}'
2022-07-21 18:07:32FC Invoke End RequestId: 0FFF111D2A6B444B7FE46AE2xxxxxxxx
```

### 方式二、使用 Serverless Devs 工具编译部署
该方式使用模拟数据进行调用测试

#### 1. 修改 s.yaml 配置

[Yaml完整配置](https://github.com/devsapp/fc/tree/main/docs/zh/yaml)

- 根据需要修改 access 配置

- 添加 mns 队列触发器配置

  ```yaml
      triggers:  
        - name: eventbridgeTriggerWithMNSSource       
        type: eventbridge        
        config:      
          triggerEnable: true
          asyncInvocationType: false
          eventRuleFilterPattern: '{"source":["MNS-${functionName}-eventbridgeTriggerWithMNSSource"]}'
          eventSourceConfig:
            eventSourceType: MNS
            eventSourceParameters:
              sourceMNSParameters:
                QueueName: fc-example
                IsBase64Decode: false
  ```

#### 2. 安装依赖并部署

编译部署代码包
```shell
s deploy
```
#### 3. 使用模拟数据测试

```shell
s invoke --event-file event.json
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
FunctionCompute python3 runtime inited.
FC Invoke Start RequestId: 30ee4b56-4ab8-4152-8680-bb3098fadb01
2022-07-21T10:57:12.346Z 30ee4b56-4ab8-4152-8680-bb3098fadb01 [INFO] b'{  "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",  "source":"MNS-Function-mnstrigger",  "specversion":"1.0",  "type":"mns:Queue:SendMessage",  "datacontenttype":"application/json; charset=utf-8",  "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",  "time":"2021-04-08T06:28:17.093Z",  "aliyunaccountid":"1649015465574023",  "aliyunpublishtime":"2021-10-15T07:06:34.028Z",  "aliyunoriginalaccountid":"164901546557****",  "aliyuneventbusname":"MNS-Function-mnstrigger",  "aliyunregionid":"cn-chengdu",  "aliyunpublishaddr":"42.120.XX.XX",  "data":{      "requestId":"606EA3074344430D4C81****",      "messageId":"C6DB60D1574661357FA227277445****",      "messageBody":"TEST"  }}'
FC Invoke End RequestId: 30ee4b56-4ab8-4152-8680-bb3098fadb01
Duration: 2.60 ms, Billed Duration: 3 ms, Memory Size: 128 MB, Max Memory Used: 25.02 MB
========= FC invoke Logs end =========
FC Invoke instanceId: c-62d93108-3ad49d12bf074baf8549
FC Invoke Result:
mns_queue trigger event = b'{  "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",  "source":"MNS-Function-mnstrigger",  "specversion":"1.0",  "type":"mns:Queue:SendMessage",  "datacontenttype":"application/json; charset=utf-8",  "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",  "time":"2021-04-08T06:28:17.093Z",  "aliyunaccountid":"1649015465574023",  "aliyunpublishtime":"2021-10-15T07:06:34.028Z",  "aliyunoriginalaccountid":"164901546557****",  "aliyuneventbusname":"MNS-Function-mnstrigger",  "aliyunregionid":"cn-chengdu",  "aliyunpublishaddr":"42.120.XX.XX",  "data":{      "requestId":"606EA3074344430D4C81****",      "messageId":"C6DB60D1574661357FA227277445****",      "messageBody":"TEST"  }}'


End of method: invoke
```

## 注意事项
1. MNS消息服务和函数计算建议部署在同一个地域