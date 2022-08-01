# FCToODPSSampleJava11 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToODPSSampleJava11&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=FCToODPSSampleJava11" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToODPSSampleJava11&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=FCToODPSSampleJava11" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToODPSSampleJava11&type=packageDownload">
  </a>
</p>

<description>

在阿里云函数计算上快速部署一个可以将数据插入到 ODPS(MaxCompute) 数据表的 java11 函数。

</description>


## 前期准备

### 权限准备

使用该项目，确认您的操作账户拥有以下的产品权限 / 策略：


| 服务/业务 | 函数计算 |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess</br>AliyunDataWorksFullAccess</br>AliyunBSSOrderAccess |     


### 资源准备

* 1: 开通 ODPS(MaxCompute) 服务。
* 2: 创建 ODPS 项目资源。**（如果您已经创建完成或有可以直接使用的项目资源，可跳过此步。）**
    * 2-1: 进入 ODPS 控制台首页，点击**创建项目**。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/52a09da0785fa1a473953c2a3598d57c/image.png)
    * 2-2: 填写**工作空间名称**，其他选项如无特殊需求可略过，完成后点击**创建项目**。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/eadd3641a31f7e03c3b6577f39e480c0/image.png)
    * 2-3: 填写**实例显示名称**，其他选项如无特殊需求可略过，完成后点击**确认创建**，此时创建 odps 项目的过程就完成了。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/4c0669d9eadb8cb1706294e21eb35920/image.png)
* 3: 创建 ODPS 数据表资源。**（如果您已经创建完成或有可以直接使用的数据表资源，可跳过此步。）**
    * 3-1: 进入 ODPS 控制台首页，选择 step2 创建的项目，并点击**数据开发**。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/4c122d3042ca9ebca83cb96df43081a8/image.png)
    * 3-2: 右键点击下图的**业务流程**，点击新建业务流程，并按照要求填写，此 demo 创建的业务流程名称为**fc_odps_sink_v1**。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/c704ead1235f64dc05d55ee032a224bf/image.png)
    * 3-3: 点击**新建**->**新建表**->**表**，配置表名称和业务路径，点击完成即进入到表配置页面。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/aa6d7c21f7923f5bc1e3b09e281ae9f6/image.png)
    * 3-4: 按下图设置中文名和表字段，完成后提交到生产环境。（注意：表字段是有序的，例如此 demo 表字段的顺序是 id,name,age。）
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/4ddcba884fcb5247dbd5c222e781297d/image.png)
    * 3-5: 点击左侧的**表管理**，刷新一下页面即可看到刚刚创建的数据表，此时创建 odps 数据表的过程就完成了。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/4b33ca9d55e94007c9cfbfee73a418ac/image.png)



## 使用步骤
### 应用部署

<appcenter>

- :fire: 通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=FCToODPSSampleJava11) ，
  [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=FCToODPSSampleJava11)  该应用。

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init FCToODPSSampleJava11 -d FCToODPSSampleJava11`
    - 进入项目，并进行项目部署：`cd FCToODPSSampleJava11 && s deploy -y`

</deploy>

无论通过上面哪种方式部署，都需要为应用配置参数，具体如下：
* accessKeyID/accessKeySecret: 账户的 ak 密钥对，用于在函数中访问 ODPS(MaxCompute)。
* odpsProject: ODPS(MaxCompute) 项目名称，需要字母开头，只能包含字母、下划线和数字。
* odpsEndpoint: ODPS(MaxCompute) 服务地址，如无特殊需求可选择外网访问地址，地址详见：https://help.aliyun.com/document_detail/34951.html。
* odpsTableName: ODPS(MaxCompute) 数据表名称。

### 应用调用
**调用参数**

应用部署完成后，可构造函数请求参数来对函数发起调用，测试函数正确性。 其示例可以参考：
```
[1, "xiaoming", 11]
```

**调用方式**
* 控制台调用：
    * 登录[函数计算控制台](https://fcnext.console.aliyun.com/cn-hangzhou/services) ，找到部署的函数。
    * 点击**测试函数**->**配置测试参数**，将上文的调用参数粘贴到下图所示位置中。
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/b1328a1411ea706bcb596a9e05193d62/image.png)
    ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/7371e9d5a5434dff730d95cecb40923f/image.png)
    * 完成后点击测试函数，则此条数据即通过函数插入到 ODPS 对应表中。

* s 工具调用：
    * 进入应用项目工程下，执行下面命令：`s invoke --event-file event-example/fc-to-odps-sample.json`。
    * 函数调用完成后，则此条数据即通过函数插入到 ODPS 对应表中。

<appdetail id="flushContent">

### 测试验证
函数调用成功后，该如何在 ODPS 侧查看数据是否写入成功？
* 1: 进入 ODPS 数据开发页面，按照下图所示点击**新建**->**新建节点**->**ODPS SQL**，创建完成后会弹出一个可以编写 sql 的界面，进入下一步。

  ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/d5f223b0484c1cb91b6b9a25ea6ec9e0/image.png)
* 2：写 sql 语句 `SELECT * FROM {tableName};`，点击运行 sql，完成后会返回表内的数据，您可检查表中的数据是否包含测试输入的 event 数据。
  ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/618742477e1b8a22fa953ab93381b8f1/image.png)


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