# runtime-extension-arms-java8-event 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=runtime-extension-arms-java8-event&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=runtime-extension-arms-java8-event" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=runtime-extension-arms-java8-event&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=runtime-extension-arms-java8-event" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=runtime-extension-arms-java8-event&type=packageDownload">
  </a>
</p>

<description>

快速部署一个由 Custom Java8 事件类型实现的ARMS应用程序监控扩展到阿里云函数计算。
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

| 服务/业务 | 资源 |     
| --- |  --- |   
| 资源/创建 | ARMS 应用 |

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/runtime-extension-arms)
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
    - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为quick-start-sample-codes。
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 runtime-extension-arms-java8-event。
    - Arms LicenseKey: Arms 提供的 LicenseKey，获取方法参考：https://help.aliyun.com/document_detail/195421.htm?spm=a2c4g.11186623.0.0.71426f66rCvF93#arms-cs-k8s-java

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=runtime-extension-arms-java8-event) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=runtime-extension-arms-java8-event)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
  - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
  - 初始化项目：`s init runtime-extension-arms-java8-event -d runtime-extension-arms-java8-event` 
  - 填入在以上模块介绍的参数
  - 进入项目，并进行项目部署：`cd runtime-extension-arms-java8-event && s deploy -y`
- 环境变量 `FC_EXTENSION_ARMS: true`（必须设置）
  - 添加该环境变量后，在一次函数调用结束时，不会立刻冻结函数实例，会等待10s再冻结函数实例，以确保Loggie Agent 扩展成功上报日志
  - 注意：该方式会有额外的收费，具体收费策略与实例Prefreeze回调相同，详见官方文档[函数实例生命周期-计费说明](https://help.aliyun.com/document_detail/203027.html#section-t95-ow2-tuf)
- 使用 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 调试
  - 运行 `s invoke -e "hello"` 进行远程调试，首次配置ARMS应用可能有一些延迟，建议多调用几次。
  - 调用后查看ARMS控制台应用监控信息
- 通过控制台调试
  - 登陆 FC 控制台并测试函数
  - 调用后查看ARMS控制台应用监控信息
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