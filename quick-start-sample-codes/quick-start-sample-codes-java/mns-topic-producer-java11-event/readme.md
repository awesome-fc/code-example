# mns-topic-producer-java11-event 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-topic-producer-java11-event&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=mns-topic-producer-java11-event" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-topic-producer-java11-event&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=mns-topic-producer-java11-event" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-topic-producer-java11-event&type=packageDownload">
  </a>
</p>

<description>

快速部署一个[消息服务MNS](https://help.aliyun.com/document_detail/27414.html) 主题模型生产者示例函数到阿里云函数计算，与函数计算中的 MNS 主题触发器一起实现了消息服务的生产者-消费者模型。其中MNS 主题触发器函数可查看应用 [mns-topic-trigger-fc-event-java11](http://www.devsapp.cn/details.html?name=mns-topic-trigger-fc-event-java11)

</description>

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess <br> AliyunMNSFullAccess |

| 服务/业务 | 访问控制(RAM) |     
| --- |  --- |   
| 资源/创建 | 确保 AliyunFCDefaultRole 存在，该权限内容可以参考[这里](https://help.aliyun.com/document_detail/181589.html) |

使用该项目，您需要准备好以下资源：

| 服务/业务 | MNS |     
| --- |  --- |   
| 资源/创建 | MNS 主题 |  

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/mns-topic-producer-java11-event)
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
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 mns-topic-producer-java11-event。
    - 账户ID (account id): 您需要提供主账户的 ID。
    - 主题名 (queue name): 您需要提供您创建的 MNS 主题的名称。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=mns-topic-producer-java11-event) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=mns-topic-producer-java11-event)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init mns-topic-producer-java11-event -d mns-topic-producer-java11-event` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd mns-topic-producer-java11-event && s deploy -y`
  
- 使用 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 调试
  - 运行 `s invoke` 进行远程调试
  - 调用函数时收到的响应如下所示:
    ```bash
    ========= FC invoke Logs begin =========
    FC Initialize Start RequestId: 48941a58-7e8d-4b2f-8ddc-xxxxxxxxxx
    2022-07-19 08:44:41.161 [INFO] [48941a58-7e8d-4b2f-8ddc-xxxxxxxxxx] init mns client time cost: 504ms
    FC Initialize End RequestId: 48941a58-7e8d-4b2f-8ddc-xxxxxxxxxx
    FC Invoke Start RequestId: c814ae3f-0ad6-409e-a4d0-xxxxxxxxxx
    2022-07-19 08:44:52.080 [INFO] [c814ae3f-0ad6-409e-a4d0-xxxxxxxxxx] publish message succ, message id:3405CA51807661357FC76xxxxxxxxxx, message body md5:D56E8E0A94F7B385AE1xxxxxxxxxx
    FC Invoke End RequestId: c814ae3f-0ad6-409e-a4d0-xxxxxxxxxx
    Duration: 348.41 ms, Billed Duration: 349 ms, Memory Size: 256 MB, Max Memory Used: 128.43 MB
    ========= FC invoke Logs end =========
    FC Invoke instanceId: c-62d66ef7-5bcb408bxxxxxxxxxx
    FC Invoke Result:
    publish message succ, message id:3405CA51807661357FC76xxxxxxxxxx, message body md5:D56E8E0A94F7B385AE1xxxxxxxxxx
    End of method: invoke
      ```
- 端对端测试
  - 登陆 FC 控制台并测试函数
  - 控制台返回结果如下所示:
    ```bash
    publish message succ, message id:3405CA51807661357FC76xxxxxxxxxx, message body 
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