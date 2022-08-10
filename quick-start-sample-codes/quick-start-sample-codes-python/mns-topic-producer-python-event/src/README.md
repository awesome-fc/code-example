# Python3 消息服务 MNS 主题模型生产者示例

本示例为您展示了 python runtime 的 [消息服务MNS](https://help.aliyun.com/document_detail/27414.html) 主题模型生产者示例。
本示例使用了MNS 的主题模型作为示例，与函数计算中的 MNS 主题触发器一起实现了消息服务的生产者-消费者模型。
MNS 的配置在函数的环境变量配置中（参考s.yaml)。

本示例使用 [MNS官方Python SDK](https://help.aliyun.com/document_detail/32294.html)。

## 准备开始
- 一个可用的mns主题，可参考MNS官方文档[主题模型快速入门-创建主题](https://help.aliyun.com/document_detail/34424.html) 创建。
- 有MNS权限的RAM用户
  - 建议直接使用函数计算默认的角色 AliyunFCDefaultRole
  - 也可参考MNS官方文档[开通消息服务MNS并授权](https://help.aliyun.com/document_detail/27423.html)，函数计算需要该RAM密钥访问MNS主题。
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始

### 方式一、使用控制台创建

#### 1. 安装依赖和部署代码包

```shell
# 安装依赖到 /code 目录
cd code && pip3 install -r requirements.txt -t .
# 打包文件
cd code && zip -r python3-mns-topic-producer.zip *
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
![img_1.png](https://img.alicdn.com/imgextra/i1/O1CN01fuQDxG1cZG9R5jkAH_!!6000000003614-2-tps-2742-334.png)

函数环境变量配置：
![img_2.png](https://img.alicdn.com/imgextra/i4/O1CN01FsiWdg28rdkm5DMp3_!!6000000007986-2-tps-1962-508.png)

#### 4. 设置服务角色配置
在编辑服务页面，选择服务角色，推荐选择函数计算默认设置的角色 AliyunFCDefaultRole。
也可以自定义服务角色，并添加权限策略AliyunMNSFullAccess，或自定义权限策略，详情见文档 [授权策略和示例](https://help.aliyun.com/document_detail/27447.html)
![img_3.png](https://img.alicdn.com/imgextra/i3/O1CN01U35W371pYspseip5E_!!6000000005373-2-tps-2562-1014.png)

#### 5. 测试函数

返回结果如下所示
```bash
Publish Message Succeed. MessageBody:I am a test message. MessageID:20A37C322A6B4E2D486664xxxxxxxxxx
```

### 方式二、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 修改 role 配置，角色需要有mns的读写权限，建议使用函数计算默认的role, `acs:ram::{AccountID}:role/aliyunfcdefaultrole`
- 修改 environmentVariables 配置，填入 MnsEndpoint 和 TopicName

```yaml
        environmentVariables:
          MNS_ENDPOINT: "{{ mnsEndpoint }}" # 设置MNS访问地址
          MNS_TOPIC_NAME: "{{ mnsTopicName }}" # 设置MNS主题名称
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
FC Invoke Start RequestId: 39e64053-add7-40cf-8df8-xxxxxxxxxx
...
FC Invoke End RequestId: 39e64053-add7-40cf-8df8-xxxxxxxxxx
Duration: 115.10 ms, Billed Duration: 116 ms, Memory Size: 128 MB, Max Memory Used: 37.27 MB
========= FC invoke Logs end =========
FC Invoke instanceId: c-62d78387-e1a681251axxxxxxxxxx
FC Invoke Result:
Publish Message Succeed. MessageBody:I am a test message. MessageID:20A37C322A6B4E2D5C3564xxxxxxxxxx
End of method: invoke
```

## 注意事项
1. MNS消息服务和函数计算建议部署在同一个地域
