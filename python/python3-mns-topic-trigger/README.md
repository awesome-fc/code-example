# Python3 消息服务 MNS 触发器示例

本示例为您展示了 python runtime 的 [消息服务MNS](https://help.aliyun.com/document_detail/27414.html) 主题模型触发器示例。
本示例使用了 MNS 的主题模型作为示例，与示例  python3-mns-topic-producer 一起实现了消息服务的生产者-消费者模型。
MNS的配置在函数的环境变量配置中（参考s.yaml)。

## 准备开始
- 一个可用的mns主题，可参考MNS官方文档[主题模型快速入门-创建主题](https://help.aliyun.com/document_detail/34424.html) 创建。
- 有 MNS 权限的 RAM 用户
  - 也可参考MNS官方文档[开通消息服务MNS并授权](https://help.aliyun.com/document_detail/27423.html)，函数计算需要该RAM密钥访问MNS主题。
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始

### 方式一、使用控制台创建

#### 1. 编译打包

```shell
# 打包文件
cd code && zip -r python3-mns-topic-trigger.zip *
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
  "TopicOwner":"topic account id",
  "Message":"mock mns message",
  "Subscriber":"subscriber account id",
  "PublishTime":1658235558094,
  "SubscriptionName":"test-5bf13c7e",
  "MessageMD5":"652BF0E6297840015247C3xxxxxxx",
  "TopicName":"fc-example",
  "MessageId":"3405CA51807661353B3xxxxxxxx"
}
```
测试返回结果如下所示：
```bash
mns_topic trigger event = b'{  "TopicOwner":"topic account id",  "Message":"mock mns message",  "Subscriber":"subscriber account id",  "PublishTime":1658235558094,  "SubscriptionName":"test-5bf13c7e",  "MessageMD5":"652BF0E6297840015247C3xxxxxxx",  "TopicName":"fc-example",  "MessageId":"3405CA51807661353B3xxxxxxxx"}
```

返回日志如下所示
```bash
2022-07-20 15:50:15FunctionCompute python3 runtime inited.
2022-07-20 15:50:15FC Invoke Start RequestId: e677e5aa-771b-404b-9d98-78e9064a5ea6
2022-07-20 15:50:152022-07-20 15:50:15 e677e5aa-771b-404b-9d98-78e9xxxxxxxx [INFO] mns_topic trigger event = b'{\n  "TopicOwner":"topic account id",\n  "Message":"mock mns message",\n  "Subscriber":"subscriber account id",\n  "PublishTime":1658235558094,\n  "SubscriptionName":"test-5bf13c7e",\n  "MessageMD5":"652BF0E6297840015247C3xxxxxxx",\n  "TopicName":"fc-example",\n  "MessageId":"3405CA51807661353B3xxxxxxxx"\n}'
2022-07-20 15:50:15FC Invoke End RequestId: e677e5aa-771b-404b-9d98-78e9xxxxxxxx
```

- 3.2 Stream类型的测试事件
```bash
mock mns message
```

测试返回结果如下所示：
```bash
mns_topic trigger event = b'mock mns message'
```

返回日志如下所示
```bash
2022-07-20 15:59:49FunctionCompute python3 runtime inited.
2022-07-20 15:59:50FC Invoke Start RequestId: 2f3583e6-94e2-479e-b297-8501xxxxxxxx
2022-07-20 15:59:502022-07-20 15:59:50 2f3583e6-94e2-479e-b297-8501xxxxxxxx [INFO] b'mock mns message'
2022-07-20 15:59:50FC Invoke End RequestId: 2f3583e6-94e2-479e-b297-8501xxxxxxxx
```

#### 4. 配置MNS触发器
- 选择 topic 模型 MNS 触发器
- 选择 JSON 的 Event 格式

> 注意：若选择 STREAM 的 Event 格式，在代码中则不需要将 Event 解析成json。

<img src="assets/20220720102639.jpg" alt="img_1.png" style="zoom: 40%;" />

> 在创建触发器过程中，页面会提示授权相关信息并让你创建 aliyunmnsnotificationrole  角色，按照页面指导创建即可。

#### 5. 通过MNS控制台触发测试函数

![img_2.png](assets/20220720104405.jpg)

测试返回结果如下所示：
```bash
mns_topic trigger event = b'{"TopicOwner":"15812231xxxxxxxx","Message":"json meesage from MNS console","Subscriber":"15812231xxxxxxxx","PublishTime":1658306182000,"SubscriptionName":"ree-2dfbf25f","MessageMD5":"7F1120AE2C15B843365E9A88xxxxxxxx","TopicName":"fc-example","MessageId":"20A37C322A6B444C1B386569xxxxxxxx"}'
```

在函数计算控制台查看请求日志，如下所示：
```bash
2022-07-20 16:36:22FunctionCompute python3 runtime inited.
2022-07-20 16:36:22FC Invoke Start RequestId: 03D74DFB-13A7-5061-9D4C-FCD1xxxxxxxx
2022-07-20 16:36:222022-07-20 16:36:22 03D74DFB-13A7-5061-9D4C-FCD1xxxxxxxx [INFO] b'{"TopicOwner":"15812231xxxxxxxx","Message":"json meesage from MNS console","Subscriber":"15812231xxxxxxxx","PublishTime":1658306182000,"SubscriptionName":"ree-2dfbf25f","MessageMD5":"7F1120AE2C15B843365E9A88xxxxxxxx","TopicName":"fc-example","MessageId":"20A37C322A6B444C1B386569xxxxxxxx"}'
2022-07-20 16:36:22FC Invoke End RequestId: 03D74DFB-13A7-5061-9D4C-FCD1xxxxxxxx
```

### 方式二、使用 Serverless Devs 工具编译部署
该方式使用模拟数据进行调用测试

#### 1. 修改 s.yaml 配置

[Yaml完整配置](https://gitee.com/devsapp/fc/blob/main/docs/zh/yaml/readme.md#yaml%E5%AE%8C%E6%95%B4%E9%85%8D%E7%BD%AE)

- 根据需要修改 access 配置

- 添加 [mns 触发器](https://gitee.com/devsapp/fc/blob/main/docs/zh/yaml/triggers.md#mns%E8%A7%A6%E5%8F%91%E5%99%A8) 配置

  ```yaml
        triggers:
          - name: {TriggerName}
            description: ''
            sourceArn: acs:mns:{Region}:{AccountID}:/topics/{TopicName}
            type: mns_topic
            role: acs:ram::{AccountID}:role/aliyunfcdefaultrole
            qualifier: LATEST
            config:
              filterTag: ''
              notifyContentFormat: JSON
              notifyStrategy: BACKOFF_RETRY
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
FC Invoke Start RequestId: f6c84568-6251-44cf-8f2d-715exxxxxxxx
2022-07-20T07:58:17.591Z f6c84568-6251-44cf-8f2d-715exxxxxxxx [INFO] b'{  "TopicOwner":"topic account id",  "Message":"mock mns message",  "Subscriber":"subscriber account id",  "PublishTime":1658235558094,  "SubscriptionName":"test-5bf13c7e",  "MessageMD5":"652BF0E6297840015247C3xxxxxxx",  "TopicName":"fc-example",  "MessageId":"3405CA51807661353B3xxxxxxxx"}'
FC Invoke End RequestId: f6c84568-6251-44cf-8f2d-715exxxxxxxx
Duration: 9.79 ms, Billed Duration: 10 ms, Memory Size: 128 MB, Max Memory Used: 24.50 MB
========= FC invoke Logs end =========
FC Invoke instanceId: c-62d7b590-2a0ccbf206baxxxxxxxx
FC Invoke Result:
mns_topic trigger event = b'{  "TopicOwner":"topic account id",  "Message":"mock mns message",  "Subscriber":"subscriber account id",  "PublishTime":1658235558094,  "SubscriptionName":"test-5bf13c7e",  "MessageMD5":"652BF0E6297840015247C3xxxxxxx",  "TopicName":"fc-example",  "MessageId":"3405CA51807661353B3xxxxxxxx"}'
End of method: invoke
```

## 注意事项
1. MNS消息服务和函数计算建议部署在同一个地域