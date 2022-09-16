# rabbitmq-trigger-fc-event-springboot help documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rabbitmq-producer-fc-event-springboot&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rabbitmq-producer-fc-event-springboot" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rabbitmq-producer-fc-event-springboot&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rabbitmq-producer-fc-event-springboot" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rabbitmq-producer-fc-event-springboot&type=packageDownload">
  </a>
</p>

## Preliminary preparation

### Permission preparation

Using this item, verify that your operational account has the following product permissions/policies:


| Service/Business     | Functional Computing                                         |
| -------------------- | ------------------------------------------------------------ |
| Permissions/Policies | AliyunFCFullAccess<br/>AliyunAMQPFullAccess |


### Resource preparation

  * For an available RabbitMQ instance, please refer to the official document of Message Queue for RabbitMQ [create resources](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-resources).

  * Create the [Vhost](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-a-vhost) and [Queue](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-a-queue) in the details of rabbitmq instance.

# Code & Preview

- [ :smiley_cat:source code](https://github.com/devsapp/)
- In order to successfully deploy this sample code, you need to provide the following parameters during the deployment process:
  - Region: You need to configure the region where your Function Compute service needs to be deployed through this parameter. The default value is cn-qingdao (Qingdao).
    - The geographic options available to you are:
      - cn-beijing (Beijing)
      - cn-hangzhou (Hangzhou)
      - cn-shanghai (Shanghai)
      - cn-qingdao (Qingdao)
      - cn-zhangjiakou (Zhangjiakou)
      - cn-huhehaote (Hohhot)
      - cn-shenzhen (Shenzhen)
      - cn-chengdu (Chengdu)
      - cn-hongkong (Hongkong)
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
  - Service name: You need to name your Function Compute service. The service name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-128, the default value is rabbitmq-trigger-quick-start.
  - Function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64. The default is rabbitmq-trigger-event-function-springboot.
  - Instance ID: You need to specify the instance ID of the RabbitMQ. Please refer to [here](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-an-instance) for more details.
  - Virtual Host Name: You need to specify the virtual host name of the RabbitMQ. Please refer to [here](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-a-vhost) for more details.
  - Queue Name: You need to specify the queue name of the RabbitMQ. Please refer to [here](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/create-a-queue) for more details.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

- :fire: via [Serverless Application Center](https://fcnext.console.aliyun.com/applications/create?template=rabbitmq-trigger-fc-event-springboot),
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun. com/applications/create?template=rabbitmq-trigger-fc-event-springboot) the application.

</appcenter>

- Deploy via [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install):

  - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install), and perform [Authorization Information Configuration](https://www.serverless-devs.com/ fc/config);
  - Initialize the project: `s init rabbitmq-trigger-fc-event-springboot -d rabbitmq-trigger-fc-event-springboot`
  - Fill in the parameters described in the above modules
  - Enter the project and deploy the project: `cd rabbitmq-trigger-fc-event-springboot && s deploy -y`
- local debugging
  - Enter the application project project and execute the following command: `s invoke --event-file event-example/rabbitmq-eventbridge-fc-sample.json`.
  - You can view the logs and results after the simulated event triggers the function.


- End-to-end testing

  - Send messages to the RabbitMQ instance to trigger the invocation of the function. Please refer to [here](https://www.alibabacloud.com/help/en/message-queue-for-rabbitmq/latest/use-an-sdk-to-send-and-receive-messages) for more details.

  

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
| <center>WeChat Official Account: `serverless`</center>       | <center>WeChat Assistant: `xiaojiangwh`</center>             | <center>DingTalk Group:`33947367`</center>                   |

</p>

</devgroup>

