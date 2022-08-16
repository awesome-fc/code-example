# kafka-producer-fc-event-golang 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=kafka-producer-fc-event-golang&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=kafka-producer-fc-event-golang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=kafka-producer-fc-event-golang&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=kafka-producer-fc-event-golang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=kafka-producer-fc-event-golang&type=packageDownload">
  </a>
</p>


## 前期准备

### 权限准备

使用该项目，确认您的操作账户拥有以下的产品权限 / 策略：


| 服务/业务 | 函数计算                                                     |
| --------- | ------------------------------------------------------------ |
| 权限/策略 | AliyunFCFullAccess<br/>AliyunKafkaFullAccess<br/>AliyunVPCReadOnlyAccess |


### 资源准备

  * 一个可用的Kafka消息队列，可参考消息队列Kafka版官方文档[消息队列快速入门](https://help.aliyun.com/document_detail/99949.html)。

    - 创建VPC专有网络（推荐在生产环境中也使用VPC），可参考[VPC官方文档](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg#task-1012575)。VPC控制台[链接](https://vpcnext.console.aliyun.com/)。至此即可拥有VPC和相应交换机。

    > 部署Kafka实例时会提示创建可用的VPC专有网络

  * 在Kafka控制台创建需要使用的Kafka Topic与Consumer Group。

### 环境准备

由于Go的Kafka客户端程序包含CGO，虽然Go拥有交叉编译器，但如果没有安装相应交叉编译C的工具链，无法直接编译出跨平台的可执行文件。即当我们使用了CGO时，要想实现跨平台编译，同时需要让C/C++代码也支持跨平台编译。

函数计算的环境为Linux / amd64，其他操作系统要构建相应可执行文件则需要相应环境：

- Mac

  ```bash
  # 下载linux编译工具链
  brew install FiloSottile/musl-cross/musl-cross
  ```
  
- Windows / Mac

  ```bash
  # 拉取镜像
  docker pull karalabe/xgo-latest
  ```
# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/)
- 为了能够成功部署本样例代码，您在部署过程中需要提供以下参数：
  - 地域 (region): 您需要通过这个参数配置您函数计算服务需要部署的地域，默认值为 cn-hangzhou (杭州)。
    - 为您提供的地域选项为：
      - cn-beijing (北京)
      - cn-hangzhou (杭州)
      - cn-shanghai (上海)
      - cn-qingdao (青岛)
      - cn-zhangjiakou (张家口)
      - cn-huhehaote (呼和浩特)
      - cn-shenzhen (深圳)
      - cn-chengdu (成都)
      - cn-hongkong (香港)
      - ap-southeast-1 (新加坡)
      - ap-southeast-2 (悉尼)
      - ap-southeast-3 (吉隆坡)
      - ap-southeast-5 (雅加达)
      - ap-northeast-1 (东京)
      - eu-central-1 (法兰克福)
      - eu-west-1 (伦敦)
      - us-west-1 (硅谷)
      - us-east-1 (弗吉尼亚)
      - ap-south-1 (孟买)
  - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为 kafka-producer-quick-start。
  - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 kafka-producer-event-function-golang。
  - vpcId: 我们推荐您使用VPC访问Kafka，选择创建Kafka实例时使用的VPC。注意需要在函数计算支持的 az。
  - vswitchIds:  使用vpc中的vswitch id，用于内网访问 kafka，注意需要在函数计算支持的az。
  - securityGroupId:  kafka 实例所在 vpc 的安全组id，可在`云服务器 ECS`控制台`网络与安全`菜单项找到。
  - Kafka接入点 (bootstrapServers): 您购买的Kafka实例的默认接入点，可以在实例详情中找到。
  - topicName: Kafka实例中某个topic name，此topic的数据生产会触发部署函数，需要您提前创建。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=kafka-producer-fc-event-golang) ，
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=kafka-producer-fc-event-golang)  该应用。（此方法在此应用下仅支持Linux环境。推荐使用 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署）

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：

  - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
  - 初始化项目：`s init kafka-producer-fc-event-golang -d kafka-producer-fc-event-golang`
  - 填入在以上模块介绍的参数
  - 进入项目目录，`cd kafka-producer-fc-event-golang`根据使用环境选择相应`pre-deploy`中`run`的字段。
  - 进行项目部署： `s deploy -y`
