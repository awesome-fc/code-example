# oss-trigger-fc-event-nodejs14 help documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-event-nodejs14&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=oss-trigger-fc-event-nodejs14" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-event-nodejs14&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=oss-trigger-fc-event-nodejs14" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-event-nodejs14&type=packageDownload">
  </a>
</p>

<description>

Quickly deploy a Node.js 14 Event-type OSS trigger function to Alibaba Cloud Function Compute

</description>


## Preliminary preparation 
To use this program, it is recommended that you have the following product permissions/policies prepared:

| Service/Business | Function Compute   |
| --- |--------------------|
| Permissions/Policies | AliyunFCFullAccess |

To use this project, please prepare the following resources:

| Service/Business | OSS                                         |
|------------------|---------------------------------------------|
| Resource/Create  | OSS Bucket                                  |
| Resource/Upload  | Upload image file (JPEG, PNG) to OSS Bucket |

| Service/Business | Access control(RAM)                                                                                                                      |
| --- |------------------------------------------------------------------------------------------------------------------------------------------|
| Resource/Creat | Ensure AliyunFCDefaultRole exist，The content of this permission can refer to[there](https://help.aliyun.com/document_detail/181589.html) |


<codepre id="codepre">

# Code & Preview

- [ :smiley_cat: source code](https://github.com/devsapp/start-fc/blob/main/event-function/oss-trigger-fc-event-nodejs14)
- To deploy the sample code, you need to provide the following parameters during the deployment process:
  - region: You need to configure the region where your Function Compute service needs to be deployed through this parameter，The default value is cn-hangzhou.
    - The region options to choose from：
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
  - service name: You need to name your Function Compute service. The service name can only contain letters, numbers, underscores, and dashes. Cannot start with a number or a dash.Length between 1-128.The default value is oss-trigger-quick-start.
  - function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64.The default value is oss-trigger-event-function-nodejs.
  - oss bucket region:  You need to provide the region where the OSS bucket created in the resource preparation above is located. The region option is the same as the region parameter above. The default value is cn-hangzhou
  - account id: You need to provide the ID of the master account so that Function Compute can learn the location of your OSS bucket.
  - bucket name: You need to provide the name of the OSS bucket you created.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

-  :fire:  By [Serverless Application Center](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-event-nodejs14) ，
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-event-nodejs14)  this application.

</appcenter>

- Deploy by [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) ：
  - [Install Serverless Devs Cli ](https://www.serverless-devs.com/serverless-devs/install) ，and[Authorization information configuration](https://www.serverless-devs.com/fc/config) ；
  - Initialize the project：`s init oss-trigger-fc-event-nodejs14 -d oss-trigger-fc-event-nodejs14`
  - Fill in the parameters described in the above modules
  - Enter the project and deploy the project：`cd oss-trigger-fc-event-nodejs14 && s deploy -y`

- Test the code
  - Run `s cli fc-event oss` and generate OSS Trigger's Sample payload/event
  - An example of the generated Event is that the Event is a simulation of a real OSS triggering an incoming Event.
    ```bash
    {
      "events": [
        {
          "eventName": "ObjectCreated:PutObject",
          "eventSource": "acs:oss",
          "eventTime": "2017-04-21T12:46:37.000Z",
          "eventVersion": "1.0",
          "oss": {
            "bucket": {
              "arn": "acs:oss:cn-shanghai:123456789:bucketname",
              "name": "testbucket",
              "ownerIdentity": "123456789",
              "virtualBucket": ""
            },
            "object": {
              "deltaSize": 122539,
              "eTag": "688A7BF4F233DC9C88A80BF985AB7329",
              "key": "image/a.jpg",
              "size": 122539
            },
            "ossSchemaVersion": "1.0",
            "ruleId": "9adac8e253828f4f7c0466d941fa3db81161****"
          },
          "region": "cn-shanghai",
          "requestParameters": {
            "sourceIPAddress": "140.205.***.***"
          },
          "responseElements": {
            "requestId": "58F9FF2D3DF792092E12044C"
          },
          "userIdentity": {
            "principalId": "123456789"
          }
        }
      ]
    }
    ```
  - You need to replace the oss resource part in the above example. The regionName, accountId, and bucketName need to be replaced with your own bucket information. Please replace fileName with the image name you uploaded earlier.
  ```bash
  "oss": {
      "bucket": {
        "arn": "acs:oss:${regionName}:${accountId}:${bucketName}",
        "name": "${bucketName}",
        "ownerIdentity": "${accountId}",
        "virtualBucket": ""
      },
     "object": {
        "deltaSize": 122539,
        "eTag": "688A7BF4F233DC9C88A80BF985AB7329",
        "key": "${fileName}",
        "size": 122539
     },
  }
   "region": "${regionName}",
  ```
  - Trigger tests with test samples,run `s cli fc invoke --service-name ${serviceNamme} --function-name ${functionName} --event-file event-template/oss-event.json --region ${regionName}`
  - After successful execution, you can query the OSS bucket, and the backup images will be placed in the copy folder.

</deploy>

<appdetail id="flushContent">

# Application details



This application should be only used for learning and reference. You can carry out secondary development and improvement based on this project to realize your own business logic.



</appdetail>

<devgroup>

## Developer community

If you have feedback about errors or future expectations, you can give and exchange in  [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) .If you want to join our discussion group or keep up to date with the latest developments in FC components, you can do so through the following channels:

<p align="center">

| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
|--- | --- | --- |
| <center>WeChat public account：`serverless`</center> | <center>WeChat Assistant：`xiaojiangwh`</center> | <center>Dingding group：`33947367`</center> | 

</p>

</devgroup>