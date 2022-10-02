# .NET Core 3.1 POCO 事件回调示例程序
本示例为您展示了 .NET Core 3.1 runtime 的 Stream 事件回调示例程序。


## 准备开始
- 已安装 .NET SDK 3.1 或 .NET SDK 2.1 (使用命令 `dotnet --list-sdks` 查看已安装的 SDK 版本)
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置

#### 2. 安装依赖并部署

编译部署代码包
```shell
s deploy
```
> 注意: `pom.xml` 中有配置 `pre-deploy` 脚本 `dotnet publish -c Release`, 在部署前会调用 `dotnet publish -c Release` 编译打包。

#### 3. 调用测试

```shell
s invoke --event-file event
```

调用函数时收到的响应如下所示：

```bash
Reading event file content:
Hello World

========= FC invoke Logs begin =========
FunctionCompute dotnetcore3.1 runtime inited.
FC Invoke Start RequestId: fe418d7a-ec9e-4adf-8565-xxxxxxxxxx
2022-09-22T08:29:57.331Z fe418d7a-ec9e-4adf-8565-xxxxxxxxxx [INFO] Handle request: fe418d7a-ec9e-4adf-8565-xxxxxxxxxx
FC Invoke End RequestId: fe418d7a-ec9e-4adf-8565-xxxxxxxxxx

Duration: 172.34 ms, Billed Duration: 173 ms, Memory Size: 256 MB, Max Memory Used: 13.12 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-632c1d04-48527c4xxxxxxxxxx

FC Invoke Result:
Hello World


End of method: invoke
```

### 方式二、使用控制台创建

#### 1. 编译打包

```shell
# 编译部署
cd HelloFcApp && dotnet publish -c Release -o ./target
# 打包文件
cd ./target && zip -r dotnet3-blank-stream-event.zip HelloFcApp.dll
```

#### 2. 创建函数

选择服务（或创建服务）后，单击创建函数

- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 .NET Core 3.1
- 选择函数触发方式：通过事件请求触发
- [请求处理程序](https://help.aliyun.com/document_detail/112379.html)（函数入口）设为：HelloFcApp::Example.Hello::StreamHandler
- 其他设置使用默认

> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)


#### 3. 配置测试参数

在函数管理页面的测试函数标签栏配置测试参数

#### 4. 测试函数

返回结果如下所示

```bash
Hello World

```
