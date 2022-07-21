# Python3 消息服务 MNS 队列模型生产者示例

本示例为您展示了 python runtime 的 [消息服务MNS](https://help.aliyun.com/document_detail/27414.html) 队列模型生产者示例。
本示例使用了MNS 的队列模型作为示例，与函数计算中的 MNS 队列触发器一起实现了消息服务的生产者-消费者模型。
MNS 的配置在函数的环境变量配置中（参考s.yaml)。

> 若使用主题模型，请参考python3-mns-topic-producer示例。

本示例使用 [MNS官方Python SDK](https://help.aliyun.com/document_detail/32294.html)。

## 准备开始
- 一个可用的mns队列，可参考MNS官方文档[队列模型快速入门-创建队列](https://help.aliyun.com/document_detail/34417.html) 创建。
- 有MNS权限的RAM用户
  - 建议直接使用函数计算默认的角色 AliyunFCDefaultRole
  - 也可参考MNS官方文档[开通消息服务MNS并授权](https://help.aliyun.com/document_detail/27423.html)，函数计算需要该RAM密钥访问MNS队列。
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始

### 方式一、使用控制台创建

#### 1. 安装依赖和部署代码包

```shell
# 安装依赖到 /code 目录
cd code && pip3 install -r requirements.txt -t .
# 打包文件
cd code && zip -r python3-mns-queue-producer.zip *
```

#### 2. 创建函数
选择服务（或创建服务）后，单击创建函数，如图所示
- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 Python 3.6/3.9
- 选择函数触发方式：通过事件请求触发
- 其他设置使用默认

> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)

#### 3. 设置 initializer 回调函数配置和环境变量配置

回调函数配置：
![img_2.png](assets/20220719164834.jpg)

函数环境变量配置：
![img_2.png](assets/20220719164825.jpg)

#### 4. 设置服务角色配置
在编辑服务页面，选择服务角色，推荐选择函数计算默认设置的角色 AliyunFCDefaultRole。
也可以自定义服务角色，并添加权限策略AliyunMNSFullAccess，或自定义权限策略，详情见文档 [授权策略和示例](https://help.aliyun.com/document_detail/27447.html)
![img_3.png](assets/20220719171807.jpg)

#### 5. 测试函数

返回结果如下所示
```bash
Send Message Succeed. MessageBody:I am a test message. MessageID:494D03462A6B444C35256Axxxxxxxxxx
```

### 方式二、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 修改 role 配置，角色需要有mns的读写权限，建议使用函数计算默认的role, `acs:ram::{AccountID}:role/aliyunfcdefaultrole`
- 修改 environmentVariables 配置，填入 MnsEndpoint 和 QueueName

```yaml
        environmentVariables:
          MnsEndpoint: "http://{AccountID}.mns.{Region}.aliyuncs.com" # 设置MNS访问地址
          QueueName: "fc-example" # 设置MNS队列名称
```

#### 2. 安装依赖并部署

安装依赖库 aliyun-mns

```shell
# 使用s工具安装依赖，需要使用 docker
s build --use-docker
```

> 注意: 使用 `s build --use-docker` 会将依赖安装到代码包的 `.s/python` 目录下，可以在本地的 `.s/build/artifacts/{serviceName}/{functionName}/` 查看。

部署代码

```bash
# 部署代码，会自动添加环境变量 PYTHONUSERBASE=/code/.s/python，该环境变量是必须的
s deploy
```

#### 3. 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
FunctionCompute python3 runtime inited.
FC Invoke Start RequestId: 3772f753-36f1-4898-aa18-ef14xxxxxxxx
...
FC Invoke End RequestId: 3772f753-36f1-4898-aa18-ef14xxxxxxxx
Duration: 117.00 ms, Billed Duration: 117 ms, Memory Size: 128 MB, Max Memory Used: 37.05 MB
========= FC invoke Logs end =========
FC Invoke instanceId: c-62d919bc-2a9e69184f94xxxxxxxx
FC Invoke Result:
Send Message Succeed. MessageBody:I am a test message. MessageID:494D03462A6B4E2D7F986AB5xxxxxxxx
End of method: invoke

```

## 注意事项
1. MNS消息服务和函数计算建议部署在同一个地域
