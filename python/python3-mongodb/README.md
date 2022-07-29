# Python3 MongoDB示例
本示例为您展示了Python runtime的 MongoDB 使用示例。
在本示例中，MongoDB 数据库配置在函数的环境变量配置中（参考s.yaml，yaml 配置详见 https://gitee.com/devsapp/fc/tree/main/docs/zh/yaml)，initialize 回调函数从环境变量中获取数据库配置，创建 MongoDB 连接，这样可以实现链接在整个函数实例生命周期内的复用。preStop 回调函数负责关闭 MongoDB 连接。回调函数与函数实例生命周期的关系详见 https://help.aliyun.com/document_detail/203027.html

本示例 Driver 使用4.1.1版本。版本兼容详情见 https://www.mongodb.com/docs/drivers/pymongo/#compatibility

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
        initializer: index.initialize
        instanceLifecycleConfig:
          preStop:
            handler: index.pre_stop
            timeout: 20
```

#### 2. 安装依赖并部署

安装依赖库 pymongo
```shell
# 使用s工具安装依赖，需要使用 docker
s build --use-docker
```
> 注意: 使用 `s build --use-docker` 会将依赖安装到代码包的 `.s/python` 目录下，可以在本地的 `.s/build/artifacts/{serviceName}/{functionName}/` 查看。

部署代码
```bash
# 部署代码，会自动添加环境变量 PYTHONUSERBASE=/code/.s/python，该环境变量是必须的
s deploy
```

#### 3. 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
s invoke
========= FC invoke Logs begin =========
FunctionCompute python3 runtime inited.
FC Initialize Start RequestId: b2******
FC Initialize End RequestId: b2******
FC Invoke Start RequestId: b2******
FC Invoke End RequestId: b2******

Duration: 176.03 ms, Billed Duration: 177 ms, Memory Size: 128 MB, Max Memory Used: 62.08 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-6******

FC Invoke Result:
{'_id': ObjectId('62cb9e5a5c21fd08dbf68490'), 'name': '张三', 'age': 18.0}


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

