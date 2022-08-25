# start-php-mongodb 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-php-mongodb&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=start-php-mongodb" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-php-mongodb&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=start-php-mongodb" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-php-mongodb&type=packageDownload">
  </a>
</p>

<description>

本示例为您展示了 PHP Runtime 的 MongoDB 使用示例。
在本示例中，表格存储实例配置在函数的环境变量配置中，initializer 回调函数从环境变量中获取配置，创建 MongoDB 连接，可以实现在整个函数实例生命周期内复用该连接。

MongoDB扩展是函数计算 PHP 环境的非内置拓展，需要自行安装，详见[文档](https://help.aliyun.com/document_detail/89032.html)。本示例已完成安装，扩展位于 src/code/extension 目录下。

</description>

<table>

## 准备开始
- 一个可用的 MongoDB 数据库，可以参考以下命令创建测试数据库。

```bash
use users
db.users.insert([
  {"name": "张三", "age": 18},
  {"name": "李四", "age": 20}
])
```

</table>

<codepre id="codepre">

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

- :fire: 通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=start-php-mongodb) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=start-php-mongodb)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init start-php-mongodb -d start-php-mongodb`   
    - 进入项目，并进行项目部署：`cd start-php-mongodb && s deploy -y`

</deploy>

<appdetail id="flushContent">

# 应用详情

## 初始化参数
| 参数名称     | 参数类型 | 是否必填 | 例子                                                     | 参数含义           |
| ------------ | -------- | -------- | -------------------------------------------------------- | ------------------ |
| serviceName  | String   | 选填     | start-php-mongodb                                | 函数服务名称名     |
| functionName | String   | 选填     | start-php-mongodb                                | 函数名称           |
| roleArn      | String   | 必填     | acs:*ram*::\<accountId>:role/aliyuncdnserverlessdevsrole | 函数执行角色       |
| MONGO_URL     | String   | 必填     | mongodb://xxxxxxxx | MongoDB 数据库地址 |  |
| MONGO_DATABASE     | String   | 必填     | table-instance                                           | 数据库名         |

### 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
s invoke

========= FC invoke Logs begin =========
FC Invoke Start RequestId: fd24aa39******
\nFC Invoke End RequestId: fd24aa39******

Duration: 80.16 ms, Billed Duration: 81 ms, Memory Size: 256 MB, Max Memory Used: 12.85 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62ff777d******

FC Invoke Result:
[{"_id":{"$oid":"62ff76916e3ab64609eebb99"},"name":"张三","age":18}]


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
