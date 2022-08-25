# start-python3-mongodb-http 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-python3-mongodb-http&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=start-python3-mongodb-http" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-python3-mongodb-http&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=start-python3-mongodb-http" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=start-python3-mongodb-http&type=packageDownload">
  </a>
</p>

<description>

本示例为您展示了 Python Runtime 的 MongoDB 使用示例。
在本示例中，表格存储实例配置在函数的环境变量配置中，initializer 回调函数从环境变量中获取配置，创建 MongoDB 连接，可以实现在整个函数实例生命周期内复用该连接，preStop 回调函数负责关闭 MongoDB 连接。

本示例 Driver 使用4.1.1版本。版本兼容详情见 https://www.mongodb.com/docs/drivers/pymongo/#compatibility

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

- :fire: 通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=start-python3-mongodb-http) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=start-python3-mongodb-http)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init start-python3-mongodb-http -d start-python3-mongodb-http`   
    - 进入项目，并进行项目部署：`cd start-python3-mongodb-http && s deploy -y`

</deploy>

<appdetail id="flushContent">

# 应用详情

## 初始化参数
| 参数名称     | 参数类型 | 是否必填 | 例子                                                     | 参数含义           |
| ------------ | -------- | -------- | -------------------------------------------------------- | ------------------ |
| serviceName  | String   | 选填     | start-python3-mongodb-http                                | 函数服务名称名     |
| functionName | String   | 选填     | start-python3-mongodb-http                                | 函数名称           |
| roleArn      | String   | 必填     | acs:*ram*::\<accountId>:role/aliyuncdnserverlessdevsrole | 函数执行角色       |
| MONGO_URL     | String   | 必填     | mongodb://xxxxxxxx | MongoDB 数据库地址 |  |
| MONGO_DATABASE     | String   | 必填     | table-instance                                           | 数据库名         |

### 调用测试

```shell
# 请求地址为 s deploy 返回的地址
curl https://******.cn-hangzhou.fcapp.run\?name\=张三
```

请求收到的响应如下所示：

```bash
{'_id': ObjectId('6306db41f15bf12e96341993'), 'name': '张三', 'age': 18.0}
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
