# Golang 定时器触发FC函数示例

 本示例为您展示了 Golang 的定时器触发FC函数示例。

 ## 准备开始
 - [可选] 安装并配置 Serverless Devs 工具。(https://help.aliyun.com/document_detail/195474.html)

 ## 快速开始

 ### 方式一. 使用控制台创建

 #### 1. 安装依赖和部署代码包

> 部署Go代码包更多详情可见(https://help.aliyun.com/document_detail/418490.html)

 ```shell
 # 在code目录下安装依赖并编译
 go mod tidy
 GOOS=linux GOARCH=amd64 go build main.go
 # 压缩可执行文件
 zip fc-golang-demo.zip main
 ```

 #### 2. 创建函数
 选择服务（或创建服务）后，单击创建函数，如图所示
 - 选择 `使用标准 Runtime 从零创建`
 - 填入函数名称
 - 代码上传方式选择`通过zip包上传代码`上传相应代码压缩包
 - 选择运行环境 Go 1
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
 c-62e46764-132ecab75fa946deb5462022-07-30 11:43:58FC Invoke Start RequestId: 37421a15-6150-462b-bf0b-685e48b66128
 c-62e46764-132ecab75fa946deb5462022-07-30 11:43:582022-07-30 11:43:58 37421a15-6150-462b-bf0b-685e48b66128 [INFO] main.go:20: triggerName:  timer-go1
 c-62e46764-132ecab75fa946deb5462022-07-30 11:43:582022-07-30 11:43:58 37421a15-6150-462b-bf0b-685e48b66128 [INFO] main.go:21: triggerTime:  2022-07-30T03:43:58Z
 c-62e46764-132ecab75fa946deb5462022-07-30 11:43:582022-07-30 11:43:58 37421a15-6150-462b-bf0b-685e48b66128 [INFO] main.go:22: payload: testPayload
 c-62e46764-132ecab75fa946deb5462022-07-30 11:43:58FC Invoke End RequestId: 37421a15-6150-462b-bf0b-685e48b66128
 c-62e46764-132ecab75fa946deb5462022-07-30 11:44:58FC Invoke Start RequestId: f2145f49-6a99-4a64-bd88-6c49c6a37e25
 c-62e46764-132ecab75fa946deb5462022-07-30 11:44:582022-07-30 11:44:58 f2145f49-6a99-4a64-bd88-6c49c6a37e25 [INFO] main.go:20: triggerName:  timer-go1
 c-62e46764-132ecab75fa946deb5462022-07-30 11:44:582022-07-30 11:44:58 f2145f49-6a99-4a64-bd88-6c49c6a37e25 [INFO] main.go:21: triggerTime:  2022-07-30T03:44:58Z
 c-62e46764-132ecab75fa946deb5462022-07-30 11:44:582022-07-30 11:44:58 f2145f49-6a99-4a64-bd88-6c49c6a37e25 [INFO] main.go:22: payload: testPayload
 ```



 ### 方式二、使用 Serverless Devs 工具编译部署

 #### 1. 修改 s.yaml 配置
 - [可选] 根据需要修改相应region

 #### 2. 安装依赖并部署

 安装依赖库

 ```shell
 # 使用s工具安装依赖，需要使用 docker
 s build --use-docker
 ```

 > 注意: 使用 `s build --use-docker` 会将依赖安装到代码包的 `.s/python` 目录下，可以在本地的 `.s/build/artifacts/{serviceName}/{functionName}/` 查看。

 部署代码

 ```bash
 # 部署代码，会自动添加环境变量 PYTHONUSERBASE=/code/.s/python，该环境变量是必须的
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
FC Invoke Start RequestId: 953240e4-a6e3-452b-a2ed-aef6518dc3af
2022-07-30T08:17:49.606Z 953240e4-a6e3-452b-a2ed-aef6518dc3af [INFO] main.go:20: triggerName:  nodejs14-timer
2022-07-30T08:17:49.606Z 953240e4-a6e3-452b-a2ed-aef6518dc3af [INFO] main.go:21: triggerTime:  2022-07-29T10:02:58Z
2022-07-30T08:17:49.606Z 953240e4-a6e3-452b-a2ed-aef6518dc3af [INFO] main.go:22: payload: TestPayload
FC Invoke End RequestId: 953240e4-a6e3-452b-a2ed-aef6518dc3af

Duration: 1.13 ms, Billed Duration: 2 ms, Memory Size: 128 MB, Max Memory Used: 9.69 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4e919-8073a286587848f3b97c

FC Invoke Result:
"Timer Payload: TestPayload"


End of method: invoke
 ```
