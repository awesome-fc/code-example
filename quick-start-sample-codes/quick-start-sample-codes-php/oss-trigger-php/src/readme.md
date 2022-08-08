# oss-trigger-fc-http-php 帮助文档

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

快速部署一个 php 的 http 类型的 oss trigger 函数到阿里云函数计算，该函数会将备份http request中到object到copy目录下

</description>

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |
| --- |  --- |
| 权限/策略 | AliyunFCFullAccess |

使用该项目，您需要准备好以下资源：

| 服务/业务 | OSS                             |
| --- |---------------------------------|
| 资源/创建 | OSS Bucket                      |
| 资源/上传 | 图片文件 (JPEG, PNG) 上传至 OSS Bucket |

| 服务/业务 | 访问控制(RAM) |
| --- |  --- |
| 资源/创建 | 确保 AliyunFCDefaultRole 存在，该权限内容可以参考[这里](https://help.aliyun.com/document_detail/181589.html) |  

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/http-function/oss-trigger-fc-http-php)
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
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 oss-triger-http-function-php。
    - 账户ID (account id): 您需要提供主账户的 ID，以便函数计算获悉您 oss 资源的位置。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-http-php) ，
   [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=oss-trigger-fc-http-php)  该应用。

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init oss-trigger-fc-http-php -d oss-trigger-fc-http-php`
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd oss-trigger-fc-http-php && s deploy -y`

- 代码测试
  #### 通过postman构造http请求
    - 到函数计算控制台找到部署的函数，复制其公网访问地址
      ![](https://img.alicdn.com/imgextra/i1/O1CN01R7fWRr1xDN9b6e0fa_!!6000000006409-0-tps-2047-607.jpg)
    - 在postman构造请求，粘贴公网访问地址，其中request body需包含如下内容
    ```bash
    {
    "endpoint": "xxx",  // 目标OSS的endpoint
    "bucket": "xxx", //bucket 名称
    "object": "exampledir/examplefile.jpeg" // 欲操作的文件的完整路径
    }
  ```
  - 通过GET方式发送请求
  ![](https://img.alicdn.com/imgextra/i2/O1CN01yVwHJr26EaksATfmO_!!6000000007630-0-tps-1706-776.jpg)
  - 查看postman请求结果，同时可以到目标OSS查看，object已被备份到copy目录下

  
    

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