# Java11 timer triggers FC function example

 This example shows you a Java11 timer trigger FC function example.

## ready to start

 - [Optional] Install and configure Serverless Devs tools. (https://help.aliyun.com/document_detail/195474.html)

## quick start

### Method 1. Use the console to create

#### 1. Install dependencies and deploy code packages

```shell
# Install dependencies and compile them into jar packages, the corresponding jar packages will be in the target directory
mvn clean package
````

#### 2. Create function

 After selecting a service (or creating a service), click Create Function as shown

 - Select `Create from scratch with standard Runtime`
 - Fill in the function name
 - Code upload method select `Upload code via jar package` to upload the corresponding jar compressed package
 - Select the runtime environment Java 11
 - Select function trigger method: trigger by event request
 - Configure the trigger to select `timed trigger`, and fill in the time interval such as 1 (minutes)
 - Use default for other settings

> For the detailed function creation process, see the document: [Create a function using the console](https://help.aliyun.com/document_detail/51783.html)

#### 3. Set service role configuration

 On the Edit Service page, select the service role. It is recommended to select the default role AliyunFCDefaultRole for Function Compute.
 You can also customize service roles, add permission policies, or customize permission policies. For details, see the document [Authorization Policies and Examples](https://help.aliyun.com/document_detail/253969.html)



#### 4. Test function

If the log service is enabled, you can see the log information in the `call log`, the time interval is 1 minute.

```bash
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:49FC Invoke Start RequestId: 2dba780c-f1f6-4fd6-a572-d51f91a3a29b
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:492022-07-29 10:25:49.153 [INFO] [2dba780c-f1f6-4fd6-a572-d51f91a3a29b] triggerTime: 2022-07-29T
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:492022-07-29 10:25:49.154 [INFO] [2dba780c-f1f6-4fd6-a572-d51f91a3a29b] triggerName: timer-java11
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:492022-07-29 10:25:49.154 [INFO] [2dba780c-f1f6-4fd6-a572-d51f91a3a29b] payload: testPayload
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:25:49FC Invoke End RequestId: 2dba780c-f1f6-4fd6-a572-d51f91a3a29b
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:49FC Invoke Start RequestId: c6e86b7e-a319-4e9b-9796-3204c2b3e971
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:492022-07-29 10:26:49.560 [INFO] [c6e86b7e-a319-4e9b-9796-3204c2b3e971] triggerTime: 2022-07-29
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:492022-07-29 10:26:49.561 [INFO] [c6e86b7e-a319-4e9b-9796-3204c2b3e971] triggerName: timer-java11
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:492022-07-29 10:26:49.561 [INFO] [c6e86b7e-a319-4e9b-9796-3204c2b3e971] payload: testPayload
c-62e3b541-92ebfa07e7a04c658bb82022-07-29 18:26:49FC Invoke End RequestId: c6e86b7e-a319-4e9b-9796-3204c2b3e971
````



### Method 2. Compile and deploy using Serverless Devs tools

#### 1. Modify s.yaml configuration

 - [Optional] Modify the corresponding region as needed

#### 2. Install dependencies and deploy

 Install dependent libraries

```shell
# Use the s tool to install dependencies, you need to use docker
s build --use-docker
````

 deploy code

```bash
# deploy code
s deploy -y
````

#### 3. Invoke the test

```shell
s invoke -e '{
    "triggerTime": "2022-07-29T10:02:58Z",
    "triggerName": "nodejs14-timer",
    "payload": "TestPayload"
}'
````

 The response received when calling the function looks like this:

```bash
========= FC invoke Logs begin =========
FC Invoke Start RequestId: bb1f5703-7e9c-4494-8582-2cad07dbf9f7
[Name] Register [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@58372a00] as [com.aliyun.serverless.runtime.classloader.FunctionClassLoader@com.aliyun.serverless.runtime.classloader.FunctionClassLoader@/code/HelloFCJavaKafka- 1.0-SNAPSHOT.jar/code/original-HelloFCJavaKafka-1.0-SNAPSHOT.jar]: hash [d4d9f0d4] (normal mode)
2022-07-30 08:53:46.510 [INFO] [bb1f5703-7e9c-4494-8582-2cad07dbf9f7] triggerTime: 2022-07-29T10:02:58Z
2022-07-30 08:53:46.510 [INFO] [bb1f5703-7e9c-4494-8582-2cad07dbf9f7] triggerName: nodejs14-timer
2022-07-30 08:53:46.510 [INFO] [bb1f5703-7e9c-4494-8582-2cad07dbf9f7] payload: TestPayload
FC Invoke End RequestId: bb1f5703-7e9c-4494-8582-2cad07dbf9f7

Duration: 233.88 ms, Billed Duration: 234 ms, Memory Size: 128 MB, Max Memory Used: 81.87 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4f199-8b9aa92268c64154950d

FC Invoke Result:
Timer Payload: TestPayload


End of method: invoke
````
