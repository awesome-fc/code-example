# timer-trigger-fc-event-python3 help documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-python3&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=timer-trigger-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-python3&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=timer-trigger-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=timer-trigger-fc-event-python3&type=packageDownload">
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
  - Function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64. The default is timer-trigger-event-function-python3.
  - Time trigger expression (cronExpression): You need to fill in the expression to trigger the function. The time trigger expression supports two settings: @every, cron expression. For example: @every 1m / @every 1h30m.
  - Input content (payLoad): represents the input content of the trigger event itself.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

- :fire: via [Serverless Application Center](https://fcnext.console.aliyun.com/applications/create?template=timer-trigger-fc-event-python3) ,
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun. com/applications/create?template=timer-trigger-fc-event-python3) the application.

</appcenter>

- Deploy via [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install):

  - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install), and perform [Authorization Information Configuration](https://www.serverless-devs.com/ fc/config);
  - Initialize the project: `s init timer-trigger-fc-event-python3 -d timer-trigger-fc-event-python3`
  - Fill in the parameters described in the above modules
  - Enter the project and deploy the project: `cd timer-trigger-fc-event-python3 && s deploy -y`
- local debugging
  - Enter the application project project and execute the following command: `s invoke --event-file event-example/timer-fc-sample.json`.
  - You can view the logs and results after the simulated event triggers the function.

```bash
========= FC invoke Logs begin =========
FunctionCompute python3 runtime initiated.
FC Invoke Start RequestId: c4022f76-ce36-4cf4-b86d-b08bacfafed7
2022-08-04T03:32:54.313Z c4022f76-ce36-4cf4-b86d-b08bacfafed7 [INFO] event: b'{ "triggerTime": "2022-07-29T10:02:58Z", "triggerName": "TestTimer" , "payload": "TestPayload"}'
2022-08-04T03:32:54.313Z c4022f76-ce36-4cf4-b86d-b08bacfafed7 [INFO] triggerName: TestTimer
2022-08-04T03:32:54.313Z c4022f76-ce36-4cf4-b86d-b08bacfafed7 [INFO] triggerTime = 2022-07-29T10:02:58Z
2022-08-04T03:32:54.313Z c4022f76-ce36-4cf4-b86d-b08bacfafed7 [INFO] payload = TestPayload
FC Invoke End RequestId: c4022f76-ce36-4cf4-b86d-b08bacfafed7

Duration: 3.19 ms, Billed Duration: 4 ms, Memory Size: 128 MB, Max Memory Used: 24.73 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62eb3de6-35638a6ddb9d4a0c8301

FC Invoke Result:
Timer Payload: TestPayload


End of method: invoke
````



- End-to-end testing

  - Log in to the Function Compute console, find the function you just deployed, and view the `call log` (if the log is not activated, click one-click activation), and you can view that the function trigger log is triggered and printed at the corresponding time interval.

  The log is as follows:

   ```bash
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:34:09FC Invoke Start RequestId: 7997cb0a-68dc-4a29-a8ad-279be13758ff
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:34:092022-08-04 11:34:09 7997cb0a-68dc-4a29-a8ad-279be13758ff [INFO] event: b'{"triggerTime":"2022-08-04T03:34:09Z","triggerName":"timer","payload":"TestPayload"}'
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:34:092022-08-04 11:34:09 7997cb0a-68dc-4a29-a8ad-279be13758ff [INFO] triggerName: timer
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:34:092022-08-04 11:34:09 7997cb0a-68dc-4a29-a8ad-279be13758ff [INFO] triggerTime = 2022-08-04T03:34:09Z
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:34:092022-08-04 11:34:09 7997cb0a-68dc-4a29-a8ad-279be13758ff [INFO] payload = TestPayload
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:34:09FC Invoke End RequestId: 7997cb0a-68dc-4a29-a8ad-279be13758ff
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:36:09FC Invoke Start RequestId: 9c2cf28f-395d-4d56-93cf-69a0aa026b7c
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:36:092022-08-04 11:36:09 9c2cf28f-395d-4d56-93cf-69a0aa026b7c [INFO] event: b'{"triggerTime":"2022-08-04T03:36:09Z","triggerName":"timer","payload":"TestPayload"}'
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:36:092022-08-04 11:36:09 9c2cf28f-395d-4d56-93cf-69a0aa026b7c [INFO] triggerName: timer
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:36:092022-08-04 11:36:09 9c2cf28f-395d-4d56-93cf-69a0aa026b7c [INFO] triggerTime = 2022-08-04T03:36:09Z
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:36:092022-08-04 11:36:09 9c2cf28f-395d-4d56-93cf-69a0aa026b7c [INFO] payload = TestPayload
c-62eb3de6-35638a6ddb9d4a0c83012022-08-04 11:36:09FC Invoke End RequestId: 9c2cf28f-395d-4d56-93cf-69a0aa026b7c

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



