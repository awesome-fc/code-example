# golang tracing 示例程序
本示例为您展示了 Golang Runtime 使用jaeger SDK的链路追踪使用示例。




## 准备开始
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 在 service 处添加角色配置。
```yaml
role: acs:ram::{your account id}:role/aliyunfcdefaultrole
```

#### 2. 安装依赖并部署

编译部署代码包
```shell
s deploy
```


#### 3. 调用测试

```shell
s invoke -e '{"name":"your name"}'
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
2022/10/14 09:28:45.857120 start
FC Invoke Start RequestId: 8b0ba671-a2ab-4f29-9ac4-c4858537e8f7
hello world
FC Invoke End RequestId: 8b0ba671-a2ab-4f29-9ac4-c4858537e8f7

Duration: 260.36 ms, Billed Duration: 261 ms, Memory Size: 128 MB, Max Memory Used: 6.94 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-63492bcd-df1d5ea028ae413880b4

FC Invoke Result:
"hello world! 你好，your name!"


End of method: invoke
```



### 方式二、使用控制台创建

#### 1. 编译打包

```shell
# 编译部署
GOOS=linux GOARCH=amd64 go build main.go
# 打包文件
zip main.zip main
```


#### 2. 创建服务和函数

创建服务
- 点击创建服务
- 点击显示高级选项
- 开启链路追踪功能
- 其他设置可使用默认


  ![img_5](assets/img_1.png)

选择服务（或创建服务）后，单击创建函数，如图所示
- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 Go1
- 选择函数触发方式：通过事件请求触发
- [请求处理程序](https://help.aliyun.com/document_detail/323526.html)（函数入口）设为：main
- 其他设置使用默认


> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)


#### 3. 配置测试参数

点击配置参数，创建新测试事件

![img_5](assets/img_2.png)

#### 4. 测试函数

点击链路追踪，查看函数的调用链路、耗时时间等信息。

![img_5](assets/img_3.png)