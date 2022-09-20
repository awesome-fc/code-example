# rds-mysql-fc-http-java11 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-http-java11&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rds-mysql-fc-http-java11" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-http-java11&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rds-mysql-fc-http-java11" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-http-java11&type=packageDownload">
  </a>
</p>

<description>

快速部署一个 JAVA 11 的 HTTP 类型的读写 Mysql 数据库函数到阿里云函数计算。在本案例中提供公网方式连接到 RDS MySQL数据库。

</description>

## 前期准备
使用该项目，推荐您拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess |

使用该项目，您需要提前准备好 MySQL 数据库并执行以下 SQL 语句创建表:
  ```sql
    CREATE TABLE `users` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
      `name` varchar(20) NOT NULL,
      `age` tinyint(11) NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
  ```
 MySQL 数据库既可以选择阿里云 MySQL 数据库也可以选择其它 MySQL 数据库。在本案例中我们使用公网方式连接到阿里云 MySQL 数据库，若使用vpc方式连接阿里云 MySQL 数据库，请参考文档[配置网络](https://help.aliyun.com/document_detail/72959.html)配置VPC网络。
- 使用阿里云数据库RDS MySQL实例,创建教程可以参考[这里](https://help.aliyun.com/document_detail/26117.htm?spm=a2c4g.11186623.0.0.12a47634PzmWPx)
  - 登陆 RDS 控制台为 MySQL 实例[申请外网地址](https://help.aliyun.com/document_detail/26128.html),便于公网访问数据库
  - [设置 IP 白名单](https://help.aliyun.com/document_detail/96118.html),本案例作为测试，可以将白名单配置成 0.0.0.0/0。（不要在生产环境使用!)
  - 需要提供正确的数据库 URL 地址、数据库名称、用户、密码，用于连接数据库
- 使用其它 MySQL 数据库
  - 需要提供正确的数据库 URL 地址、数据库名称、用户、密码，用于连接数据库

<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/rds-mysql-fc-http-java11)
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
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 rds-mysql-event-function-java。
    - MySQL 地址 (mysql url):  MySQL 地址,可登陆RDS控制台->实例列表->基本信息->网络类型->外网地址 查看，用于连接数据库。
    - MySQL 端口号 (mysql port):  MySQL 端口号,默认3306。
    - MySQL 数据库名称 (mysql database name):  MySQL 数据库名,可登陆RDS控制台->实例列表->数据库管理->数据库名称 查看，用于连接数据库。
    - MySQL 用户名 (mysql user): MySQL 用户,可登陆RDS控制台->实例列表->账号管理->用户账号 查看，用于连接数据库。
    - MySQL 密码 (mysql password):  MySQL 密码,可登陆RDS控制台->实例列表->账号管理->用户账号 查看，用于连接数据库。

</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=rds-mysql-fc-http-java11) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=rds-mysql-fc-http-java11)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init rds-mysql-fc-http-java11 -d rds-mysql-fc-http-java11` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd rds-mysql-fc-http-java11 && s deploy -y`

## 应用详情
- 环境变量

| 参数名称 | 参数类型 | 是否必填 | 例子 |   参数含义 | 
| --- |  --- |  --- | --- | --- |
| MYSQL_ENDPOINT | String | 必填 | rm-uf6rrswxxxxxxxxxxxx.mysql.rds.aliyuncs.com |   数据库网址，用于连接数据库 |
| MYSQL_PORT | String | 必填 | 3306 |   数据库端口 | 
| MYSQL_DBNAME | String | 必填 | test |   数据库名称 | 
| MYSQL_USER | String | 必填 | fc |   数据库用户名 |  
| MYSQL_PASSWORD | String | 必填 | xxxxxxxx |   数据库密码 |  

- 代码测试
  - 通过postman构造http请求
  - 登陆函数控制台->服务->函数管理->触发器管理->配置信息->公网访问地址,粘贴公网地址
    ![img_1_公网地址](https://img.alicdn.com/imgextra/i3/O1CN01jD0eTR29hquBAB57V_!!6000000008100-2-tps-2976-820.png)
  - 在postman构造请求，粘贴公网访问地址，其中request body需包含想要插入数据库的 JSON 格式内容，包含用户名字和年龄字段。
    ```
    {
      "name": "wanger",
      "age": "12"
    }
    ```
  - 通过 POST 方法发送请求
  - 结果如下所示:
    ![img_2_postman结果](https://img.alicdn.com/imgextra/i4/O1CN01YOKX8J1zv07UNH5do_!!6000000006775-2-tps-2720-686.png)
如果用户不熟悉postman工具也可直接在控制台编辑测试用例
  ![](https://img.alicdn.com/imgextra/i2/O1CN01FgFm8Y1LLCWs8MLzE_!!6000000001282-2-tps-3428-1326.png)
- 数据库访问限制
  - 使用云数据库时，一般都会有访问控制，需要[设置 IP 白名单](https://help.aliyun.com/document_detail/96118.html),本案例作为测试，可以将白名单配置成 0.0.0.0/0。（不要在生产环境使用!)。
  - 在生产环境，可以使用以下两种方式访问：
    - VPC方式（**推荐**） 
    参考文档：[配置网络](https://help.aliyun.com/document_detail/72959.html)
    - 公网方式
    参考文档：[配置固定公网IP地址](https://help.aliyun.com/document_detail/410740.html)
  - 本示例不是连接池方式，若要使用连接池，可以参考文档 [Connection Pooling with Connector/J](https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-usagenotes-j2ee-concepts-connection-pooling.html)

本应用仅作为学习和参考使用，您可以基于本项目进行二次开发和完善，实现自己的业务逻辑

</deploy>

<appdetail id="flushContent">

# 常见问题
- 未设置白名单，MySQL 网址或端口设置错误
    ```bash
     "errorMessage": "Communications link failure\n\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server."
    ```
- MySQL 用户名、密码错误
    ```bash
     "errorMessage": "Access denied for user 'fc'@'120.76.207.131' (using password: YES)"
    ```
- MySQL 数据库名称错误
    ```bash
     "errorMessage": "Unknown database 'users1'" 
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