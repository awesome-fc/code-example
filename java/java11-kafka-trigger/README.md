# Java11 Kafka消息队列触发器示例

本示例为您展示了 Java11  [消息队列Kafka](https://help.aliyun.com/document_detail/68151.html?spm=5176.167616.J_5253785160.5.2dfe6feexRPqMj) 消息读取示例。

本示例使用了Kafka消息队列作为示例，与函数计算中的消息队列Kafka版触发器一起实现了消息队列的读取。



 ## 准备开始
 - 一个可用的Kafka消息队列，可参考消息队列Kafka版官方文档[消息队列快速入门](https://help.aliyun.com/document_detail/99949.html)。

   - 创建VPC专有网络（推荐在生产环境中也使用VPC），可参考[VPC官方文档](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg#task-1012575)。VPC控制台[链接](https://vpcnext.console.aliyun.com/)。至此即可拥有VPC和相应交换机。

   > 部署Kafka实例时会提示创建可用的VPC专有网络

 - [可选] 安装并配置 Serverless Devs 工具。(https://help.aliyun.com/document_detail/195474.html)

 ## 快速开始

 ### 方式一. 使用控制台创建

#### 1. 安装依赖和部署代码包

```shell
# 安装依赖并编译为jar包，对应jar包将在target目录下
mvn clean package
```

 #### 2. 创建服务

推荐在Kafka实例相同Region创建服务。

创建服务时在`高级选项`中`服务角色`选择AliyunFcDefaultRole（如没有则根据提示创建相应角色），并开启`允许访问VPC`，选取创建Kafka实例时所选择的`专有网络`、`交换机`与对应的`安全组(Kafka实例部署后自动创建)`。

![CreateService.png](assets/CreateService.png)



 #### 3. 创建函数

 选择服务（或创建服务）后，单击创建函数，如图所示
 - 选择 `从零开始创建`
 - 填入函数名称
 - 选择运行环境 Java 11
 - 选择函数触发方式：通过事件请求触发
 - 配置触发器选择`消息队列Kafka版`
 - 其他设置使用默认

 > 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)

 #### 4. 配置Kafka Connector

在消息队列Kafka版控制台，选择`Connector任务列表`，选择对应的Kafka实例创建Connector。有关Connector详情可见(https://help.aliyun.com/document_detail/171242.html)。

- 输入名称
- 配置触发函数的Topic（即该Topic有消息时触发函数）
- 选择刚创建好的服务与函数
- 部署Connector

![Connector.png](assets/Connector.png)



 #### 5. **使用模拟数据测试**或直接在Connector`操作`中测试

 模拟数据：

```
[{"key":"testkey","offset":14,"overflowFlag":false,"partition":10,"timestamp":1659271844051,"topic":"HelloTopic","value":"Test Payload","valueSize":14}]
```

日志如下：

 ```bash
 c-62e67aa5-8501421476c2421c996e2022-07-31 20:50:45FunctionCompute python3 runtime inited.
 c-62e67aa5-8501421476c2421c996e2022-07-31 20:50:45FC Invoke Start RequestId: 789a7345-4340-47a7-b2f9-e529ff05fe68
 c-62e67aa5-8501421476c2421c996e2022-07-31 20:50:452022-07-31 20:50:45 789a7345-4340-47a7-b2f9-e529ff05fe68 [INFO] kafka whole message:[{"key":"testkey","offset":14,"overflowFlag":false,"partition":10,"timestamp":1659271844051,"topic":"HelloTopic","value":"Test Payload","valueSize":14}]
 c-62e67aa5-8501421476c2421c996e2022-07-31 20:50:452022-07-31 20:50:45 789a7345-4340-47a7-b2f9-e529ff05fe68 [INFO] message topic:HelloTopic
 c-62e67aa5-8501421476c2421c996e2022-07-31 20:50:452022-07-31 20:50:45 789a7345-4340-47a7-b2f9-e529ff05fe68 [INFO] message value:Test Payload
 c-62e67aa5-8501421476c2421c996e2022-07-31 20:50:45FC Invoke End RequestId: 789a7345-4340-47a7-b2f9-e529ff05fe68
 ```



> 目前Kafka触发器 Serverless devs工具支持不够完善，推荐使用控制台创建Kafka触发器完成消息的读取。
