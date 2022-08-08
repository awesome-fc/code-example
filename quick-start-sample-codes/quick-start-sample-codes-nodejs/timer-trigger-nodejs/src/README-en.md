# Nodejs14 timer trigger FC function example

 This example shows you Nodejs14's timer trigger FC function example.

 ## ready to start

 - [Optional] Install and configure Serverless Devs tools. (https://help.aliyun.com/document_detail/195474.html)

 ## quick start

 ### Method 1. Use the console to create

 #### 1. Install dependencies and deploy code packages

 ```shell
# Install dependencies in the code directory and compile
npm install
# compress
zip code.zip -r ./*
 ````

 #### 2. Create function

 After selecting a service (or creating a service), click Create Function as shown

 - Select `Create from scratch with standard Runtime`
 - Fill in the function name
 - Select the code upload method `Upload code via zip package` to upload the corresponding code zip package
 - Select the runtime environment Nodejs14
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
62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:58FC Invoke Start RequestId: 209d5189-18c4-4b9e-9c3f-12d255cc14d4
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] whole event: {"triggerTime":"2022- 02:58Z","triggerName":"nodejs14-timer","payload":"TestPayload"}
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] triggerName: nodejs14-timer
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] triggerTime: 2022-07-29T10
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:582022-07-29 18:02:58 209d5189-18c4-4b9e-9c3f-12d255cc14d4 [verbose] triggerMessgae: TestPayload
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:02:58FC Invoke End RequestId: 209d5189-18c4-4b9e-9c3f-12d255cc14d4
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:58FC Invoke Start RequestId: d5d22563-78ba-4198-b819-4d30154bdace
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] whole event: {"triggerTime":"2022-07 03:58Z","triggerName":"nodejs14-timer","payload":"TestPayload"}
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] triggerName: nodejs14-timer
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] triggerTime: 2022-07-29T10:0
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:582022-07-29 18:03:58 d5d22563-78ba-4198-b819-4d30154bdace [verbose] triggerMessgae: TestPayload
c-62e3b016-c9d04dfa1562484cbd8b2022-07-29 18:03:58FC Invoke End RequestId: d5d22563-78ba-4198-b819-4d30154bdace
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
FC Invoke Start RequestId: e8ff75e9-08e6-4fd9-be64-d91894700d86
load code for handler:index.handler
} "payload": "TestPayload"-timer",:58Z",9-be64-d91894700d86 [verbose] whole event: {
2022-07-30T08:44:59.548Z e8ff75e9-08e6-4fd9-be64-d91894700d86 [verbose] triggerName: nodejs14-timer
2022-07-30T08:44:59.548Z e8ff75e9-08e6-4fd9-be64-d91894700d86 [verbose] triggerTime: 2022-07-29T10:02:58Z
2022-07-30T08:44:59.548Z e8ff75e9-08e6-4fd9-be64-d91894700d86 [verbose] triggerMessgae: TestPayload
FC Invoke End RequestId: e8ff75e9-08e6-4fd9-be64-d91894700d86

Duration: 3.92 ms, Billed Duration: 4 ms, Memory Size: 128 MB, Max Memory Used: 48.84 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4ef8b-5959048c584a48f8bd95

FC Invoke Result:
timer trigger:TestPayload


End of method: invoke
 ````
