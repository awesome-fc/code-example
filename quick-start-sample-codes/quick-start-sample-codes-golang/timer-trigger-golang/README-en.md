# timer-trigger-fc-event-golang help documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-golang&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=timer-trigger-fc-event-golang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-golang&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=timer-trigger-fc-event-golang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-golang&type=packageDownload">
  </a>
</p>



## Preliminary preparation

### Permission preparation

Using this item, verify that your operational account has the following product permissions/policies:


| Service/Business     | Functional Computing |
| -------------------- | -------------------- |
| Permissions/Policies | AliyunFCFullAccess   |

# Code & Preview

- [ :smiley_cat:source code](https://github.com/devsapp/)
- In order to successfully deploy this sample code, you need to provide the following parameters during the deployment process:
  - Region: You need to configure the region where your Function Compute service needs to be deployed through this parameter. The default value is cn-hangzhou (Hangzhou).
    - The geographic options available to you are:
      - cn-beijing (Beijing)
      - cn-hangzhou (Hangzhou)
      - cn-shanghai (Shanghai)
      - cn-qingdao (Qingdao)
      - cn-zhangjiakou (Zhangjiakou)
      - cn-huhehaote (Hohhot)
      - cn-shenzhen (Shenzhen)
      - cn-chengdu (Chengdu)
      - cn-hongkong (Hong Kong)
      - ap-southeast-1 (Singapore)
      - ap-southeast-2 (Sydney)
      - ap-southeast-3 (Kuala Lumpur)
      - ap-southeast-5 (Jakarta)
      - ap-northeast-1 (Tokyo)
      - eu-central-1 (Frankfurt)
      - eu-west-1 (London)
      - us-west-1 (Silicon Valley)
      - us-east-1 (Virginia)
      - ap-south-1 (Mumbai)
  - Service name: You need to name your Function Compute service. The service name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-128, and the default value is timer-trigger-quick-start.
  - Function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64. The default is timer-trigger-event-function-golang.
  - Time trigger expression (cronExpression): You need to fill in the expression to trigger the function. The time trigger expression supports two settings: @every, cron expression. For example: @every 1m / @every 1h30m.
  - Input content (payLoad): represents the input content of the trigger event itself.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

- :fire: Through [Serverless Application Center](https://fcnext.console.aliyun.com/applications/create?template=timer-trigger-fc-event-golang),
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun. com/applications/create?template=timer-trigger-fc-event-golang) the application.

</appcenter>

- Deploy via [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install):

  - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install), and perform [Authorization Information Configuration](https://www.serverless-devs.com/ fc/config);
  - Initialize project: `s init timer-trigger-fc-event-golang -d timer-trigger-fc-event-golang`
  - Fill in the parameters described in the above modules
  - Enter the project and deploy the project: `cd timer-trigger-fc-event-golang && s deploy -y`
- local debugging
  - Enter the application project project and execute the following command: `s invoke --event-file event-example/timer-fc-sample.json`.
  - You can view the logs and results after the simulated event triggers the function.

```bash
========= FC invoke Logs begin =========
2022/08/03 13:28:45.587867 start
FC Invoke Start RequestId: 975f927d-8830-4939-a636-db624401a8da
2022-08-03T13:28:45.603Z 975f927d-8830-4939-a636-db624401a8da [INFO] main.go:22: triggerName: TestTimer
2022-08-03T13:28:45.603Z 975f927d-8830-4939-a636-db624401a8da [INFO] main.go:23: triggerTime: 2022-07-29T10:02:58Z
2022-08-03T13:28:45.603Z 975f927d-8830-4939-a636-db624401a8da [INFO] main.go:24: payload: TestPayload
FC Invoke End RequestId: 975f927d-8830-4939-a636-db624401a8da

Duration: 1.66 ms, Billed Duration: 2 ms, Memory Size: 128 MB, Max Memory Used: 9.57 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62ea780d-30ea332513b3473eacdc

FC Invoke Result:
"Timer Payload: TestPayload"


End of method: invoke
````



- End-to-end testing

  - Log in to the Function Compute console, find the function you just deployed, and view the `call log` (if the log is not activated, click one-click activation), and you can view that the function trigger log is triggered and printed at the corresponding time interval.

  The log is as follows:

  ```bash
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:29:20FC Invoke Start RequestId: 2ee71a74-4266-43ec-8a6a-8d6e6dc06ddb
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:29:202022-08-03 21:29:20 2ee71a74-4266-43ec-8a6a-8d6e6dc06ddb [INFO] main.go:22: triggerName: timer
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:29:202022-08-03 21:29:20 2ee71a74-4266-43ec-8a6a-8d6e6dc06ddb [INFO] main.go:23:triggerTime:2022-08-03T13 :20Z
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:29:202022-08-03 21:29:20 2ee71a74-4266-43ec-8a6a-8d6e6dc06ddb [INFO] main.go:24: payload: TestPayload
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:29:20FC Invoke End RequestId: 2ee71a74-4266-43ec-8a6a-8d6e6dc06ddb
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:31:20FC Invoke Start RequestId: 49b019f1-747a-4955-9de6-1e931769a790
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:31:202022-08-03 21:31:20 49b019f1-747a-4955-9de6-1e931769a790 [INFO] main.go:22: triggerName: timer
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:31:202022-08-03 21:31:20 49b019f1-747a-4955-9de6-1e931769a790 [INFO] main.go:23:triggerTime: 23022-08-03 :20Z
  c-62ea780d-30ea332513b3473eacdc2022-08-03 21:31:202022-08-03 21:31:20 49b019f1-747a-4955-9de6-1e931769a790 [INFO] main.go:24: payload: TestPayload
  c-62

</deploy>

<appdetail id="flushContent">

# application details



This application is only used for learning and reference. You can carry out secondary development and improvement based on this project to realize your own business logic.



</appdetail>

<devgroup>

## Developer Community

If you have feedback about bugs or future expectations, you can give feedback and exchange in [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues). If you want to join our discussion group or keep up to date with the latest developments in FC components, you can do so through the following channels:

<p align="center">




| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| <center>WeChat public account:`serverless`</center>          | <center>WeChat assistant: `xiaojiangwh`</center>             | <center>Dingding exchange group:`33947367`</center>          |

</p>

</devgroup>