# oss-trigger-fc-event-python3 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-event-python3&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=oss-trigger-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-event-python3&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=oss-trigger-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=oss-trigger-fc-event-python3&type=packageDownload">
  </a>
</p>

<description>

快速部署一个 Python 3.6 的 Event 类型的 OSS trigger 函数到阿里云函数计算。

</description>

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess <br>AliyunOSSFullAccess |

| 服务/业务 | 访问控制(RAM) |     
| --- |  --- |   
| 资源/创建 | 确保 AliyunFCDefaultRole 存在，该权限内容可以参考[这里](https://help.aliyun.com/document_detail/181589.html) |

使用该项目，您需要准备好以下资源：

| 服务/业务 | OSS |     
| --- |  --- |   
| 资源/创建 | OSS Bucket |  
| 资源/上传 | 图片文件 (JPG, PNG) 上传至 OSS Bucket |

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/oss-trigger-fc-event-python3)
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
    - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为 oss-trigger-quick-start。
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 oss-triger-event-function-python。
    - OSS Bucket 资源所在区域 (oss bucket region): 您需要提供您上述资源准备中创建的 OSS bucket 的所在区域，地域选项同上述 region 参数。默认值 cn-hangzhou (杭州)
    - 账户ID (account id): 您需要提供主账户的 ID，以便函数计算获悉您 OSS bucket 的位置。
    - Bucket 名 (bucket name): 您需要提供您创建的 OSS bucket 的名称。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-event-python3) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-event-python3)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init oss-trigger-fc-event-python3 -d oss-trigger-fc-event-python3` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd oss-trigger-fc-event-python3 && s deploy -y`
  
- 本地调试
  - 运行 `s cli fc-event oss` 生成 OSS Trigger 的 Event 样例参数
  - 生成的 Event 样例为，该 Event 为真实 OSS 触发传入 Event 的模拟。
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
  - 您需要将以上样例中 oss 和 region 资源部分进行替换， 进入目录 `cd event-template`, 对生成的 event 模板根据实际情况进行变更 `vim oss-event.json`
  - 其中的 regionName，accountId，bucketName 需要根据您创建的 bucket 信息进行替换， 请将 fileName 替换为您之前上传的图片名称。
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
  },
  "region": "${regionName}",
  ```
  - 使用测试样例进行触发测试，运行 `s cli fc invoke --service-name ${serviceNamme} --function-name ${functionName} --event-file event-template/oss-event.json --region ${regionName}`
  - 执行成功后您可以查询 OSS bucket，被处理后的图片会被放入 processed 文件夹中。
- 端对端测试
  - 登陆 OSS 控制台
  - 上传一个图片到您之前触发器的制定路径
  - 登陆函数计算控制台，找到刚才部署的函数，查看 `调用日志`, 如果没有开通日志请点击一键开通
  
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