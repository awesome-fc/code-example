# kafka-producer-fc-event-golang help documentation

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


## Preliminary preparation

### Permission preparation

Using this item, verify that your operational account has the following product permissions/policies:


| Service/Business     | Functional Computing                                         |
| -------------------- | ------------------------------------------------------------ |
| Permissions/Policies | AliyunFCFullAccess<br/>AliyunKafkaFullAccess<br/>AliyunVPCReadOnlyAccess |


### Resource preparation

  * For an available Kafka message queue, please refer to the official document of message queue Kafka version [Quick Start of Message Queue](https://help.aliyun.com/document_detail/99949.html).

    - Create a VPC private network (VPC is recommended in the production environment), please refer to [VPC official document](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg# task-1012575). VPC console [link](https://vpcnext.console.aliyun.com/). At this point, you can have a VPC and corresponding switches.

    > When deploying a Kafka instance, you will be prompted to create an available VPC private network

  * Create the Kafka Topic and Consumer Group to be used in the Kafka console.

### Environment preparation

Since Go's Kafka client program includes CGO, although Go has a cross-compiler, if the corresponding cross-compiled C toolchain is not installed, a cross-platform executable file cannot be directly compiled. That is, when we use CGO, in order to achieve cross-platform compilation, we need to make C/C++ code also support cross-platform compilation.

The environment for function computing is Linux/amd64, and other operating systems need corresponding environments to build corresponding executable files:

- Mac

  ```bash
  # Download the linux compilation toolchain
  brew install FiloSottile/musl-cross/musl-cross
  ````
  
- Windows/Mac

  ```bash
  # pull image
  docker pull karalabe/xgo-latest
  ````
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
  - Service name: You need to name your Function Compute service. The service name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-128, the default value is kafka-producer-quick-start.
  - Function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64. The default is kafka-producer-event-function-golang.
  - vpcId: We recommend that you use a VPC to access Kafka, and select the VPC used to create a Kafka instance. Note that you need to fill in the az supported by Function Compute.
  - vswitchIds: Use the vswitch id in the vpc to access kafka on the intranet. Note that az needs to be supported by Function Compute.
  - securityGroupId: The security group id of the vpc where the kafka instance is located, which can be found in the `ECS` console `Network and Security` menu item.
  - Kafka access point (bootstrapServers): the default access point for the Kafka instance you purchased, which can be found in the instance details
  - topicName: The topic name of the Kafka instance. The data production of this topic will trigger the deployment function, which needs to be created in advance.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

- :fire: via [Serverless Application Center](https://fcnext.console.aliyun.com/applications/create?template=kafka-producer-fc-event-golang),
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun. com/applications/create?template=kafka-producer-fc-event-golang) the application. (This method only supports Linux environment under this application. It is recommended to use [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) for deployment)

</appcenter>

- Deploy via [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install):

  - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install), and perform [Authorization Information Configuration](https://www.serverless-devs.com/ fc/config);
  - Initialize the project: `s init kafka-producer-fc-event-golang -d kafka-producer-fc-event-golang`
  - Fill in the parameters described in the above modules
  - Enter the project directory, `cd kafka-producer-fc-event-golang` select the field of `run` in the corresponding `pre-deploy` according to the usage environment.
  - Do project deployment: `s deploy -y`
- local debugging
  - Enter the application project project and execute the following command: `s invoke -e '{"Key": "test go kafka producer"}'`.
  - You can view the result after the simulated event triggers the function.

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

- End-to-end testing

  - Log in to the Function Compute console and enter the corresponding function to configure the test parameter as `{"key":"test kafka producer go"}` (note that the parameter needs to be configured as a json format with the key value of "Key" ["Key" is in The settings in the demo program can be modified]) Click `Test function`
  
  The returned results are as follows:

  ```bash
  "Finish sending the message to kafka: test kafka producer go!"
  ```
  The log is as follows:
  
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
  
  - Log in to the Kafka console, view the corresponding topic `message details` of the corresponding instance, find the corresponding partition, and then query the corresponding message.
  
  

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