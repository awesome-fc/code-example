# Golang timer triggers FC function example

 This example shows you an example of Golang's timer-triggered FC function.

 ## ready to start

 - [Optional] Install and configure Serverless Devs tools. (https://help.aliyun.com/document_detail/195474.html)

 ## quick start

 ### Method 1. Use the console to create

 #### 1. Install dependencies and deploy code packages

> For more details on deploying the Go code package, see (https://help.aliyun.com/document_detail/418490.html)

 ```shell
# Install dependencies in the code directory and compile
go mod tidy
GOOS=linux GOARCH=amd64 go build main.go
# Compress the executable
zip fc-golang-demo.zip main
 ````

 #### 2. Create function

 After selecting a service (or creating a service), click Create Function as shown

 - Select `Create from scratch with standard Runtime`
 - Fill in the function name
 - Select the code upload method `Upload code via zip package` to upload the corresponding code zip package
 - Select the runtime environment Go 1
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
c-62e46764-132ecab75fa946deb5462022-07-30 11:43:58FC Invoke Start RequestId: 37421a15-6150-462b-bf0b-685e48b66128
c-62e46764-132ecab75fa946deb5462022-07-30 11:43:582022-07-30 11:43:58 37421a15-6150-462b-bf0b-685e48b66128 [INFO] main.go:20: triggerName: timer-go1
c-62e46764-132ecab75fa946deb5462022-07-30 11:43:582022-07-30 11:43:58 37421a15-6150-462b-bf0b-685e48b66128 [INFO] main.go:21:triggerTime: 2022-07-30T :58Z
c-62e46764-132ecab75fa946deb5462022-07-30 11:43:582022-07-30 11:43:58 37421a15-6150-462b-bf0b-685e48b66128 [INFO] main.go:22: payload: testPayload
c-62e46764-132ecab75fa946deb5462022-07-30 11:43:58FC Invoke End RequestId: 37421a15-6150-462b-bf0b-685e48b66128
c-62e46764-132ecab75fa946deb5462022-07-30 11:44:58FC Invoke Start RequestId: f2145f49-6a99-4a64-bd88-6c49c6a37e25
c-62e46764-132ecab75fa946deb5462022-07-30 11:44:582022-07-30 11:44:58 f2145f49-6a99-4a64-bd88-6c49c6a37e25 [INFO] main.go:20: triggerName: timer-go1
c-62e46764-132ecab75fa946deb5462022-07-30 11:44:582022-07-30 11:44:58 f2145f49-6a99-4a64-bd88-6c49c6a37e25 [INFO] main.go:21:triggerTime: 2022-07-30T0 :58Z
c-62e46764-132ecab75fa946deb5462022-07-30 11:44:582022-07-30 11:44:58 f2145f49-6a99-4a64-bd88-6c49c6a37e25 [INFO] main.go:22: payload: testPayload
 ````



 ### Method 2: Compile and deploy using Serverless Devs tools

 #### 1. Modify s.yaml configuration

 - [Optional] Modify the corresponding region as needed

 #### 2. Install dependencies and deploy

 Install dependent libraries

 ```shell
# Use the s tool to install dependencies, you need to use docker
s build --use-docker
 ````

 > Note: Using `s build --use-docker` will install dependencies into the `.s/python` directory of the code package, which can be found locally in `.s/build/artifacts/{serviceName}/{functionName}/ ` View.

 deploy code

 ```bash
# Deploy the code, the environment variable will be automatically added PYTHONUSERBASE=/code/.s/python, this environment variable is required
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
FC Invoke Start RequestId: 953240e4-a6e3-452b-a2ed-aef6518dc3af
2022-07-30T08:17:49.606Z 953240e4-a6e3-452b-a2ed-aef6518dc3af [INFO] main.go:20: triggerName: nodejs14-timer
2022-07-30T08:17:49.606Z 953240e4-a6e3-452b-a2ed-aef6518dc3af [INFO] main.go:21: triggerTime: 2022-07-29T10:02:58Z
2022-07-30T08:17:49.606Z 953240e4-a6e3-452b-a2ed-aef6518dc3af [INFO] main.go:22: payload: TestPayload
FC Invoke End RequestId: 953240e4-a6e3-452b-a2ed-aef6518dc3af

Duration: 1.13 ms, Billed Duration: 2 ms, Memory Size: 128 MB, Max Memory Used: 9.69 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62e4e919-8073a286587848f3b97c

FC Invoke Result:
"Timer Payload: TestPayload"


End of method: invoke
 ````
