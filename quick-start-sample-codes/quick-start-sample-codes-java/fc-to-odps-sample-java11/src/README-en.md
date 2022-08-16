# FCToODPSSampleJava11 Documentation

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

Deploy a python3 function that can insert data into an ODPS (MaxCompute) table on Alibaba Cloud Function Compute.

</description>


## Preliminary Preparation

### Permission Preparation

Using this item, verify that your operational account has the following product permissions/policies:


| Service              | Function Compute                                                          |     
|----------------------|---------------------------------------------------------------------------|   
| permissions/policies | AliyunFCFullAccess</br>AliyunDataWorksFullAccess</br>AliyunBSSOrderAccess |     


### Resource Preparation

  * 1: Activate ODPS (MaxCompute) service.
  * 2: Create ODPS project resources.**（If you have already created or have project resources that can be used directly, you can skip this step.）**
    * 2-1: Enter the ODPS console home page and create project.
    * 2-2: Fill in the workspace name, other options can be skipped if there are no special requirements, click **Create Project** after completion.
    * 2-3: Fill in the instance display name. Other options can be skipped if there are no special requirements. At this point, the process of creating the odps project is completed.
  * 3: Create ODPS Table Resource. **（If you have already created or have project resources that can be used directly, you can skip this step.）**
    * 3-1: Go to the home page of the ODPS console, select the project created in step 2, and click **Data Development**.
    * 3-2: Right-click business process in the figure below, click **New Business Process**, and fill in as required.
    * 3-3: Click **New**->**New Table**->**Table**, configure the table name and business path, and click Finish to enter the table configuration page.
    * 3-4: Set the Chinese name and table fields, and submit it to the production environment after completion. (Note: Table fields are ordered.)
    * 3-5: Click **Table Management** on the left, refresh the page to see the data table just created, and the process of creating the odps data table is complete.

      

## Steps For Usage
### Application Deployment

<appcenter>

- :fire: Use [Serverless Applications](https://fcnext.console.aliyun.com/applications/create?template=FCToODPSSampleJava11) ，
[![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=FCToODPSSampleJava11) deploy the application.

</appcenter>

- Use [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) deploy：
    - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install) ，and make [authorization information configuration](https://www.serverless-devs.com/fc/config) ；
    - Initialize the project：`s init FCToODPSSampleJava11 -d FCToODPSSampleJava11`   
    - Enter the project and deploy the project：`cd FCToODPSSampleJava11 && s deploy -y`

</deploy>

Regardless of the deployment method above, you need to configure the parameters for the application, as follows:
  * accessKeyID/accessKeySecret: The account's ak key pair, which is used to access ODPS (MaxCompute) in the function.
  * odpsProject: The ODPS (MaxCompute) project name needs to start with a letter and can only contain letters, underscores and numbers.
  * odpsEndpoint: ODPS (MaxCompute) service address, if there is no special requirement, you can choose the external network access address, see the address for details：https://help.aliyun.com/document_detail/34951.html。
  * odpsTableName: ODPS(MaxCompute) table name。

### Call Application
**Call Parameters**

After the application is deployed, you can construct a function request parameter to call the function to test the correctness of the function. For examples, please refer to:
```
[
  [1, "xiaoming", 11],
  [2, "xiaoli", 12]
]
```

**Call Method**
  * Console Call：
    * Login [Function Compute Console](https://fcnext.console.aliyun.com/cn-hangzhou/services) ，Find deployed functions.
    * Click **Test function**->**Configure test parameters**, and paste the calling parameters above.
    * Click the test function after completion, and this piece of data will be inserted into the ODPS corresponding table through the function.
  
  * s Call：
    * Enter the application project project and execute the following command：`s invoke --event-file event-example/fc-to-odps-sample.json`.
    * After the function call is completed, this piece of data is inserted into the ODPS corresponding table through the function.

<appdetail id="flushContent">

### Test Verification
After the function is called successfully, how to check whether the data is written successfully on the ODPS side?
  * 1: Enter the ODPS data development page, click **New**->**New Node**->**ODPS SQL**, after the creation is complete, an interface where you can write sql will pop up, and go to the next step.
  * 2：Write sql statement `SELECT * FROM {tableName};`，Click to run sql, and the data in the table will be returned after completion. You can check whether the data in the table contains the event data entered by the test.
  

<appdetail id="flushContent">

# Application Details

This application is only used for learning and reference. You can carry out secondary development and improvement based on this project to realize your own business logic.

</appdetail>

<devgroup>

## Developer Community

If you have feedback about bugs or future expectations, you can give feedback and exchange in [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) .If you want to join our discussion group or keep up to date with the latest developments in FC components, you can do so through the following channels:

<p align="center">

| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
|-----------------------------------------------------------------------------------------------------------------------------------| --- |-----------------------------------------------------------------------------------------------------------------------------------|
| <center>WeChat Public Account：`serverless`</center>                                                                               | <center>WeChat Assistant：`xiaojiangwh`</center> | <center>Dingding Group：`33947367`</center>                                                                                        | 

</p>

</devgroup>