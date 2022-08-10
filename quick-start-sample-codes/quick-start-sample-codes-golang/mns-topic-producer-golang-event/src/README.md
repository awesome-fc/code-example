# Golang 消息服务 MNS 主题模型生产者示例

本示例为您展示了 golang runtime 的 [消息服务MNS](https://help.aliyun.com/document_detail/27414.html) 主题模型生产者示例。
本示例使用了MNS 的主题模型作为示例，与函数计算中的 MNS 主题触发器一起实现了消息服务的生产者-消费者模型。
MNS 的配置在函数的环境变量配置中（参考s.yaml)。

本示例使用 MNS 官方 [Go SDK](https://help.aliyun.com/document_detail/116629.html)

## 准备开始
- 一个可用的mns主题，可参考MNS官方文档[主题模型快速入门-创建主题](https://help.aliyun.com/document_detail/34424.html) 创建。
- 有MNS权限的RAM用户
  - 建议直接使用函数计算默认的角色 AliyunFCDefaultRole
  - 也可参考MNS官方文档[开通消息服务MNS并授权](https://help.aliyun.com/document_detail/27423.html)，函数计算需要该RAM密钥访问MNS主题。
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始

### 方式一、使用控制台创建

#### 1. 编译打包

```shell
# 编译部署
GOOS=linux GOARCH=amd64 go build main.go
# 打包文件
zip main.zip main
```

> 以上命令只适用于 Linux/Mac 环境，Windows 环境可参考官方文档: [在 Windows 下编译打包](https://help.aliyun.com/document_detail/418490.html#section-qfg-n9c-m9v)

#### 2. 创建函数
选择服务（或创建服务）后，单击创建函数，如图所示
- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 Go 1
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
"Publish succ, message id: CC9C55A980767F854A158DA3xxxxxxxx, messagebody md5: 48E9198EE9E413E274A0E9F2xxxxxxxx"
```

### 方式二、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置

- 根据需要修改 access 配置
- 修改 environmentVariables 配置，填入 MnsEndpoint 和 TopicName

```yaml
        environmentVariables:
          MNS_ENDPOINT: "{{ mnsEndpoint }}" # 设置MNS访问地址
          MNS_TOPIC_NAME: "{{ mnsTopicName }}" # 设置MNS主题名称
```

#### 2. 安装依赖并部署

部署代码

```bash
s deploy
```

#### 3. 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
2022/07/28 04:04:50.883302 start
FC Initialize Start RequestId: dd329059-8bc0-41b9-bd07-45da857bxxxx
FC Initialize End RequestId: dd329059-8bc0-41b9-bd07-45da857bxxxx
FC Invoke Start RequestId: 00cb1be4-126e-47a5-b7bc-0efad01fxxxx
FC Invoke End RequestId: 00cb1be4-126e-47a5-b7bc-0efad01fxxxx

Duration: 77.34 ms, Billed Duration: 78 ms, Memory Size: 128 MB, Max Memory Used: 12.16 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e20ae2-9eb6a54c1f154d39xxxx

FC Invoke Result:
"Publish succ, message id: CC9C55A980767F854A158DA3xxxxxxxx, messagebody md5: 48E9198EE9E413E274A0E9F2xxxxxxxx"


End of method: invoke

```

## 注意事项
1. MNS消息服务和函数计算建议部署在同一个地域
