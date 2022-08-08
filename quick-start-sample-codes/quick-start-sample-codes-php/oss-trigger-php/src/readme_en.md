# oss-trigger-fc-http-php help documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-http-php&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=oss-trigger-fc-http-php" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-http-php&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=oss-trigger-fc-http-php" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-http-php&type=packageDownload">
  </a>
</p>

<description>

Quickly deploy a php http type oss trigger function to Alibaba Cloud Function Computing, this function will backup http request to object to copy directory

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

| Service/Business | Access control(RAM) |
| -- |  --- |
| Resource/Create | Ensure AliyunFCDefaultRole exist，The content of this permission can refer to[here](https://help.aliyun.com/document_detail/181589.html) |


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

-  :fire:  By [Serverless application center](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-http-php) ，
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-http-php)  this application.

</appcenter>

-  Deploy by [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) ：
  - [Install Serverless Devs Cli ](https://www.serverless-devs.com/serverless-devs/install) ，and[Authorization information configuration](https://www.serverless-devs.com/fc/config) ；
  - Initialize the project： `s init oss-trigger-fc-event-nodejs14 -d oss-trigger-fc-event-nodejs14`
  - Fill in the parameters described in the above modules
  - Enter the project and deploy the project：`cd oss-trigger-fc-event-nodejs14 && s deploy -y`

- Code testing
  #### Construct http request through postman
    - Go to the Function Compute console to find the deployed function and copy its public network access address
      ![](https://img.alicdn.com/imgextra/i1/O1CN01R7fWRr1xDN9b6e0fa_!!6000000006409-0-tps-2047-607.jpg)
    - Construct a request in postman, paste the public network access address, and the request body must contain the following content
    ```bash
    {
    "endpoint": "xxx",  // The endpoint of the target OSS
    "bucket": "xxx", //bucket name
    "object": "exampledir/examplefile.jpeg" // full path to the file to be manipulated
    }
  ```
  - Send request via GET
  ![](https://img.alicdn.com/imgextra/i2/O1CN01yVwHJr26EaksATfmO_!!6000000007630-0-tps-1706-776.jpg)
  - View the postman request result, and at the same time, you can view the target OSS. The object has been backed up to the copy directory.

  
    

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