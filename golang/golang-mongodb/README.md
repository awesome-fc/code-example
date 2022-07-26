# Golang Runtime 的 MongoDB 示例
本示例为您展示了 Golang Runtime 的 Mongodb 使用示例。
在本示例中，MongoDB 数据库配置在函数的环境变量配置中（参考s.yaml，yaml 配置详见 https://gitee.com/devsapp/fc/tree/main/docs/zh/yaml)，initializer 回调函数从环境变量中获取数据库配置，创建 MongoDB 连接，这样可以实现链接在整个函数实例生命周期内的复用。preStop 回调函数负责关闭 MongoDB 连接。回调函数与函数实例生命周期的关系详见 https://help.aliyun.com/document_detail/203027.html

本示例 Driver 使用1.10版本。版本兼容详情见https://www.mongodb.com/docs/drivers/go/current/compatibility

## 准备开始
- 一个可用的 MongoDB 数据库，可以参考以下命令创建测试数据库。

```bash
use users
db.users.insert([
  {"name": "张三", "age": 18},
  {"name": "李四", "age": 20}
])
```

- [可选] 安装并配置 Serverless Devs 工具，以及 docker daemon。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 修改 environmentVariables 配置，填入 MONGO_URL, MONGO_DATABASE
- 使用 initialize 和 prestop 回调，需要在 s.yaml 中配置

```yaml
        initializationTimeout: 20
        initializer: "true"
        instanceLifecycleConfig:
          preStop:
            handler: "true"
            timeout: 30
```

#### 2. 安装依赖并部署

部署代码
```bash
# 部署代码
s deploy
```

#### 3. 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
2022/07/26 08:18:01.996744 start
FC Initialize Start RequestId: 9f******
FC Initialize End RequestId: 9f******
FC Invoke Start RequestId: 80******
FC Invoke End RequestId: 80******

Duration: 158.00 ms, Billed Duration: 158 ms, Memory Size: 128 MB, Max Memory Used: 14.18 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-******

FC Invoke Result:
{"_id":"62df91e92cd055af86f892fd","age":18,"name":"张三"}


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
