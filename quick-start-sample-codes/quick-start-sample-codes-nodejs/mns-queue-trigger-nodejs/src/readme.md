# mns-queue-trigger-nodejs 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-queue-trigger-nodejs&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=mns-queue-trigger-nodejs" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-queue-trigger-nodejs&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=mns-queue-trigger-nodejs" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=mns-queue-trigger-nodejs&type=packageDownload">
  </a>
</p>

<description>

快速部署一个由 Nodejs 实现的消息队列 MNS 触发器（队列模型）示例函数到阿里云函数计算。

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
| 资源/创建 | MNS 队列 |  

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/mns-queue-trigger-nodejs)
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
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 mns-queue-trigger-nodejs。
    - 账户ID (account id): 您需要提供主账户的 ID。
    - 队列名 (queue name): 您需要提供您创建的 MNS queue 的名称。
    - 解码（base64 decode）: 您选择是否自动解码队列中的消息内容，true 即自动解码。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=mns-queue-trigger-nodejs) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=mns-queue-trigger-nodejs)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init mns-queue-trigger-nodejs -d mns-queue-trigger-nodejs` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd mns-queue-trigger-nodejs && s deploy -y`
  
- 使用 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 调试
  - 运行 `s invoke --event-file event.json` 进行调试
  - 文件 event.json 中的内容为模拟事件内容。
    ```bash
    {
      "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",
      "source":"MNS-Function-mnstrigger",
      "specversion":"1.0",
      "type":"mns:Queue:SendMessage",
      "datacontenttype":"application/json; charset=utf-8",
      "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",
      "time":"2021-04-08T06:28:17.093Z",
      "aliyunaccountid":"16490154********",
      "aliyunpublishtime":"2021-10-15T07:06:34.028Z",
      "aliyunoriginalaccountid":"164901546557****",
      "aliyuneventbusname":"MNS-Function-mnstrigger",
      "aliyunregionid":"cn-chengdu",
      "aliyunpublishaddr":"42.120.XX.XX",
      "data":{
          "requestId":"606EA3074344430D4C81****",
          "messageId":"C6DB60D1574661357FA227277445****",
          "messageBody":"TEST"
      }
    }
    ```
  - 调用函数时收到的响应如下所示:
    ```bash
    ========= FC invoke Logs begin =========
    FC Invoke Start RequestId: c73252bd-a1eb-4264-8c4c-e39da47b****
    2022-07-26T07:30:49.402Z c73252bd-a1eb-4264-8c4c-e39da47b**** [verbose] message publish time: 2021-10-15T07:06:34.028Z
    2022-07-26T07:30:49.402Z c73252bd-a1eb-4264-8c4c-e39da47b**** [verbose] message trigger time cost: 24539055.374 s
    2022-07-26T07:30:49.402Z c73252bd-a1eb-4264-8c4c-e39da47b**** [verbose] messageBody: TEST
    FC Invoke End RequestId: c73252bd-a1eb-4264-8c4c-e39da47b****
    Duration: 2.66 ms, Billed Duration: 3 ms, Memory Size: 128 MB, Max Memory Used: 43.81 MB
    ========= FC invoke Logs end =========
    FC Invoke instanceId: c-62df9727-fc2c2f592af241ef****
    FC Invoke Result:
    succ
    End of method: invoke
      ```
- 通过控制台调试
  - 登陆 MNS 控制台向队列发送一条消息
  ![img_1.jpg](https://cdn.jsdelivr.net/gh/penghuima/ImageBed@master/img/blog_file/PicGo-Github-ImgBed20220802120226.jpg)
  - 登陆函数计算控制台，找到刚才部署的函数，查看 `调用日志`, 如果没有开通日志请点击一键开通
  - 函数日志内容如下所示:
    ```bash
    2022-07-26 15:26:31FC Invoke Start RequestId: EB0A77CA8076444C6A50840F3D52****
    2022-07-26 15:26:31load code for handler:index.handler
    2022-07-26 15:26:312022-07-26 15:26:31 EB0A77CA8076444C6A50840F3D52**** [verbose] message publish time: 2022-07-26T07:26:31.604Z
    2022-07-26 15:26:312022-07-26 15:26:31 EB0A77CA8076444C6A50840F3D52**** [verbose] message trigger time cost: 0.244 s
    2022-07-26 15:26:312022-07-26 15:26:31 EB0A77CA8076444C6A50840F3D52**** [verbose] messageBody: message from mns console
    2022-07-26 15:26:31FC Invoke End RequestId: EB0A77CA8076444C6A50840F3D52****
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