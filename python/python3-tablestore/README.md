# Python3 表格存储示例

本示例为您展示了 Python runtime 的表格存储使用示例。
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
- 使用 initialize 回调，需要在 s.yaml 中配置

```yaml
        initializationTimeout: 20
        initializer: index.initialize
```

#### 2. 安装依赖并部署

安装依赖
```shell
# 使用s工具安装依赖，需要使用 docker
s build --use-docker
```
> 注意: 使用 `s build --use-docker` 会将依赖安装到代码包的 `.s/python` 目录下，可以在本地的 `.s/build/artifacts/{serviceName}/{functionName}/` 查看。

部署代码
```shell
# 部署代码，会自动添加环境变量 PYTHONUSERBASE=/code/.s/python，该环境变量是必须的
s deploy
```

#### 3. 调用测试

```shell
s invoke
```

调用函数时收到的响应如下所示：

```bash
========= FC invoke Logs begin =========
FunctionCompute python3 runtime inited.
FC Initialize Start RequestId: 58******
FC Initialize End RequestId: 58******
FC Invoke Start RequestId: 58******
FC Invoke End RequestId: 58******

Duration: 68.86 ms, Billed Duration: 69 ms, Memory Size: 128 MB, Max Memory Used: 48.77 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-******

FC Invoke Result:
[('age', '20', 1657531733801), ('home', '北京', 1657618107569), ('name', '张三', 1657531733801)]


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

