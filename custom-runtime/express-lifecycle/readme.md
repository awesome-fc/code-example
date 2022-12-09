# nodeJS Express 函数实例生命周期回调示例

本示例为您展示了 Custom Runtime 的 Express 使用示例。本示例将使用 Express 实现 函数实例生命周期回调。

## 准备开始

- [可选] 安装并配置 Serverless Devs 工具。（ https://help.aliyun.com/document_detail/195474.html ）

## 快速开始
### 方式一、使用 Serverless Devs 工具编译部署

#### 1. 修改 s.yaml 配置
- 根据需要修改 access 配置
- 指定端口。

#### 2. 安装依赖并部署

编译部署代码包
```shell
s deploy
```
> 注意: `s.yaml` 中有配置 `pre-deploy` 脚本  `npm install --production --registry=https://registry.npmmirror.com`, 在部署前会调用该命令安装依赖。

#### 3. 调用测试

```shell
curl https://express-web-xxxxxx.cn-hangzhou.fcapp.run
```
开启函数计算调用日志，可以看到如下日志

```bash
FC Initialize Start RequestId:b192fdb1-0acf-4604-9d0e-48b51b1614f4
FC Initialize End RequestId:b192fdb1-0acf-4604-9d0e-48b51b1614f4
FC Invoke Start RequestId: b192fdb1-0acf-4604-9d0e-48b51b1614f4
FC Invoke End RequestId: b192fdb1-0acf-4604-9d0e-48b51b1614f4
FC prefreeze Start RequestId:b192fdb1-0acf-4604-9d0e-48b51b1614f4
FC prefreeze End RequestId:b192fdb1-0acf-4604-9d0e-48b51b1614f4
FC prestop Start RequestId:af1a53d7-791a-45d5-a4d1-3f5c9898106a
FC prestop End RequestId:af1a53d7-791a-45d5-a4d1-3f5c9898106a
```

### 方式二、使用控制台创建
todo...



