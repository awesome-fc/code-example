# Golang Runtime 的表格存储示例
本示例为您展示了 Golang Runtime 的表格存储使用示例。
在本示例中，表格存储实例配置在函数的环境变量配置中（参考s.yaml)，initializer 回调函数从环境变量中获取配置，创建表格存储客户端，可以实现在整个函数实例生命周期内复用该客户端。

## 准备开始
- 一个可用的表格存储实例，开通实例等教程可以查看阿里云官方文档。(https://help.aliyun.com/product/27278.html)
- 配置服务角色。

   当函数计算需访问阿里云其他云服务时，需要为函数计算授予相应的权限。详见文档(https://help.aliyun.com/document_detail/181589.html)。
   但是函数计算默认的服务角色 AliyunFCDefaultRole 不包含表格存储的权限，因此需要为该角色添加表格存储权限，也可以创建并使用新的角色。
   * 进入[RAM角色管理](https://ram.console.aliyun.com/roles)。
   * 找到 AliyunFCDefaultRole 并点击右侧**添加权限**。
   * 在**添加权限**页面，查找“表格存储”，选择 AliyunOTSFullAccess 权限，单击确定为角色添加表格存储权限。
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）


## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 在 service 处添加角色配置。
```yaml
role: acs:ram::{your account id}:role/aliyunfcdefaultrole
```
- 修改 environmentVariables 配置，填入 ENDPOINT, INSTANCE_NAME
- 使用 initializer 回调，需要在 s.yaml 中配置

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
2022/07/26 09:35:37.215418 start
FC Initialize Start RequestId: 2c******
FC Initialize End RequestId: 2c******
FC Invoke Start RequestId: 2c******
FC Invoke End RequestId: 2c******

Duration: 137.28 ms, Billed Duration: 138 ms, Memory Size: 128 MB, Max Memory Used: 19.82 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-******

FC Invoke Result:
{"PrimaryKey":{"PrimaryKeys":[{"ColumnName":"region","Value":"abc","PrimaryKeyOption":0},{"ColumnName":"id","Value":1,"PrimaryKeyOption":0}]},"Columns":[{"ColumnName":"age","Value":"20","Timestamp":1657531733801},{"ColumnName":"home","Value":"北京","Timestamp":1657618107569},{"ColumnName":"name","Value":"张三","Timestamp":1657531733801}],"ConsumedCapacityUnit":{"Read":1,"Write":0},"RequestId":"0005e4b2-05df-963c-69ef-700b9c940bae"}


End of method: invoke
```

### 方式二、使用控制台创建
todo...


## 表格存储访问限制

在生产环境，可以使用以下两种方式访问：

1. VPC方式（**推荐**） <br>
   参考文档：https://help.aliyun.com/document_detail/84514.html
2. 代理方式<br>
   参考文档：https://help.aliyun.com/document_detail/91243.html
