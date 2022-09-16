# Java11 Kafka message queue trigger example

This example shows you a Java11 [Message Queue Kafka](https://help.aliyun.com/document_detail/68151.html?spm=5176.167616.J_5253785160.5.2dfe6feexRPqMj) message reading example.

This example uses the Kafka message queue as an example, and implements the reading of the message queue together with the message queue Kafka version trigger in Function Compute.



 ## ready to start

 - For an available Kafka message queue, please refer to the official document of message queue Kafka version [Quick Start of Message Queue](https://help.aliyun.com/document_detail/99949.html).

   - Create a VPC private network (VPC is recommended in the production environment), please refer to [VPC official document](https://help.aliyun.com/document_detail/65398.htm?spm=a2c4g.11186623.0.0.61be4c9d4aGfpg# task-1012575). VPC console [link](https://vpcnext.console.aliyun.com/). At this point, you can have a VPC and corresponding switches.

   > When deploying a Kafka instance, you will be prompted to create an available VPC private network

 - [Optional] Install and configure Serverless Devs tools. (https://help.aliyun.com/document_detail/195474.html)

 ## quick start

 ### Method 1. Use the console to create

#### 1. Install dependencies and deploy code packages

```shell
# Install dependencies and compile them into jar packages, the corresponding jar packages will be in the target directory
mvn clean package
````

 #### 2. Create service

It is recommended to create services in the same Region of the Kafka instance.

 #### 3. Create function

 After selecting a service (or creating a service), click Create Function as shown

 - Select `Create from scratch`
 - Fill in the function name
 - Select the runtime environment Java 11
 - Select function trigger method: trigger by event request
 - No selection is made at the configuration trigger
 - Use default for other settings

 > For the detailed function creation process, see the document: [Create a function using the console](https://help.aliyun.com/document_detail/51783.html)



 #### 4. Configure EventBridge (you can configure it when you create a function in the future)

In the EventBridge console (https://eventbridge.console.aliyun.com/), select `Event Stream`, `Create Event Stream`:

- enter name
- Select the corresponding Kafka instance, topic and group. (that is, the Topic triggers the function when a message enters)
- news site select the latest site
- Network configuration select default network

<img src="assets/CreateEventBridge.png" alt="CreateEventBridge.png" style="zoom:50%;" />

- The rule defaults to `{}` empty object
- Select `Function Calculation` as the target, and select the corresponding trigger service and function.

<img src="assets/Objective.png" alt="Objective" style="zoom:50%;" />

After the creation is successful, `enable` the event flow can wait for the startup to succeed.



 #### 5. Test the test with mock data

 The simulated data is the real Event trigger data:

````
["{\"data\":{\"topic\":\"HelloTopic\",\"partition\":3,\"offset\":3,\"timestamp\":1659355857017,\"headers\ ":{\"headers\":[],\"isReadOnly\":false},\"value\":\"b\\u0027{\\\\n \\\"Test\\\\": \ \\"TestKafkaEBtrigger2\\"\\\\n}\\u0027\"},\"id\":\"91389190-9b3f-4df8-bebd-bb32433a02fe\",\"source\":\"acs :alikafka\",\"specversion\":\"1.0\",\"type\":\"alikafka:Topic:Message\",\"datacontenttype\":\"application/json; charset\\u003dutf- 8\",\"time\":\"2022-08-01T12:10:57.017Z\",\"subject\":\"acs:alikafka:alikafka_pre-cn-7pp2t2jwj001:topic:HelloTopic\",\ "aliyunaccountid\":\"1938858730552836\"}"]
````

The log is as follows:

 ```bash
c-62e7c2d2-c36b80589fb24d8484812022-08-01 20:10:58FC Invoke Start RequestId: ad215848-8f9f-4432-b3dc-3cd8a0c59d77
c-62e7c2d2-c36b80589fb24d8484812022-08-01 20:10:59[Name] Register [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@58372a00] as [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@com.aliyun .serverless.runtime.classloader.FunctionClassLoader@]: hash [f2bdbf56] (normal mode)
c-62e7c2d2-c36b80589fb24d8484812022-08-01 20:10:592022-08-01 12:10:59.035 [INFO] [ad215848-8f9f-4432-b3dc-3cd8a0c59d77] Event: ["{\"data\":{\ "topic\":\"HelloTopic\",\"partition\":3,\"offset\":3,\"timestamp\":1659355857017,\"headers\":{\"headers\":[] ,\"isReadOnly\":false},\"value\":\"b\\u0027{\\\\n \\\"Test\\\": \\\"TestKafkaEBtrigger2\\\\"\\\\ \n}\\u0027\"},\"id\":\"91389190-9b3f-4df8-bebd-bb32433a02fe\",\"source\":\"acs:alikafka\",\"specversion\": \"1.0\",\"type\":\"alikafka:Topic:Message\",\"datacontenttype\":\"application/json; charset\\u003dutf-8\",\"time\":\ "2022-08-01T12:10:57.017Z\",\"subject\":\"acs:alikafka:alikafka_pre-cn-7pp2t2jwj001:topic:HelloTopic\",\"aliyunaccountid\":\"1938858730552836\"} "]
Kafka Topic: HelloTopic
c-62e7c2d2-c36b80589fb24d8484812022-08-01 20:10:592022-08-01 12:10:59.242 [INFO] [ad215848-8f9f-4432-b3dc-3cd8a0c59d77] Message Value: b'{\n "Test": " TestKafkaEBtrigger2"\n}'
c-62e7c2d2-c36b80589fb24d8484812022-08-01 20:10:59FC Invoke End RequestId: ad215848-8f9f-4432-b3dc-3cd8a0c59d77
 ````



 ### Method 2. Compile and deploy using Serverless Devs tools

 #### 1. Modify s.yaml configuration

- Modify region, serviceName, functionName (set the same region as the Kafka instance).

- Modify the triggers configuration, fill in the Kafka InstanceId, ConsumerGroup, and Topic of the trigger function (all need to be created in advance), and finally set the consumption location to the latest location (latest) or the earliest location (earliest).


 #### 2. Install dependencies and deploy

 Install dependent libraries

 ```shell
# Use the s tool to install dependencies, you need to use docker
s build --use-docker
 ````

 deploy code

 ```bash
s deploy -y
 ````

 #### 3. Invoke the test

Trigger data tests with real events.

 ```shell
s invoke -e '["{\"data\":{\"topic\":\"HelloTopic\",\"partition\":9,\"offset\":3,\"timestamp\":1659346376797 ,\"headers\":{\"headers\":[],\"isReadOnly\":false},\"value\":\"b\\u0027{\\\\n \\\"Test\ \\": \\\"TestKafkaEBtrigger\\\"\\\\n}\\u0027\"},\"id\":\"1cb591f9-987e-41d9-b974-0342e9acb90a\",\"source\ ":\"acs:alikafka\",\"specversion\":\"1.0\",\"type\":\"alikafka:Topic:Message\",\"datacontenttype\":\"application/json; charset\\u003dutf-8\",\"time\":\"2022-08-01T09:32:56.797Z\",\"subject\":\"acs:alikafka:alikafka_pre-cn-7pp2t2jwj001:topic: HelloTopic\",\"aliyunaccountid\":\"1938858730552836\"}"]'
 ````

 The response received when calling the function looks like this:

 ```bash
========= FC invoke Logs begin =========
FC Invoke Start RequestId: 8a7a6cce-f123-4204-84a8-fa3f0f3b95bf
[Name] Register [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@58372a00] as [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@com.aliyun.serverless.runtime.classloader.FunctionClassLoader@/code/FCJavaKafkaTrigger-1.0-SNAPSHOT.jar/code/original-FCJavaKafkaTrigger-1.0-SNAPSHOT.jar]: hash [d4d9f0d4] (normal mode)
2022-08-02 02:06:35.360 [INFO] [8a7a6cce-f123-4204-84a8-fa3f0f3b95bf] Event: ["{\"data\":{\"topic\":\"HelloTopic\",\"partition\":9,\"offset\":3,\"timestamp\":1659346376797,\"headers\":{\"headers\":[],\"isReadOnly\":false},\"value\":\"b\\u0027{\\\\n    \\\"Test\\\": \\\"TestKafkaEBtrigger\\\"\\\\n}\\u0027\"},\"id\":\"1cb591f9-987e-41d9-b974-0342e9acb90a\",\"source\":\"acs:alikafka\",\"specversion\":\"1.0\",\"type\":\"alikafka:Topic:Message\",\"datacontenttype\":\"application/json; charset\\u003dutf-8\",\"time\":\"2022-08-01T09:32:56.797Z\",\"subject\":\"acs:alikafka:alikafka_pre-cn-7pp2t2jwj001:topic:HelloTopic\",\"aliyunaccountid\":\"1938858730552836\"}"]
2022-08-02 02:06:35.521 [INFO] [8a7a6cce-f123-4204-84a8-fa3f0f3b95bf] Kafka Topic: HelloTopic
2022-08-02 02:06:35.521 [INFO] [8a7a6cce-f123-4204-84a8-fa3f0f3b95bf] Message Value: b'{\n    "Test": "TestKafkaEBtrigger"\n}'
FC Invoke End RequestId: 8a7a6cce-f123-4204-84a8-fa3f0f3b95bf

Duration: 269.65 ms, Billed Duration: 270 ms, Memory Size: 128 MB, Max Memory Used: 54.12 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e886aa-77e6f1dbaa2c43c0960b

FC Invoke Result:
Produce ok, Topic: HelloTopic Value: b'{\n    "Test": "TestKafkaEBtrigger"\n}'


End of method: invoke
 ```

