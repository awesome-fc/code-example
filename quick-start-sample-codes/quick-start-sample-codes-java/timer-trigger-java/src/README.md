# Java11 定时器触发FC函数示例

 本示例为您展示了 Java11 的定时器触发FC函数示例。

## 准备开始

 - [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始

### 方式一. 使用控制台创建

#### 1. 安装依赖和部署代码包

```shell
# 安装依赖并编译为jar包，对应jar包将在target目录下
mvn clean package
```

#### 2. 创建函数

 选择服务（或创建服务）后，单击创建函数，如图所示

 - 选择 `使用标准 Runtime 从零创建`
 - 填入函数名称
 - 代码上传方式选择`通过jar包上传代码`上传相应jar压缩包
 - 选择运行环境 Java 11
 - 选择函数触发方式：通过事件请求触发
 - 配置触发器选择`定时触发器`，并填入时间间隔如1（分钟）
 - 其他设置使用默认

> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)

#### 3. 设置服务角色配置

 在编辑服务页面，选择服务角色，推荐选择函数计算默认设置的角色 AliyunFCDefaultRole。
 也可以自定义服务角色，并添加权限策略，或自定义权限策略，详情见文档 [授权策略和示例](https://help.aliyun.com/document_detail/253969.html)



#### 4. 测试函数

如果开启日志服务，可以在`调用日志`中看到一下日志信息，时间间隔为1分钟。

```bash
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:49FC Invoke Start RequestId: 2dba780c-f1f6-4fd6-a572-d51f91a3a29b
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:492022-07-29 10:25:49.153 [INFO] [2dba780c-f1f6-4fd6-a572-d51f91a3a29b] triggerTime: 2022-07-29T10:25:49Z
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:492022-07-29 10:25:49.154 [INFO] [2dba780c-f1f6-4fd6-a572-d51f91a3a29b] triggerName: timer-java11
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:492022-07-29 10:25:49.154 [INFO] [2dba780c-f1f6-4fd6-a572-d51f91a3a29b] payload: testPayload
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:49FC Invoke End RequestId: 2dba780c-f1f6-4fd6-a572-d51f91a3a29b
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:49FC Invoke Start RequestId: c6e86b7e-a319-4e9b-9796-3204c2b3e971
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:492022-07-29 10:26:49.560 [INFO] [c6e86b7e-a319-4e9b-9796-3204c2b3e971] triggerTime: 2022-07-29T10:26:49Z
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:492022-07-29 10:26:49.561 [INFO] [c6e86b7e-a319-4e9b-9796-3204c2b3e971] triggerName: timer-java11
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:492022-07-29 10:26:49.561 [INFO] [c6e86b7e-a319-4e9b-9796-3204c2b3e971] payload: testPayload
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:49FC Invoke End RequestId: c6e86b7e-a319-4e9b-9796-3204c2b3e971
```



### 方式二. 使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置

 - [可选] 根据需要修改相应region

#### 2. 安装依赖并部署

 安装依赖库

```shell
# 使用s工具安装依赖，需要使用 docker
s build --use-docker
```

 部署代码

```bash
# 部署代码
s deploy -y
```

#### 3. 调用测试

```shell
s invoke -e '{
    "triggerTime": "2022-07-29T10:02:58Z",
    "triggerName": "nodejs14-timer",
    "payload": "TestPayload"
}'
```

 调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
FC Invoke Start RequestId: bb1f5703-7e9c-4494-8582-2cad07dbf9f7
[Name] Register [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@58372a00] as [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@com.aliyun.serverless.runtime.classloader.FunctionClassLoader@/code/HelloFCJava-1.0-SNAPSHOT.jar/code/original-HelloFCJava-1.0-SNAPSHOT.jar]: hash [d4d9f0d4] (normal mode)
2022-07-30 08:53:46.510 [INFO] [bb1f5703-7e9c-4494-8582-2cad07dbf9f7] triggerTime: 2022-07-29T10:02:58Z
2022-07-30 08:53:46.510 [INFO] [bb1f5703-7e9c-4494-8582-2cad07dbf9f7] triggerName: nodejs14-timer
2022-07-30 08:53:46.510 [INFO] [bb1f5703-7e9c-4494-8582-2cad07dbf9f7] payload: TestPayload
FC Invoke End RequestId: bb1f5703-7e9c-4494-8582-2cad07dbf9f7

Duration: 233.88 ms, Billed Duration: 234 ms, Memory Size: 128 MB, Max Memory Used: 81.87 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4f199-8b9aa92268c64154950d

FC Invoke Result:
Timer Payload: TestPayload


End of method: invoke
```

