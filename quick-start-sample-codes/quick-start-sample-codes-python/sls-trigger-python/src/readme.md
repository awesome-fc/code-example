# sls-trigger-fc-event-python3 帮助文档

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=sls-trigger-fc-event-python3&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=sls-trigger-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=sls-trigger-fc-event-python3&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=sls-trigger-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=sls-trigger-fc-event-python3&type=packageDownload">
  </a>
</p>

<description>

快速部署一个 Python 3.6 的 Event 类型的 sls trigger 函数到阿里云函数计算，该函数会将从事件源日志库中拉取日志。

</description>


## 前期准备
使用该项目，推荐您的账号拥有以下的产品权限 / 策略：

| 服务/业务 | 函数计算(FC) |     
| --- |  --- |   
| 权限/策略 | AliyunFCFullAccess<br>AliyunLogFullAccess<br>AliyunLogInvokeFCAccess |  

使用该项目，您需要准备好以下资源：

| 服务/业务 | 日志服务(SLS) |     
| --- |  --- |   
| 资源/创建 | log project，日志项目 |  
| 资源/创建 | source log store，触发源日志库  |
| 资源/创建 | record log store，触发日志记录日志库  |

| 服务/业务 | 访问控制(RAM) |     
| --- |  --- |   
| 资源/创建 | service role name, 服务配置角色，请确保该角色拥有 AliyunLogFullAccess 权限 |  


<codepre id="codepre">

# 代码 & 预览

- [ :smiley_cat:  源代码](https://github.com/devsapp/start-fc/blob/main/event-function/sls-trigger-fc-event-python3)
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
    - 服务名 (service name): 您需要给您的函数计算服务进行命名，服务名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-128 之间，默认值为 sls-trigger-quick-start。
    - 函数名 (function name): 您需要给您的函数计算函数进行命名，函数名称，只能包含字母、数字、下划线和中划线。不能以数字、中划线开头。长度在 1-64 之间。默认值为 sls-triger-event-function-python。
    - 账户ID (account id): 您需要提供主账户的 ID，以便函数计算获悉您 sls 资源的位置。
    - 服务角色名（service role name）: 您需要提供服务的配置角色名称，请确保该角色拥有 AliyunLogFullAccess 权限
    - 日志项目名称 (log project): 您需要选择已创建的日志项目
    - 事件源日志库名称 (source log store): 您需要选择已创建的日志库提供事件源日志数据，触发器会定时从该日志仓库中订阅数据到函数服务进行自定义加工
    - 最大重试次数 (max retry time): 日志服务触发函数执行时，如果遇到错误，所允许的最大尝试次数，取值范围：[0,100]
    - 触发时间间隔 (trigger interval): 日志服务触发函数运行的时间间隔，取值范围：[3,600]，单位：秒
    - 记录日志库名称 (record log store): 您需要选择已创建的日志库来记录触发日志，日志服务触发函数执行过程的日志会记录到该日志仓库中


</codepre>

<deploy>

## 部署 & 体验

<appcenter>

-  :fire:  通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=sls-trigger-fc-event-python3) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=sls-trigger-fc-event-python3)  该应用。 

</appcenter>

- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
    - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://www.serverless-devs.com/fc/config) ；
    - 初始化项目：`s init sls-trigger-fc-event-python3 -d sls-trigger-fc-event-python3` 
    - 填入在以上模块介绍的参数
    - 进入项目，并进行项目部署：`cd sls-trigger-fc-event-python3 && s deploy -y`
  
- 代码测试

  #### 构造 sls event 参数进行模拟触发
  
  - 运行 `s cli fc-event sls` 生成 sls Trigger 的 Event 样例参数
  - 生成的 Event 样例为，该 Event 为真实 sls 触发传入 Event 的模拟。
    ```bash
    {
      "parameter": {},
      "source": {
        "endpoint": "http://cn-shanghai-intranet.log.aliyuncs.com",
        "projectName": "log-com",
        "logstoreName": "log-en",
        "shardId": 0,  # 日志分区信息
        "beginCursor": "MTUyOTQ4MDIwOTY1NTk3ODQ2Mw==",  # 起点游标，表示从什么位置开始读取数据
        "endCursor": "MTUyOTQ4MDIwOTY1NTk3ODQ2NA=="   # 结束游标，表示读取数据到什么地方结束
      },
      "jobName": "1f7043ced683de1a4e3d8d70b5a412843d817a39",
      "taskId": "c2691505-38da-4d1b-998a-f1d4bb8c9994",
      "cursorTime": 1529486425
    }
    ```
  - 您需要将以上样例中 sls 资源部分进行替换，其中的 `endpoint`, `projectName`, `logstoreName` 需要根据您创建的事件源日志库信息进行替换，`beginCursor`, `shardId`、 `endCursor` 需要基于实际写入的日志数据内容进行指定。
  - 使用测试样例进行触发测试，运行 `s cli fc invoke --service-name ${serviceNamme} --function-name ${functionName} --event-file event-template/test-sls.json --region ${regionName}`
  - 执行成功后您可以查看函数执行日志。
  
  #### 向事件源数据库写入日志数据进行真实触发
  
  - 通过 SLS put_logs 接口向事件源日志库写入日志数据，写入数据的 python 代码示例:
  ```python
  # encoding: utf-8
  
  import time
  from aliyun.log import *
  
  
  def main():
      endpoint = 'cn-shanghai.log.aliyuncs.com'       # 选择与上面步骤创建Project所属区域匹配的Endpoint
      accessKeyId = 'xxx'    # 使用你的阿里云访问密钥AccessKeyId
      accessKey = 'xxx'      # 使用你的阿里云访问密钥AccessKeySecret
      project = 'xxx'        # 上面步骤创建的项目名称
      logstore = 'xxx'       # 上面步骤创建的事件源日志库名称
  
      # 构建一个client
      client = LogClient(endpoint, accessKeyId, accessKey)
  
      # list 所有的logstore
      req1 = ListLogstoresRequest(project)
      res1 = client.list_logstores(req1)
      res1.log_print()
      topic = ""
      source = ""
  
      # 发送10个数据包，每个数据包有1条log
      for i in range(10):
          logitemList = []  # LogItem list
          for j in range(1):
              contents = [('index', str(i * 10 + j))]
              logItem = LogItem()
              logItem.set_time(int(time.time()))
              logItem.set_contents(contents)
              logitemList.append(logItem)
          req2 = PutLogsRequest(project, logstore, topic, source, logitemList)
          res2 = client.put_logs(req2)
          res2.log_print()
  
  if __name__ == '__main__':
      main()
  ```
  - 执行成功后您检查函数执行日志。

</deploy>

<appdetail id="flushContent">

# 应用详情



本应用仅作为学习和参考使用，您可以基于本项目进行二次开发和完善，实现自己的业务逻辑



</appdetail>

<devgroup>

## 开发者社区

您如果有关于错误的反馈或者未来的期待，您可以在 [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) 中进行反馈和交流。如果您想要加入我们的讨论组或者了解 FC 组件的最新动态，您可以通过以下渠道进行：

<p align="center">

| <img src="https://serverless-article-picture.sls-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.sls-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.sls-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
|--- | --- | --- |
| <center>微信公众号：`serverless`</center> | <center>微信小助手：`xiaojiangwh`</center> | <center>钉钉交流群：`33947367`</center> | 

</p>

</devgroup>