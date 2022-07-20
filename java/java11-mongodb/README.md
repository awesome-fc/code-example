# java11 MongoDB示例

本示例为您展示了 Java runtime 的 MongoDB 使用示例。 
在本示例中，MongoDB 数据库配置在函数的环境变量配置中（参考s.yaml，yaml 配置详见 https://gitee.com/devsapp/fc/tree/main/docs/zh/yaml)，initialize 回调函数从环境变量中获取数据库配置，创建 MongoDB 连接，preStop 回调函数负责关闭 MongoDB 连接。回调函数与函数实例生命周期的关系详见 https://help.aliyun.com/document_detail/203027.html

本示例 Driver 使用4.6版本。版本兼容详情见 https://www.mongodb.com/docs/drivers/java/sync/current/compatibility


## 准备开始
- 一个可用的 MongoDB 数据库，可以参考以下命令创建测试数据库。

```bash
use users
db.users.insert([
  {"name": "张三", "age": 18},
  {"name": "李四", "age": 20}
])
```


- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 修改 environmentVariables 配置，填入 MONGO_URL, MONGO_DATABASE
- 使用 initialize 和 prestop 回调，需要在 s.yaml 中配置

```yaml
        initializationTimeout: 60
        initializer: example.App::initialize
        instanceLifecycleConfig:
          preStop:
            handler: example.App::preStop
            timeout: 60
```

#### 2. 安装依赖并部署

编译部署代码包
```shell
s deploy
```
> 注意: `pom.xml` 中有配置 `pre-deploy` 脚本 `mvn package`, 在部署前会调用 `mvn package` 编译打包。

#### 3. 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
FC Invoke Start RequestId: 71946c57-******
2022-07-12 03:13:22.763 [INFO] [71946c57-10f7-4548-8574-866cfa29c591] get user: Document{{_id=62cb9e5a5c21fd08dbf68490, name=张三, age=18.0}}
FC Invoke End RequestId: 71946c57-******

Duration: 344.21 ms, Billed Duration: 345 ms, Memory Size: 128 MB, Max Memory Used: 102.32 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-62cce588-******

FC Invoke Result:
Document{{_id=62cb9e5a5c21fd08dbf68490, name=张三, age=18.0}}


End of method: invoke
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

