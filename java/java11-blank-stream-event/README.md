# java11 stream 事件回调示例程序
本示例为您展示了 Java runtime 的 stream 事件回调示例程序。


## 准备开始
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
> 注意: `pom.xml` 中有配置 `pre-deploy` 脚本 `mvn package`, 在部署前会调用 `mvn package` 编译打包。

#### 3. 调用测试

```shell
s invoke --event-file event.txt
```

调用函数时收到的响应如下所示：

```bash
Reading event file content:
Hello world, Hello FC!
========= FC invoke Logs begin =========
FC Invoke Start RequestId: ef87e911-7d20-4d5b-b7c3-xxxxxxxxx
2022-04-08 04:12:35.328 [INFO] Input data: Hello world, Hello FC!
FC Invoke End RequestId: ef87e911-7d20-4d5b-b7c3-xxxxxxxxx

Duration: 30.87 ms, Billed Duration: 31 ms, Memory Size: 128 MB, Max Memory Used: 66.68 MB
========= FC invoke Logs end =========

FC Invoke Result:
Hello world, Hello FC!


End of method: invoke
```

### 方式二、使用控制台创建

#### 1. 编译打包

```shell
# 编译部署
mvn package
# 打包文件
cd target && zip -r java11-stream-event-blank.zip *
```

#### 2. 创建函数
选择服务（或创建服务）后，单击创建函数，如图所示
- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 java11/java8
- 选择函数触发方式：通过事件请求触发
- 其他设置使用默认

> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)

#### 3. 配置测试参数
在函数管理页面的测试函数标签栏配置测试参数

#### 4. 测试函数

返回结果如下所示
```bash
"Hello world, Hello FC!"
```