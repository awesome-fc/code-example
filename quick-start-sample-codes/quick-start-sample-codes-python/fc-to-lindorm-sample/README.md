# FCToLindormPython3 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToLindormPython3&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=FCToLindormPython3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToLindormPython3&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=FCToLindormPython3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToLindormPython3&type=packageDownload">
  </a>
</p>

<description>

在阿里云函数计算上快速部署一个可以在 lindorm 宽表引擎上建表插入数据的函数。

</description>


## 前期准备

### 权限准备

使用该项目，确认您的操作账户拥有以下的产品权限 / 策略：


| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess</br>AliyunLindormFullAccess</br>AliyunVPCFullAccess |     


### 资源准备

  * 1：开通 Lindorm 服务。
  * 2：创建 Lindorm 实例。
    点击实例列表的创建，并按照下图填写对应选项。（记住此处填写的 vpcID 和 vswitchID，下文部署函数时也需填写此参数。）
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/ffd1cc739fade8edbb435f3e13f73e67/image.png)
  * 3：为 Lindorm 设置访问白名单： 未设置白名单是无法访问的，Lindorm 出于安全角度考虑禁止设置 0.0.0.0/0(即向所有客户端开放)，可将 vswitch 的 IPV4 网段设为访问白名单。
    * 3-1：到专有网络控制台查看上步配置的交换机，复制下图中的 IPV4 网段。
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/25b4fb29d4f5ab82838dc4e1cdec8da5/image.png)
    * 3-2：点击 Lindorm 控制台的访问控制，将复制的网段粘贴到白名单分组中。
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/d5ad661e0648050380a66909b9c712d5/image.png)

## 使用步骤
### 应用部署

<appcenter>

- :fire: 通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=FCToLindormPython3) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=FCToLindormPython3)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init FCToLindormPython3 -d FCToLindormPython3`   
    - 进入项目，并进行项目部署：`cd FCToLindormPython3 && s deploy -y`

</deploy>


无论通过上面哪种方式部署，都需要为应用配置参数，具体如下：
  * 函数计算 VPC 配置
    * vpcID：函数实例所在的 vpc id，需和 lindorm 所在的 vpc 相同。
    * vswitchID：vpc 中 vswitch id，用于专有网络访问 lindorm。
    * securityGroupID：vpc 下安全组 id，用于专有网络访问 lindorm。（如果没有可创建一个，出方向和入方向使用默认规则即可。）
  * lindormUserName/lindormPassword：访问 lindorm 宽表引擎的账户密码，默认值均为 root。（可到 `lindorm 控制台->数据库连接->宽表引擎 ` 中查看。）
  * databaseURL：访问 lindorm server 的专有网络地址。（可到 `lindorm 控制台->数据库连接->宽表引擎 ` 中查看。）
    注意：只截取 `http://` 的后半段即可，如下图。
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/3dd10276fdd9a6f64e0e8687dc6a7a30/image.png)
  * sqlTableName: 函数在宽表引擎创建的 table 表名，函数会自动在宽表引擎中创建此表并插入数据。

### 应用调用
**调用参数**

不需要填写任何参数。

**调用方式**
  * 控制台调用：
    * 登录[函数计算控制台](https://fcnext.console.aliyun.com/cn-hangzhou/services) ，找到部署的函数。
    * 点击**测试函数**，请求执行成功即可
    * 你可以在下图的日志输出中查看 info 日志，如果 info 日志输出的 row 数据和插入的数据一样，证明测试成功。
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/6c15682b95e3371b1cb4d2afad58372b/image.png)

  
  * s 工具调用：
    * 进入应用项目工程下，执行下面命令：`s invoke`。
    * 你可以在下图的日志输出中查看 info 日志，如果 info 日志输出的 row 数据和插入的数据一样，证明测试成功。
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/58677f84041df3f2fd875b32029faa10/image.png)


<appdetail id="flushContent">


## 高级功能
出于对下游服务的保护，此应用默认的函数实例并发上限为 10，如果您需要更高的并发，可自主配置提高限制。

## 应用详情

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