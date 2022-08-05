# rds-mysql-fc-event-python3 help documentation

<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-event-python3&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rds-mysql-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-event-python3&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=rds-mysql-fc-event-python3" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=rds-mysql-fc-event-python3&type=packageDownload">
  </a>
</p>

<description>

Quickly deploy a Python 3 Event type read and write Mysql database function to Alibaba Cloud Function Compute. In this case, the public network is provided to connect to the RDS MySQL database.

</description>

## Pre-preparation
To use this program, it is recommended that you have the following product permissions / strategies.

| Services / Business  | functional computing |     
| --- |  --- |   
| Permissions/Policies | AliyunFCFullAccess |

To use this project, you need to prepare the MySQL database in advance and execute the following SQL statement to create the table:
  ```sql
    CREATE TABLE `users` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
      `name` varchar(20) NOT NULL,
      `age` tinyint(11) NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
  ```
For MySQL database, you can choose either Alibaba Cloud MySQL database or other MySQL database. In this case, we use the public network to connect to the Alibaba Cloud MySQL database. If you use the vpc method to connect to the Alibaba Cloud MySQL database, please refer to the document [Configure Network](https://help.aliyun.com/document_detail/72959.html) Configure the VPC network.
- Use Alibaba Cloud Database RDS MySQL instance, the creation tutorial can refer to [here](https://help.aliyun.com/document_detail/26117.htm?spm=a2c4g.11186623.0.0.12a47634PzmWPx)
  - Log in to the RDS console as a MySQL instance [Apply for an external network address](https://help.aliyun.com/document_detail/26128.html), which is convenient for accessing the database on the public network
  - [Set IP whitelist](https://help.aliyun.com/document_detail/96118.html), in this case as a test, you can configure the whitelist to 0.0.0.0/0. (Don't use it in production!)
  - You need to provide the correct database URL address, database name, user, and password to connect to the database
- Use other MySQL databases
  - You need to provide the correct database URL address, database name, user, and password to connect to the database

<codepre id="codepre">

# Code & Preview

- [ :smiley_cat:  source code](https://github.com/devsapp/start-fc/blob/main/event-function/rds-mysql-fc-event-python3)
- In order to successfully deploy this sample code, you need to provide the following parameters during the deployment process:
    - Region: You need to configure the region where your Function Compute service needs to be deployed through this parameter. The default value is cn-hangzhou (Hangzhou).
      - The geographic options available to you are:
        - cn-beijing (beijing)
        - cn-hangzhou (hangzhou)
        - cn-shanghai (shanghai)
        - cn-qingdao (qingdao)
        - cn-zhangjiakou (zhangjiakou)
        - cn-huhehaote (huhehaote)
        - cn-shenzhen (shenzhen)
        - cn-chengdu (chengdu)
        - cn-hongkong (Hongkong)
        - ap-southeast-1 (Singapore)
        - ap-southeast-2 (sydney)
        - ap-southeast-3 (Kuala Lumpur)
        - ap-southeast-5 (Jakarta)
        - ap-northeast-1 (Tokyo)
        - eu-central-1 (Frankfurt)
        - eu-west-1 (London)
        - us-west-1 (Silicon Valley)
        - us-east-1 (Virginia)
        - ap-south-1 (Mumbai)
  - Service name: You need to name your Function Compute service. The service name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-128, the default value is rds-mysql-quick-start.
  - Function name: You need to name your function calculation function. The function name can only contain letters, numbers, underscores and dashes. Cannot start with a number or a dash. The length is between 1-64. The default is rds-mysql-event-function-python.
  - MySQL address (mysql url): MySQL address, which can be viewed by logging in to the RDS console -> Instance List -> Basic Information -> Network Type -> External Network Address to connect to the database.
  - MySQL port number (mysql port): MySQL port number, the default is 3306.
  - MySQL database name (mysql database name): MySQL database name, which can be viewed by logging in to the RDS console -> Instance List -> Database Management -> Database Name to connect to the database.
  - MySQL user name (mysql user): MySQL user, you can log in to the RDS console -> Instance List -> Account Management -> User Account to view it to connect to the database.
  - MySQL password (mysql password): MySQL password, which can be viewed by logging in to the RDS console -> Instance List -> Account Management -> User Account to connect to the database.

</codepre>

<deploy>

## Deployment & Experience

<appcenter>

-  :fire: By [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=rds-mysql-fc-event-python3),one-click deployment

</appcenter>

- Deployment via  [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) ：
    - [Install Serverless Devs Cli Developer Tools](https://www.serverless-devs.com/serverless-devs/install) ，And make [authorization information configuration](https://www.serverless-devs.com/fc/config) ；
    - Initialize the project：`s init rds-mysql-fc-event-python3 -d rds-mysql-fc-event-python3` 
    - Fill in the parameters described in the above modules
    - Enter the project and deploy the project：`cd rds-mysql-fc-event-python3 && s deploy -y`
## Application details  
- local invoke
  - Run `s invoke` for local debugging
  - The response received when calling the function looks like this:
    ```bash
    ========= FC invoke Logs begin =========
    FunctionCompute python3 runtime inited.
    FC Initialize Start RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    FC Initialize End RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    FC Invoke Start RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    2022-03-31T02:57:49.693Z 28fa11ab-81da-4cd0-b050-xxxxxxxxxx [INFO] (3, 'wangEr', 38)
    FC Invoke End RequestId: 28fa11ab-81da-4cd0-b050-xxxxxxxxxx
    Duration: 18.42 ms, Billed Duration: 19 ms, Memory Size: 128 MB, Max Memory Used: 34.80 MB
    ========= FC invoke Logs end =========
    FC Invoke Result:
    user: (3, 'wangEr', 38)
    End of method: invoke
      ```
- end-to-end testing
  - Log in to the FC console and test the function
  - The console returns the result as follows:
    ```bash
    {name=wangEr, id=3, age=38}
    ```
- Database access restrictions
  - When using cloud database, there is usually access control, and you need to [set IP whitelist](https://help.aliyun.com/document_detail/96118.html),In this case as a test, the whitelist can be configured to 0.0.0.0/0. (Don't use it in production!).
  - In a production environment, it can be accessed in the following two ways:
    - VPC method（**recommend**） 
    Reference document: [Configure Network](https://help.aliyun.com/document_detail/72959.html)
    - 公网方式
    Reference document: [Configure a fixed public IP address](https://help.aliyun.com/document_detail/410740.html)
</deploy>

<appdetail id="flushContent">

# common problem
- Whitelist not set, MySQL URL, port input error
    ```bash
     "errorMessage": "(2003, \"Can't connect to MySQL server on 'rm-uf67i8axxxxxxxxxx.mysql.rds.aliyuncs.com' (timed out)\")",
    ```
- MySQL username and password are incorrect
    ```bash
     "errorMessage": "(1045, \"Access denied for user 'fc1'@'120.xx.xx.xx' (using password: YES)\")"
    ```
- MySQL database name is wrong
    ```bash
     "errorMessage": "(1049, \"Unknown database 'users1'\")"
    ```     
</appdetail>

<devgroup>

## developer community

If you have feedback about bugs or future expectations, you can give feedback and exchange in [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) . If you want to join our discussion group or keep up to date with the latest developments in FC components, you can do so through the following channels:

<p align="center">

| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
|--- | --- | --- |
| <center>WeChat public account：`serverless`</center> | <center>WeChat Assistant：`xiaojiangwh`</center> | <center>Dingding exchange group：`33947367`</center> | 

</p>

</devgroup>