# Golang Kafka消息队列生产者示例

本示例为您展示了 Golang 的 [消息队列Kafka](https://help.aliyun.com/document_detail/68151.html?spm=5176.167616.J_5253785160.5.2dfe6feexRPqMj) 生产者示例。（消费者示例见Kafka触发器示例）

本示例使用[消息队列Kafka版官方Golang SDK](https://help.aliyun.com/document_detail/183934.html)。

 ## 准备开始

 - 一个可用的Kafka消息队列，可参考消息队列Kafka版官方文档[消息队列快速入门](https://help.aliyun.com/document_detail/99949.html)。

   - 创建VPC专有网络（推荐在生产环境中也使用VPC），可参考[VPC官方文档](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg#task-1012575)。VPC控制台[链接](https://vpcnext.console.aliyun.com/)。至此即可拥有VPC和相应交换机。

   > 部署Kafka实例时会提示创建可用的VPC专有网络

 - 能够编译CGO程序至Linux / amd64的环境（在下一部分讲述具体方法）

 - [可选] 安装并配置 Serverless Devs 工具。(https://help.aliyun.com/document_detail/195474.html)

 ## 快速开始

 ### 方式一. 使用控制台创建

 #### 1. 安装依赖和部署代码包

> 函数计算部署Go代码包的更多详情可见(https://help.aliyun.com/document_detail/418490.html)
>
> 含有CGO代码的项目如何实现跨平台编译(https://segmentfault.com/a/1190000038938300)

由于Kafka客户端包包含CGO，虽然Go拥有交叉编译器，但如果没有安装相应交叉编译C的工具链，无法直接编译出跨平台的可执行文件。即当我们使用了CGO时，要想实现跨平台编译，同时需要让C/C++代码也支持跨平台。在此说明3个解决办法：

- 直接使用Linux/amd64平台机器编译

  - ```shell
    # 直接开启CGO并编译
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build main.go
    ```

- Mac下的可行方案：

  - ```shell
    # 下载linux编译工具链
    brew install FiloSottile/musl-cross/musl-cross
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -v -tags musl
    ```

- 各平台通用方案（Docker）:

  - ```shell
    # 拉取镜像
    docker pull karalabe/xgo-latest
    # 在code目录下运行：
    docker run -v $(pwd):/go/src/gocode -w /go/src/gocode --entrypoint='' karalabe/xgo-latest /bin/bash -c "CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build main.go"
    # 运行的指令即在容器内编译，可在宿主机code目录下查看到编译好的可执行文件
    ```

最后压缩可执行文件即可。



 #### 2. 创建服务

推荐在Kafka实例相同Region创建服务。

创建服务时在`高级选项`中`服务角色`选择AliyunFcDefaultRole（如没有则根据提示创建相应角色），并开启`允许访问VPC`，选取创建Kafka实例时所选择的`专有网络`、`交换机`与对应的`安全组(Kafka实例部署后自动创建)`。

![CreateService.png](assets/CreateService.png)



#### 3. 创建函数

 创建服务后，单击创建函数

 - 选择 `使用标准 Runtime 从零创建`
 - 填入函数名称
 - 代码上传方式选择`通过zip包上传代码`上传相应代码压缩包
 - 选择运行环境 Go 1
 - 选择函数触发方式：通过事件请求触发
 - 其他设置使用默认

 > 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)



#### 4. 配置环境变量与实例生命周期回调

在函数详情中的`函数配置`模块设置环境变量并在实例生命周期回调中开启Initializer与PreStop回调程序。

其中环境变量：

- BOOTSTRAP_SERVERS设置为Kafka实例详情内`接入点信息`对应的`默认接入点`地址。

- TOPIC_NAME设置为相应发送消息到的Topic（需要在Kafka消息队列版中提前创建）

![FunctionConfig.png](assets/FunctionConfig.png)



 #### 5. 测试函数

对于Go的测试，参数需要配置为键值为"Key"的json形式（"Key"是在demo程序中设定的，可修改），如：

```json
{
    "Key": "test go"
}
```

日志输出结果：

 ```bash
2022/07/31 04:27:48.315016 start
FC Initialize Start RequestId: f1326a21-5f69-4090-8283-beb458b2d257
2022-07-31 12:27:48 f1326a21-5f69-4090-8283-beb458b2d257 [INFO] main.go:37: Initializing the kafka config
FC Initialize End RequestId: f1326a21-5f69-4090-8283-beb458b2d257
FC Invoke Start RequestId: 7e7931d4-3f62-452e-86a7-4f8190bbabb7
2022-07-31 12:27:49 7e7931d4-3f62-452e-86a7-4f8190bbabb7 [INFO] main.go:57: sending the message to kafka: test go!
2022-07-31 12:27:49 7e7931d4-3f62-452e-86a7-4f8190bbabb7 [INFO] main.go:72: Delivered message to topic HelloTopic [9] at offset 23
FC Invoke End RequestId: 7e7931d4-3f62-452e-86a7-4f8190bbabb7

 ```



 ### 方式二. 使用 Serverless Devs 工具编译部署

 #### 1. 修改 s.yaml 配置

- 修改region、serviceName、functionName（设置和Kafka实例相同的region）

- 根据平台不同，修改pre-dploy中run的命令，如果为mac平台则在安装相应linux工具链后将run设置为`CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -a -v -tags musl`即可。
- 修改vpcConfig，将Kafka实例对应的VPC ID、安全组ID（可在**云服务器 ECS**控制台`网络与安全`菜单项找到）、vSwitchID填入。
- 修改 environmentVariables 配置，填入 BOOTSTRAP_SERVERS 和 TOPIC_NAME

 #### 2. 安装依赖并部署

 安装依赖库

 ```shell
# 使用s工具安装依赖，需要使用 docker
s build --use-docker
 ```

 部署代码

 ```bash
s deploy -y
 ```

 #### 3. 调用测试

 ```shell
s invoke -e '{
    "Key": "test go serverless devs"
}'
 ```

 调用函数时收到的响应如下所示：

 ```bash
========= FC invoke Logs begin =========
2022/07/31 04:37:27.116361 start
FC Initialize Start RequestId: 85422744-98de-42a5-b9d6-67f3344f832d
2022-07-31T04:37:27.149Z 85422744-98de-42a5-b9d6-67f3344f832d [INFO] main.go:37: Initializing the kafka config
FC Initialize End RequestId: 85422744-98de-42a5-b9d6-67f3344f832d
FC Invoke Start RequestId: 85422744-98de-42a5-b9d6-67f3344f832d
2022-07-31T04:37:28.239Z 85422744-98de-42a5-b9d6-67f3344f832d [INFO] main.go:57: sending the message to kafka: test go serverless devs!
2022-07-31T04:37:28.249Z 85422744-98de-42a5-b9d6-67f3344f832d [INFO] main.go:72: Delivered message to topic HelloTopic [9] at offset 24
FC Invoke End RequestId: 85422744-98de-42a5-b9d6-67f3344f832d

Duration: 1001.97 ms, Billed Duration: 1002 ms, Memory Size: 128 MB, Max Memory Used: 16.16 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e60706-413948b4cb5d4ed18c78

FC Invoke Result:
"Finish sending the message to kafka: test go serverless devs!"


End of method: invoke
 ```

