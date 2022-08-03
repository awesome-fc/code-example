# rds-mysql-fc-event-python3 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-event-python3&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rds-mysql-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-event-python3&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rds-mysql-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-event-python3&type=packageDownload">
  </a>
</p>

<description>

快速部署一个 Python 3 的 Event 类型的读写 Mysql 数据库函数到阿里云函数计算。

</description>

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess |

| 服务/业务 | 访问控制(RAM) |     
| --- |  --- |   
| 资源/创建 | 确保 AliyunFCDefaultRole 存在，该权限内容可以参考[这里](https://help.aliyun.com/document_detail/181589.html) |

使用该项目，您需要准备好以下资源：

| 服务/业务 | 云数据库 RDS |     
| --- |  --- |   
| 资源/创建 | 云数据库RDS MySQL实例,如何创建可以参考[这里](https://help.aliyun.com/document_detail/26117.htm?spm=a2c4g.11186623.0.0.12a47634PzmWPx) |  

在 Mysql 数据库实例中参考以下sql语句创建表并插入测试数据
```sql
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `age` tinyint(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`id`, `name`, `age`) VALUES
(1, '张三', 18),
(2, '李四', 28);
```

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/rds-mysql-fc-event-python3)
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
    - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为 rds-mysql-quick-start。
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 rds-mysql-event-function-python。
    - 账户ID (account id): 您需要提供主账户的 ID。
    - Mysql 内网地址 (mysql url): 您需要提供 RDS 控制台 Mysql 实例的内网地址，用于连接数据库。
    - Mysql 端口号 (mysql port): 您需要提供 Mysql 端口号，用于连接数据库。
    - Mysql 数据库名称 (mysql database name): 您需要提供 Mysql 数据库名，用于连接数据库。
    - Mysql 用户名 (mysql user): 您需要提供 Mysql 用户名称，用于连接数据库。
    - Mysql 密码 (mysql password): 您需要提供 Mysql 密码，用于连接数据库。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=rds-mysql-fc-event-python3) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=rds-mysql-fc-event-python3)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init rds-mysql-fc-event-python3 -d rds-mysql-fc-event-python3` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd rds-mysql-fc-event-python3 && s deploy -y`
  
- 本地调试
  - 运行 `s invoke ` 进行本地调试
  - 调用函数时收到的响应如下所示:
    ```bash
    ========= FC invoke Logs begin =========
    FunctionCompute python3 runtime inited.
    FC Initialize Start RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    FC Initialize End RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    FC Invoke Start RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    2022-03-31T02:57:49.693Z 28fa11ab-81da-4cd0-b050-xxxxxxxxxx [INFO] (3, '王二', 38)
    FC Invoke End RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    Duration: 18.42 ms, Billed Duration: 19 ms, Memory Size: 128 MB, Max Memory Used: 34.80 MB
    ========= FC invoke Logs end =========
    FC Invoke Result:
    user: (3, '王二', 38)
    End of method: invoke
      ```
- 端对端测试
  - 登陆 FC 控制台并测试函数
  - 控制台返回结果如下所示:
    ```bash
    {name=王二, id=3, age=38}
    ```
- 数据库访问限制
  - 当使用云数据库时，一般都会有访问控制，比如阿里云数据库RDS中的白名单设置（ [RDS白名单设置说明](https://help.aliyun.com/document_detail/43185.html?spm=5176.19908528.help.dexternal.6c721450iLu0jH) )。如果仅仅作为测试，可以将白名单配置成 `0.0.0.0/0`。（不要在生产环境使用!)
  - 在生产环境，可以使用以下两种方式访问：
    - VPC方式（**推荐**） 
    参考文档：https://help.aliyun.com/document_detail/84514.html
    - 代理方式
    参考文档：https://help.aliyun.com/document_detail/91243.html

  - 本示例不是连接池方式，若要使用连接池，可以参考文档 [Connection Pooling with Connector/J](https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-usagenotes-j2ee-concepts-connection-pooling.html)

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