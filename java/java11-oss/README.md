# java11 oss 示例程序
本示例演示了如何使用临时密钥向OSS中上传一个文件。

com.aliyun.oss 版本见 maven  https://mvnrepository.com/artifact/com.aliyun.oss/aliyun-sdk-oss

oss sdk 使用参考 https://help.aliyun.com/document_detail/84781.html

## 准备开始
- 在指定地域创建 bucket（比如，在北京地域创建名为 `my-bucket`）
- [可选] 安装并配置 Serverless Devs 工具。（https://help.aliyun.com/document_detail/195474.html）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 修改 service 中的 role 配置，设置的role需要需要拥有 AliyunOSSFullAccess 权限，建议直接使用 AliyunFCDefaultRole 角色

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
FC Invoke Start RequestId: 7fe813fa-a81f-43c9-8c5f-xxxxxxxx
FC Invoke End RequestId: 7fe813fa-a81f-43c9-8c5f-xxxxxxxx

Duration: 1105.06 ms, Billed Duration: 1106 ms, Memory Size: 128 MB, Max Memory Used: 121.36 MB
========= FC invoke Logs end =========

FC Invoke Result:
done


End of method: invoke
```

### 方式二、使用控制台创建

#### 1. 编译打包

```shell
# 编译部署
mvn package
# 打包文件
cd target && zip -r java11-oss.zip *
```

#### 2. 创建函数
选择服务（或创建服务）后，单击创建函数，如图所示
- 选择 `从零开始创建`
- 填入函数名称
- 选择运行环境 java11/java8
- 选择函数触发方式：通过事件请求触发
- 其他设置使用默认

![img_1.png](assets/20220408110824.jpg)

> 详细创建函数流程见文档: [使用控制台创建函数](https://help.aliyun.com/document_detail/51783.html)

#### 3. 设置服务的角色配置
编辑服务的配置，将角色设置为AliyunOSSFullAccess，或者创建新的服务角色，但要保证角色中有 AliyunOSSFullAccess 权限策略。

#### 4. 测试函数

返回结果如下所示
```bash
done
```