# FCToLindormGolang Documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToLindormGolang&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=FCToLindormGolang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToLindormGolang&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=FCToLindormGolang" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=FCToLindormGolang&type=packageDownload">
  </a>
</p>

<description>

Quickly deploy a function that can create a table and insert data on the lindorm wide table engine on Alibaba Cloud Function Compute.

</description>


## Preliminary Preparation

### Permission Preparation

Using this item, verify that your operational account has the following product permissions/policies:


| Service              | Function Compute                                                     |     
|----------------------|----------------------------------------------------------------------|   
| permissions/policies | AliyunFCFullAccess</br>AliyunLindormFullAccess</br>AliyunVPCFullAccess |     

### Resource Preparation

* 1：Activate the Lindorm service.
* 2：Create a Lindorm instance.
  Click Create Instance List and fill in the corresponding options as shown in the figure below. (Remember the vpcID and vswitchID filled in here, and you also need to fill in this parameter when deploying the function below.)
  ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/ffd1cc739fade8edbb435f3e13f73e67/image.png)
* 3：Set the access whitelist for Lindorm: You cannot access without the whitelist. Lindorm prohibits setting 0.0.0.0/0 (that is, open to all clients) for security reasons. You can set the IPV4 network segment of the vswitch as the access whitelist .
    * 3-1：Go to the VPC console to view the switch configured in the previous step, and copy the IPV4 network segment in the figure below.
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/25b4fb29d4f5ab82838dc4e1cdec8da5/image.png)
    * 3-2：Click Access Control in the Lindorm console, and paste the copied network segment into the whitelist group.
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/d5ad661e0648050380a66909b9c712d5/image.png)

## Steps For Usage
### Application Deployment

<appcenter>

- :fire: Use [Serverless Applications](https://fcnext.console.aliyun.com/applications/create?template=FCToLindormGolang) ，
  [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=FCToLindormGolang) deploy the application.

</appcenter>

- Use [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) deploy：
    - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install) ，and make [authorization information configuration](https://www.serverless-devs.com/fc/config) ；
    - Initialize the project：`s init FCToLindormGolang -d FCToLindormGolang`
    - Enter the project and deploy the project：`cd FCToLindormGolang && s deploy -y`

</deploy>


Regardless of the deployment method above, you need to configure the parameters for the application, as follows:
* Function Compute VPC configuration:
    * vpcID：The vpc id where the function instance resides must be the same as the vpc where lindorm resides.
    * vswitchID：vswitch id in vpc, used to access lindorm from private network.
    * securityGroupID：The security group id under vpc is used to access lindorm in the private network. (If there is no one to create one, use the default rules for the outbound and inbound directions.)
* lindormUserName/lindormPassword：The account password for accessing the lindorm wide table engine, the default value is root. (It can be viewed in `lindorm console -> database connection -> wide table engine`.)
* databaseURL：Access the private network address of the lindorm server. (It can be viewed in `lindorm console -> database connection -> wide table engine`.)
  Note: Only the second half of `=` can be intercepted, as shown below.
  ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/51ec7ef0ade3a6f2f57dd738f6d73fb9/image.png)
* sqlTableName: The name of the table created by the function in the wide table engine, the function will automatically create this table in the wide table engine and insert data.

### Application Call
**Call Parameters**

No need to fill in any parameters.

**Call Method**
* Console Call：
    * Login [Function Compute Console](https://fcnext.console.aliyun.com/cn-hangzhou/services) ，Find deployed functions.
    * Click **Test function**->**Configure test parameters**，The request can be executed successfully.
    * You can view the info log in the log output in the figure below. If the row data output by the info log is the same as the inserted data, the test is successful.
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/fdf940efa705b779adfff2c100ffe9d9/image.png)


* s Call：
    * Enter the application project and execute the following command：`s invoke --event "{}"`.
    * You can view the info log in the log output in the figure below. If the row data output by the info log is the same as the inserted data, the test is successful。
      ![image](http://git.cn-hangzhou.oss-cdn.aliyun-inc.com/uploads/serverless/serverless-solutions/db7f514a7338d90c54be1afa35c29f0e/image.png)

<appdetail id="flushContent">

## Advanced Features

For the protection of downstream services, the default concurrency limit of function instances in this application is 10. If you need higher concurrency, you can configure to increase the limit yourself.

## Application Details

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