# java11 Spring Boot MongoDB示例

本示例为您展示了 Custom Runtime 的 Spring Boot 使用示例。本示例将使用 Spring Boot 实现 MongoDB 数据查询的接口。
在本示例中，MongoDB 配置保存在文件 ./src/main/java/resources/application.properties 文件中。s.yaml 配置详见 https://gitee.com/devsapp/fc/tree/main/docs/zh/yaml 

本示例 Driver 使用4.6版本。版本兼容详情见 https://www.mongodb.com/docs/drivers/java/sync/current/compatibility

## 准备开始
- 一个可用的 MongoDB 数据库。
使用以下命令可创建本次示例使用的 MongoDB 数据库。
```bash
use users
db.users.insert([
  {"name": "张三", "age": 18},
  {"name": "李四", "age": 20}
])
```


- [可选] 安装并配置 Serverless Devs 工具。（ https://help.aliyun.com/document_detail/195474.html ）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 添加自定义运行时启动配置以及指定端口。
- 由于 Java11 并非 Custom Runtime 的内置编程语言，本示例将 Java11 环境 (路径：./java11，环境需要用户自行下载) 和代码文件一起打包部署到函数计算，详见文档：https://help.aliyun.com/document_detail/132044.html 。
同时使用函数计算提供的 层 也可以实现上述要求，详见文档：https://help.aliyun.com/document_detail/181602.html 。

```yaml
        caPort: 8080
        customRuntimeConfig:
          command:
            - java11/bin/java
            - '-jar'
            - ./target/demo-0.0.1-SNAPSHOT.jar
```

#### 2. 安装依赖并部署

编译部署代码包
```shell
s deploy
```
> 注意: `pom.xml` 中有配置 `pre-deploy` 脚本 `mvn package`, 在部署前会调用 `mvn package` 编译打包。

#### 3. 调用测试

```shell
curl https://java-sp-mongodb-fc-example-xxxxxx.cn-beijing.fcapp.run/mongo
```
收到的响应如下所示：

```bash
Document{{_id=62cb9e5a5c21fd08dbf68490, name=张三, age=18.0}}   
```

### 方式二、使用控制台创建
todo...

## 数据库访问限制
当使用MongoDB时，一般都会有访问控制，比如云数据库MongoDB中的白名单设置（ [MongoDB白名单设置说明](https://help.aliyun.com/document_detail/88888.htm) )。

如果仅仅作为测试，可以将白名单配置成 `0.0.0.0/0`。（不要在生产环境使用!)

在生产环境，可以使用以下两种方式访问：

1. VPC方式（**推荐**） <br>
   参考文档：https://help.aliyun.com/document_detail/84514.html
2. 代理方式<br>
   参考文档：https://help.aliyun.com/document_detail/91243.html


