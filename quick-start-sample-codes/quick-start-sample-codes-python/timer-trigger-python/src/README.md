# Python3 定时器触发FC函数示例

 本示例为您展示了 Python3 的定时器触发FC函数示例。

 ## 准备开始

 - [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

 ## 快速开始

 ### 方式一. 使用控制台创建

 #### 1. 安装依赖和部署代码包

 ```shell
# 安装依赖到 /code 目录
cd code && pip3 install -r requirements.txt -t .
# 打包文件
cd code && zip -r python3-timer.zip *
 ```

 #### 2. 创建函数

 选择服务（或创建服务）后，单击创建函数，如图所示

 - 选择 `使用标准 Runtime 从零创建`
 - 填入函数名称
 - 代码上传方式选择`通过zip包上传代码`上传相应代码压缩包
 - 选择运行环境 Python 3.6
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
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:37FC Invoke Start RequestId: bd03dbc7-ce59-488d-be55-b579670afe4c
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:372022-07-29 18:17:37 bd03dbc7-ce59-488d-be55-b579670afe4c [INFO] event: b'{"triggerTime":"2022-07-29T10:17:37Z","triggerName":"timer-python3","payload":"testPayload"}'
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:372022-07-29 18:17:37 bd03dbc7-ce59-488d-be55-b579670afe4c [INFO] triggerName: timer-python3
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:372022-07-29 18:17:37 bd03dbc7-ce59-488d-be55-b579670afe4c [INFO] triggerTime = 2022-07-29T10:17:37Z
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:37FC Invoke End RequestId: bd03dbc7-ce59-488d-be55-b579670afe4c
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:37FC Invoke Start RequestId: ef169a63-4342-4019-b129-f528fd9463f8
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:372022-07-29 18:18:37 ef169a63-4342-4019-b129-f528fd9463f8 [INFO] event: b'{"triggerTime":"2022-07-29T10:18:37Z","triggerName":"timer-python3","payload":"testPayload"}'
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:372022-07-29 18:18:37 ef169a63-4342-4019-b129-f528fd9463f8 [INFO] triggerName: timer-python3
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:372022-07-29 18:18:37 ef169a63-4342-4019-b129-f528fd9463f8 [INFO] triggerTime = 2022-07-29T10:18:37Z
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:37FC Invoke End RequestId: ef169a63-4342-4019-b129-f528fd9463f8
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
FunctionCompute python3 runtime inited.
FC Invoke Start RequestId: a785f24e-b67a-400d-b3cc-4472a8587400
2022-07-30T08:49:52.165Z a785f24e-b67a-400d-b3cc-4472a8587400 [INFO] event: b'{\n    "triggerTime": "2022-07-29T10:02:58Z",\n    "triggerName": "nodejs14-timer",\n    "payload": "TestPayload"\n}'
2022-07-30T08:49:52.166Z a785f24e-b67a-400d-b3cc-4472a8587400 [INFO] triggerName: nodejs14-timer
2022-07-30T08:49:52.166Z a785f24e-b67a-400d-b3cc-4472a8587400 [INFO] triggerTime = 2022-07-29T10:02:58Z
FC Invoke End RequestId: a785f24e-b67a-400d-b3cc-4472a8587400

Duration: 3.90 ms, Billed Duration: 4 ms, Memory Size: 128 MB, Max Memory Used: 24.86 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4f0af-aa8569cde2ed43aaaf04

FC Invoke Result:
Timer Payload:TestPayload


End of method: invoke
 ```

