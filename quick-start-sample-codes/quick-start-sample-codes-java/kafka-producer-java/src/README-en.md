# Java11 Kafka message queue producer example

This example shows you the Java11 [Message Queue Kafka](https://help.aliyun.com/document_detail/68151.html?spm=5176.167616.J_5253785160.5.2dfe6feexRPqMj) producer example. (See Kafka Trigger Example for Consumer Example)

This example uses the [Message Queue Kafka Version Official Java SDK](https://help.aliyun.com/document_detail/68325.html).

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

When creating a service, select AliyunFcDefaultRole in `Service Role` in `Advanced Options` (if not, create a corresponding role according to the prompts), and enable `Allow access to VPC`, and select `VPC' and `Switch when creating a Kafka instance. ` and the corresponding `security group (created automatically after the Kafka instance is deployed)`.

![CreateService.png](assets/CreateService.png)



#### 3. Create function

  After creating the service, click Create Function as shown

 - Select `Create from scratch with standard Runtime`
 - Fill in the function name
 - Code upload method select `Upload code via jar package` to upload the corresponding jar compressed package
 - Select the runtime environment Java 11
 - Select function trigger method: trigger by event request
 - Use default for other settings

 > For the detailed function creation process, see the document: [Create a function using the console](https://help.aliyun.com/document_detail/51783.html)



#### 4. Configure environment variables and instance lifecycle callbacks

Set environment variables in the `Function Configuration` module in the function details and configure the Initializer callback procedure in the instance lifecycle callback.

Where environment variables:

- BOOTSTRAP_SERVERS is set to the `default access point` address corresponding to the `access point information` in the Kafka instance details.
- TOPIC_NAME is set to the topic to which the corresponding message is sent (need to be created in advance in the Kafka message queue version)

Initializer is set to example.App::initialize

![FunctionConfig.png](assets/FunctionConfig.png)



 #### 5. Test function

 The return result is as follows

 ```bash
Produce ok: HelloTopic-6@20
 Payload: {
    "key1": "value1",
    "key2": "value2",
    "key3": "value3"
}
 ````



 ### Method 2. Compile and deploy using Serverless Devs tools

 #### 1. Modify s.yaml configuration

- Modify region, serviceName, functionName (set the same region as the Kafka instance)

- Modify vpcConfig and fill in the VPC ID, security group ID, and vSwitchID corresponding to the Kafka instance.

- Modify the environmentVariables configuration, fill in BOOTSTRAP_SERVERS and TOPIC_NAME


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

 ```shell
s invoke -e '{
    "Key": "test java serverless devs"
}'
 ````

 The response received when calling the function looks like this:

 ```bash
========= FC invoke Logs begin =========
FC Initialize Start RequestId: 951d096c-d9d9-4105-806b-d350f31786c9
[Name] Register [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@58372a00] as [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@com.aliyun.serverless.runtime.classloader.FunctionClassLoader@/code/HelloFCJavaKafka- 1.0-SNAPSHOT.jar/code/original-HelloFCJavaKafka-1.0-SNAPSHOT.jar]: hash [d4d9f0d4] (normal mode)
SLF4J: Failed to load class "org.slf4j.impl.StaticLoggerBinder".
SLF4J: Defaulting to no-operation (NOP) logger implementation
SLF4J: See http://www.slf4j.org/codes.html#StaticLoggerBinder for further details.
FC Initialize End RequestId: 951d096c-d9d9-4105-806b-d350f31786c9
FC Invoke Start RequestId: 951d096c-d9d9-4105-806b-d350f31786c9
2022-08-08 09:41:44.123 [INFO] [951d096c-d9d9-4105-806b-d350f31786c9] Produce ok: HelloTopic-10@18
 Payload: {
"Key": "test java serverless devs"
}
FC Invoke End RequestId: 951d096c-d9d9-4105-806b-d350f31786c9
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e61808-256d56a4065243eb9951

FC Invoke Result:
Produce ok: HelloTopic-10@18
 Payload: {
    "Key": "test java serverless devs"
}


End of method: invoke
 ````
