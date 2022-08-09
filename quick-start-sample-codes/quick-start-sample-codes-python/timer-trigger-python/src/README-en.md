# Python3 timer triggers FC function example

 This example shows you a Python3 timer trigger FC function example.

 ## ready to start

 - [Optional] Install and configure Serverless Devs tools. (https://help.aliyun.com/document_detail/195474.html)

 ## quick start

 ### Method 1. Use the console to create

 #### 1. Install dependencies and deploy code packages

 ```shell
# Install dependencies to /code directory
cd code && pip3 install -r requirements.txt -t .
# package file
cd code && zip -r python3-timer.zip *
 ````

 #### 2. Create function

 After selecting a service (or creating a service), click Create Function as shown

 - Select `Create from scratch with standard Runtime`
 - Fill in the function name
 - Select the code upload method `Upload code via zip package` to upload the corresponding code zip package
 - Select the runtime environment Python 3.6
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
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:37FC Invoke Start RequestId: bd03dbc7-ce59-488d-be55-b579670afe4c
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:372022-07-29 18:17:37 bd03dbc7-ce59-488d-be55-b579670afe4c [INFO] event: b'{"triggerTime":"2022-07- :17:37Z","triggerName":"timer-python3","payload":"testPayload"}'
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:372022-07-29 18:17:37 bd03dbc7-ce59-488d-be55-b579670afe4c [INFO] triggerName: timer-python3
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:372022-07-29 18:17:37 bd03dbc7-ce59-488d-be55-b579670afe4c [INFO] triggerTime = 2022-07-29T10:17:37Z
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:17:37FC Invoke End RequestId: bd03dbc7-ce59-488d-be55-b579670afe4c
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:37FC Invoke Start RequestId: ef169a63-4342-4019-b129-f528fd9463f8
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:372022-07-29 18:18:37 ef169a63-4342-4019-b129-f528fd9463f8 [INFO] event: b'{"triggerTime":"2022 :18:37Z","triggerName":"timer-python3","payload":"testPayload"}'
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:372022-07-29 18:18:37 ef169a63-4342-4019-b129-f528fd9463f8 [INFO] triggerName: timer-python3
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:372022-07-29 18:18:37 ef169a63-4342-4019-b129-f528fd9463f8 [INFO] triggerTime = 2022-07-29T10
c-62e3b3b2-61477a6e557f4510a4d22022-07-29 18:18:37FC Invoke End RequestId: ef169a63-4342-4019-b129-f528fd9463f8
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
FunctionCompute python3 runtime initiated.
FC Invoke Start RequestId: a785f24e-b67a-400d-b3cc-4472a8587400
2022-07-30T08:49:52.165Z a785f24e-b67a-400d-b3cc-4472a8587400 [INFO] event: b'{\n "triggerTime": "2022-07-29T10:02:58Z",\n "triggerName" : "nodejs14-timer",\n "payload": "TestPayload"\n}'
2022-07-30T08:49:52.166Z a785f24e-b67a-400d-b3cc-4472a8587400 [INFO] triggerName: nodejs14-timer
2022-07-30T08:49:52.166Z a785f24e-b67a-400d-b3cc-4472a8587400 [INFO] triggerTime = 2022-07-29T10:02:58Z
FC Invoke End RequestId: a785f24e-b67a-400d-b3cc-4472a8587400

Duration: 3.90 ms, Billed Duration: 4 ms, Memory Size: 128 MB, Max Memory Used: 24.86 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4f0af-aa8569cde2ed43aaaf04

FC Invoke Result:
Timer Payload: TestPayload


End of method: invoke
 ````
