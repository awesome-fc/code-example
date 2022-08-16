# Nodejs14 定时器触发FC函数示例

 本示例为您展示了 Nodejs14 的定时器触发FC函数示例。

 ## 准备开始

 - [可选] 安装并配置 Serverless Devs 工具。(https://help.aliyun.com/document_detail/195474.html)

 ## 快速开始

 ### 方式一. 使用控制台创建

 #### 1. 安装依赖和部署代码包

 ```shell
# 在code目录下安装依赖并编译
npm install
# 压缩
zip code.zip -r ./*
 ```

 #### 2. 创建函数

 选择服务（或创建服务）后，单击创建函数，如图所示

 - 选择 `使用标准 Runtime 从零创建`
 - 填入函数名称
 - 代码上传方式选择`通过zip包上传代码`上传相应代码压缩包
 - 选择运行环境 Nodejs14
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
62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:58FC Invoke Start RequestId: 209d5189-18c4-4b9e-9c3f-12d255cc14d4
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] whole event: {"triggerTime":"2022-07-29T10:02:58Z","triggerName":"nodejs14-timer","payload":"TestPayload"}
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] triggerName:  nodejs14-timer
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] triggerTime:  2022-07-29T10:02:58Z
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] triggerMessgae:  TestPayload
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:58FC Invoke End RequestId: 209d5189-18c4-4b9e-9c3f-12d255cc14d4
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:58FC Invoke Start RequestId: d5d22563-78ba-4198-b819-4d30154bdace
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] whole event: {"triggerTime":"2022-07-29T10:03:58Z","triggerName":"nodejs14-timer","payload":"TestPayload"}
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] triggerName:  nodejs14-timer
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] triggerTime:  2022-07-29T10:03:58Z
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] triggerMessgae:  TestPayload
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:58FC Invoke End RequestId: d5d22563-78ba-4198-b819-4d30154bdace
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
FC Invoke Start RequestId: e8ff75e9-08e6-4fd9-be64-d91894700d86
load code for handler:index.handler
}   "payload": "TestPayload"-timer",:58Z",9-be64-d91894700d86 [verbose] whole event: {
2022-07-30T08:44:59.548Z e8ff75e9-08e6-4fd9-be64-d91894700d86 [verbose] triggerName:  nodejs14-timer
2022-07-30T08:44:59.548Z e8ff75e9-08e6-4fd9-be64-d91894700d86 [verbose] triggerTime:  2022-07-29T10:02:58Z
2022-07-30T08:44:59.548Z e8ff75e9-08e6-4fd9-be64-d91894700d86 [verbose] triggerMessgae:  TestPayload
FC Invoke End RequestId: e8ff75e9-08e6-4fd9-be64-d91894700d86

Duration: 3.92 ms, Billed Duration: 4 ms, Memory Size: 128 MB, Max Memory Used: 48.84 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4ef8b-5959048c584a48f8bd95

FC Invoke Result:
timer trigger:TestPayload


End of method: invoke
 ```

