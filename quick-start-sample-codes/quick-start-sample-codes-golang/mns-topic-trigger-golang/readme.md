# mns-topic-trigger-golang 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-topic-trigger-golang&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=mns-topic-trigger-golang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-topic-trigger-golang&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=mns-topic-trigger-golang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-topic-trigger-golang&type=packageDownload">
  </a>
</p>

<description>

快速部署一个 Golang 的 Event 类型的 MNS Topic trigger 函数到阿里云函数计算。

</description>

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess <br> AliyunMNSFullAccess |

| 服务/业务 | 访问控制(RAM) |     
| --- |  --- |   
| 资源/创建 | 确保 AliyunFCDefaultRole 存在，该权限内容可以参考[这里](https://help.aliyun.com/document_detail/181589.html) |
| 资源/创建 | 确保 AliyunMNSNotificationRole 存在，该权限内容可以参考[这里](https://github.com/devsapp/fc/blob/main/docs/zh/yaml/triggers.md#%E8%A7%A6%E5%8F%91%E5%99%A8%E8%A7%92%E8%89%B2%E6%9D%83%E9%99%90-2)<br> AliyunMNSNotificationRole 的创建可参考[配置MNS主题触发器](https://help.aliyun.com/document_detail/164204.html) |

使用该项目，您需要准备好以下资源：

| 服务/业务 | MNS |     
| --- |  --- |   
| 资源/创建 | MNS 主题 |  

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/mns-topic-trigger-golang)
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
    - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为 quick-start-sample-codes。
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 mns-topic-trigger-golang。
    - MNS 主题资源所在区域 (mns topic region): 您需要提供您上述资源准备中创建的MNS 主题的所在区域，地域选项同上述 region 参数。默认值 cn-hangzhou (杭州)
    - 主题名 (topic name): 您需要提供您创建的 MNS topic 的名称。
    - 过滤标签 (filter tag): 可以跳过设置为空值，只有收到包含了此处设置的过滤标签字符串的消息时，才会触发函数执行。
    - 事件格式 (notify content format): 您需要Event格式,默认值为 JSON。
    - 重试策略 (notify strategy): 你需要选择重试策略，默认值为BACKOFF_RETRY，如何选择重试策略，请参见[NotifyStrategy](https://help.aliyun.com/document_detail/27481.htm?spm=a2c4g.11186623.0.0.27bd22354xlGLu#concept-2028914)。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=mns-topic-trigger-golang) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=mns-topic-trigger-golang)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init mns-topic-trigger-golang -d mns-topic-trigger-golang` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd mns-topic-trigger-golang && s deploy -y`
  
- 本地调试
  - 运行 `s invoke --event-file event.json` 进行本地调试
  - 文件 event.json 中的内容为模拟事件内容。
    ```bash
    {
      "TopicOwner":"topic account id",
      "Message":"mock mns message",
      "Subscriber":"subscriber account id",
      "PublishTime":1658235558094,
      "SubscriptionName":"test-5bf13c7e",
      "MessageMD5":"652BF0E6297840015247C3xxxxxxx",
      "TopicName":"fc-example",
      "MessageId":"3405CA51807661353B3xxxxxxxx"
    }
    ```
  - 调用函数时收到的响应如下所示:
    ```bash
    ========= FC invoke Logs begin =========
    FC Invoke Start RequestId: 250efb53-48ac-4e76-a6b2-dabde6a0xxxx
    2022-07-28T18:43:50.784Z 250efb53-48ac-4e76-a6b2-dabde6a0xxxx [INFO] main.go:25: event: map[Message:mock mns message MessageId:3405CA51807661353B3xxxxxxxx MessageMD5:652BF0E6297840015247C3xxxxxxx PublishTime:1.658235558094e+12 Subscriber:subscriber account id SubscriptionName:test-5bf13c7e TopicName:fc-example TopicOwner:topic account id]
    FC Invoke End RequestId: 250efb53-48ac-4e76-a6b2-dabde6a0xxxx
    Duration: 0.87 ms, Billed Duration: 1 ms, Memory Size: 128 MB, Max Memory Used: 9.82 MB
    ========= FC invoke Logs end =========
    FC Invoke instanceId: c-62e2d846-995c5020672547c6xxxx
    FC Invoke Result:
    "MessageBody:mock mns message"
    End of method: invoke
      ```
- 端对端测试
  - 登陆 MNS 控制台向主题发送一条消息
  ![img_1.jpg](https://cdn.jsdelivr.net/gh/penghuima/ImageBed@master/img/blog_file/PicGo-Github-ImgBed20220802120215.jpg)
  - 登陆函数计算控制台，找到刚才部署的函数，查看 `调用日志`, 如果没有开通日志请点击一键开通
  - 函数日志内容如下所示:
    ```bash
    2022-07-29 02:51:39FC Invoke Start RequestId: CEBD8A24-8980-594C-87A6-5E5882CBxxxx
    2022-07-29 02:51:392022-07-29 02:51:39 CEBD8A24-8980-594C-87A6-5E5882CBxxxx [INFO] main.go:25: event: map[Message:json meesage from MNS console MessageId:CC9C55A98076444C2B3790CF37CBxxxx MessageMD5:D2DE9F47F7987095172CF5956D85xxxx PublishTime:1.65903429934e+12 Subscriber:143199xxxxxxxxxx SubscriptionName:mns-af29c6cf TopicName:fc-example TopicOwner:143199xxxxxxxxxx]
    2022-07-29 02:51:39FC Invoke End RequestId: CEBD8A24-8980-594C-87A6-5E5882CBxxxx
    ```
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
|--- | --- | --- |
| <center>微信公众号：`serverless`</center> | <center>微信小助手：`xiaojiangwh`</center> | <center>钉钉交流群：`33947367`</center> | 

</p>

</devgroup>