- 本地调试
  - 进入应用项目工程下，执行下面命令：`s invoke -e '{"Key": "test go kafka producer"}'`。
  - 即可查看到模拟事件触发函数后的结果。

```bash
========= FC invoke Logs begin =========
2022/08/08 06:44:34.255072 start
FC Initialize Start RequestId: be4c11f1-a99e-4421-852d-55ace08a6cad
2022-08-08T06:44:42.322Z be4c11f1-a99e-4421-852d-55ace08a6cad [INFO] main.go:37: Initializing the kafka config
FC Initialize End RequestId: be4c11f1-a99e-4421-852d-55ace08a6cad
FC Invoke Start RequestId: be4c11f1-a99e-4421-852d-55ace08a6cad
2022-08-08T06:44:42.366Z be4c11f1-a99e-4421-852d-55ace08a6cad [INFO] main.go:57: sending the message to kafka: test go kafka producer!
2022-08-08T06:44:42.394Z be4c11f1-a99e-4421-852d-55ace08a6cad [INFO] main.go:72: Delivered message to topic HelloTopic [7] at offset 19
FC Invoke End RequestId: be4c11f1-a99e-4421-852d-55ace08a6cad

Duration: 29.94 ms, Billed Duration: 30 ms, Memory Size: 128 MB, Max Memory Used: 15.91 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62f0b0d0-626f0eb5c5e94c42bf37

FC Invoke Result:
"Finish sending the message to kafka: test go kafka producer!"


End of method: invoke
```

​		

- 端对端测试

  - 登陆函数计算控制台，进入对应的函数配置测试参数为`{"key":"test kafka producer go"}`（注意参数需要配置为键值为"Key"的json形式【"Key"是在demo程序中设定的，可修改】）点击`测试函数`
  
  返回结果如下：

  ```bash
  "Finish sending the message to kafka: test kafka producer go!"
  ```
  日志如下：
  
  ```bash
  2022/08/08 06:45:36.589266 start
  FC Initialize Start RequestId: 0f2f1c1c-a89a-46fe-a588-96469cfbeba2
  2022-08-08 14:45:36 0f2f1c1c-a89a-46fe-a588-96469cfbeba2 [INFO] main.go:37: Initializing the kafka config
  FC Initialize End RequestId: 0f2f1c1c-a89a-46fe-a588-96469cfbeba2
  FC Invoke Start RequestId: 0f2f1c1c-a89a-46fe-a588-96469cfbeba2
  2022-08-08 14:45:36 0f2f1c1c-a89a-46fe-a588-96469cfbeba2 [INFO] main.go:57: sending the message to kafka: test kafka producer go!
  2022-08-08 14:45:36 0f2f1c1c-a89a-46fe-a588-96469cfbeba2 [INFO] main.go:72: Delivered message to topic HelloTopic [4] at offset 21
  FC Invoke End RequestId: 0f2f1c1c-a89a-46fe-a588-96469cfbeba2
  ```
  
  - 登陆Kafka控制台，查看对应实例的对应Topic`消息详情`，找到对应分区，即可查询到相应消息。
  
  

</deploy>

<appdetail id="flushContent">

# 应用详情



本应用仅作为学习和参考使用，您可以基于本项目进行二次开发和完善，实现自己的业务逻辑



</appdetail>

<devgroup>

## 开发者社区

您如果有关于错误的反馈或者未来的期待，您可以在 [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) 中进行反馈和交流。如果您想要加入我们的讨论组或者了解 FC 组件的最新动态，您可以通过以下渠道进行：

<p align="center">



| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| <center>微信公众号：`serverless`</center>                    | <center>微信小助手：`xiaojiangwh`</center>                   | <center>钉钉交流群：`33947367`</center>                      |

</p>

</devgroup>