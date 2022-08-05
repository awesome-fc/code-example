# start-java11-tablestore 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-java11-tablestore&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=start-java11-tablestore" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-java11-tablestore&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=start-java11-tablestore" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-java11-tablestore&type=packageDownload">
  </a>
</p>

<description>

本示例为您展示了 Java Runtime 的表格存储使用示例。
在本示例中，表格存储实例配置在函数的环境变量配置中，initializer 回调函数从环境变量中获取配置，创建表格存储客户端，可以实现在整个函数实例生命周期内复用该客户端。

</description>

<table>

## 准备开始
- 一个可用的表格存储实例，开通实例等教程可以查看阿里云官方文档。(https://help.aliyun.com/product/27278.html)
- 配置服务角色。

   当函数计算需访问阿里云其他云服务时，需要为函数计算授予相应的权限。详见文档(https://help.aliyun.com/document_detail/181589.html)。
   但是函数计算默认的服务角色 AliyunFCDefaultRole 不包含表格存储的权限，因此需要为该角色添加表格存储权限，也可以创建并使用新的角色。
   * 进入[RAM角色管理](https://ram.console.aliyun.com/roles)。
   * 找到 AliyunFCDefaultRole 并点击右侧**添加权限**。
   * 在**添加权限**页面，查找“表格存储”，选择 AliyunOTSFullAccess 权限，单击确定为角色添加表格存储权限。

</table>

<codepre id="codepre">

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

- :fire: 通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=start-java11-tablestore) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=start-java11-tablestore)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init start-java11-tablestore -d start-java11-tablestore`   
    - 进入项目，并进行项目部署：`cd start-java11-tablestore && s deploy -y`

</deploy>

<appdetail id="flushContent">

# 应用详情

## 初始化参数
| 参数名称     | 参数类型 | 是否必填 | 例子                                                     | 参数含义           |
| ------------ | -------- | -------- | -------------------------------------------------------- | ------------------ |
| serviceName  | String   | 选填     | start-java11-tablestore                                | 函数服务名称名     |
| functionName | String   | 选填     | start-java11-tablestore                                | 函数名称           |
| roleArn      | String   | 必填     | acs:*ram*::\<accountId>:role/aliyuncdnserverlessdevsrole | 函数执行角色       |
| endpoint     | String   | 必填     | https://\<instanceId>.<region>.ots-internal.aliyuncs.com | 表所在实例endpoint |  |
| instanceName     | String   | 必填     | table-instance                                           | 表所在实例         |

### 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
FC Invoke Start RequestId: d9******
FC Invoke End RequestId: d9******

Duration: 579.24 ms, Billed Duration: 580 ms, Memory Size: 128 MB, Max Memory Used: 64.28 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-******

FC Invoke Result:
{"primaryKey":[{"name":"region","value":"abc"},{"name":"id","value":1}],"attributes":[{"columnName":"age","columnValue":"20","timestamp":1657531733801},{"columnName":"home","columnValue":"北京","timestamp":1657618107569},{"columnName":"name","columnValue":"张三","timestamp":1657531733801}]}


End of method: invoke

```

</appdetail>

<devgroup>

## 开发者社区

您如果有关于错误的反馈或者未来的期待，您可以在 [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) 中进行反馈和交流。如果您想要加入我们的讨论组或者了解 FC 组件的最新动态，您可以通过以下渠道进行：

<p align="center">

| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| <center>微信公众号：`serverless`</center>                                                                                         | <center>微信小助手：`xiaojiangwh`</center>                                                                                        | <center>钉钉交流群：`33947367`</center>                                                                                           |

</p>

</devgroup>
