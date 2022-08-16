# timer-trigger-fc-event-nodejs14 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-nodejs14&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=timer-trigger-fc-event-nodejs14" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-nodejs14&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=timer-trigger-fc-event-nodejs14" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-nodejs14&type=packageDownload">
  </a>
</p>


## 前期准备

### 权限准备

使用该项目，确认您的操作账户拥有以下的产品权限 / 策略：


| 服务/业务 | 函数计算           |
| --------- | ------------------ |
| 权限/策略 | AliyunFCFullAccess |

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
  - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为 timer-trigger-quick-start。
  - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 timer-trigger-event-function-nodejs14。
  - 时间触发器表达式(cronExpression): 您需要填写表达式以触发函数，时间触发器表达式支持两种设置：@every、cron 表达。如：@every 1m / @every 1h30m。
  - 输入内容(payLoad): 代表触发器事件本身的输入内容。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=timer-trigger-fc-event-nodejs14) ，
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=timer-trigger-fc-event-nodejs14)  该应用。

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：

  - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
  - 初始化项目：`s init timer-trigger-fc-event-nodejs14 -d timer-trigger-fc-event-nodejs14`
  - 填入在以上模块介绍的参数
  - 进入项目，并进行项目部署：`cd timer-trigger-fc-event-nodejs14 && s deploy -y`
- 本地调试
  - 进入应用项目工程下，执行下面命令：`s invoke --event-file event-example/timer-fc-sample.json`。
  - 即可查看到模拟事件触发函数后的日志与结果。

```bash
========= FC invoke Logs begin =========
FC Invoke Start RequestId: 0e7355ba-98ac-4499-9d1d-fa35d2a01ec3
load code for handler:index.handler
2022-08-03T13:51:01.326Z 0e7355ba-98ac-4499-9d1d-fa35d2a01ec3 [verbose] whole event: {    "triggerTime": "2022-07-29T10:02:58Z",    "triggerName": "TestTimer",    "payload": "TestPayload"}
2022-08-03T13:51:01.326Z 0e7355ba-98ac-4499-9d1d-fa35d2a01ec3 [verbose] triggerName:  TestTimer
2022-08-03T13:51:01.326Z 0e7355ba-98ac-4499-9d1d-fa35d2a01ec3 [verbose] triggerTime:  2022-07-29T10:02:58Z
2022-08-03T13:51:01.326Z 0e7355ba-98ac-4499-9d1d-fa35d2a01ec3 [verbose] triggerMessgae:  TestPayload
FC Invoke End RequestId: 0e7355ba-98ac-4499-9d1d-fa35d2a01ec3

Duration: 3.53 ms, Billed Duration: 4 ms, Memory Size: 128 MB, Max Memory Used: 48.85 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62ea7d45-c4a7955f060a4cae9473

FC Invoke Result:
Timer Payload:TestPayload


End of method: invoke
```

​		

- 端对端测试

  - 登陆函数计算控制台，找到刚才部署的函数，查看 `调用日志`（如果没有开通日志请点击一键开通），即可查看到函数触发日志是以相应时间间隔触发打印的。
  
  日志如下：

  ```bash
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:52:30FC Invoke Start RequestId: 9c53e598-3032-4a70-8387-278073bafe93
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:52:302022-08-03 21:52:30 9c53e598-3032-4a70-8387-278073bafe93 [verbose] whole event: {"triggerTime":"2022-08-03T13:52:30Z","triggerName":"timer","payload":"TestPayload"}
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:52:302022-08-03 21:52:30 9c53e598-3032-4a70-8387-278073bafe93 [verbose] triggerName:  timer
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:52:302022-08-03 21:52:30 9c53e598-3032-4a70-8387-278073bafe93 [verbose] triggerTime:  2022-08-03T13:52:30Z
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:52:302022-08-03 21:52:30 9c53e598-3032-4a70-8387-278073bafe93 [verbose] triggerMessgae:  TestPayload
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:52:30FC Invoke End RequestId: 9c53e598-3032-4a70-8387-278073bafe93
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:54:30FC Invoke Start RequestId: 77817138-728c-4f7b-b37f-e3339fb4056c
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:54:302022-08-03 21:54:30 77817138-728c-4f7b-b37f-e3339fb4056c [verbose] whole event: {"triggerTime":"2022-08-03T13:54:30Z","triggerName":"timer","payload":"TestPayload"}
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:54:302022-08-03 21:54:30 77817138-728c-4f7b-b37f-e3339fb4056c [verbose] triggerName:  timer
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:54:302022-08-03 21:54:30 77817138-728c-4f7b-b37f-e3339fb4056c [verbose] triggerTime:  2022-08-03T13:54:30Z
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:54:302022-08-03 21:54:30 77817138-728c-4f7b-b37f-e3339fb4056c [verbose] triggerMessgae:  TestPayload
  c-62ea7d45-c4a7955f060a4cae94732022-08-03 21:54:30FC Invoke End RequestId: 77817138-728c-4f7b-b37f-e3339fb4056c
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
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| <center>微信公众号：`serverless`</center>                    | <center>微信小助手：`xiaojiangwh`</center>                   | <center>钉钉交流群：`33947367`</center>                      |

</p>

</devgroup>