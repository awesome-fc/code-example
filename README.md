
# 简介
本仓库是函数计算各 runtime 相关的示例和最佳实践建议。
- [Python Runtime 代码开发](https://help.aliyun.com/document_detail/74753.html)
- [Nodejs Runtime 代码开发](https://help.aliyun.com/document_detail/74754.html)
- [PHP Runtime 代码开发](https://help.aliyun.com/document_detail/89028.html)
- [Java Runtime 代码开发](https://help.aliyun.com/document_detail/74755.html)
- [Golang Runtime 代码开发](https://help.aliyun.com/document_detail/323505.html)
- [C# Runtime 代码开发](https://help.aliyun.com/document_detail/112377.html)
- [Custom Runtime 代码开发](https://help.aliyun.com/document_detail/132044.html) 。

> 注意： 标准Runtime python/nodejs/java/php/golang 与 custom runtime python/nodejs/golang 是不同的运行时。

## 函数计算快速开始代码示例（FC Quick Start samples）


| 示例名称                          | 运行时  | 类型  | 第三方服务  | 介绍                                                         |
| --------------------------------- | ------- | ------- | ----------- | ------------------------------------------------------------ |
| mns-queue-producer-python3-event | python3.6 | 事件请求处理程序(Event Handler) | 消息队列MNS(队列模型) | 快速部署一个由 Python3.6 事件类型实现的消息服务MNS(队列模型-生产者)示例函数到阿里云函数计算。 |
| mns-queue-producer-python3-http | python3.6 | HTTP请求处理程序（HTTP Handler） |  消息队列MNS(队列模型) |       快速部署一个由 Python3.6 事件类型实现的消息服务MNS(队列模型-生产者)示例函数到阿里云函数计算。|
| mns-queue-trigger-python |  python3.6 | MNS队列触发器（By EventBridge） | 消息队列 MNS(队列模型) | 快速部署一个由 Python3.6 实现的消息服务MNS(队列模型-消费者)触发器函数到阿里云函数计算。 |
| | | |  | |

## FC Runtime 示例

- python
    - python3-mysql: 使用initializer和preStop回调函数的mysql示例程序。
- nodejs
    - nodejs14-mysql: 使用initializer和preStop回调函数的mysql示例程序。
- java
    - java11-blank-stream-event: java11 stream 事件回调示例程序。
    - java11-blank-pojo-event: java11 POJO 事件回调示例程序。
    - java11-blank-http: java11 HTTP回调示例程序。
    - java11-oss: java11 使用临时密钥访问oss示例程序。
    - java11-mysql: 使用initializer和preStop回调函数的mysql示例程序。
- golang
    - golang-oss: 当上传图片至 oss ,触发函数自动缩放图片并重新回传至oss 。
  

## Custom Runtime 示例

- python
  - python-demo-with-lib-in-layer: 在 Custom Runtime(Python) 语言中如何引用层中的依赖

- nodes
  - nodejs-demo-with-lib-in-layer: 在 Custom Runtime(Nodejs) 语言中如何引用层中的依赖

## 注意事项

1. 该示例仅供参考，如果用于正式生产环境，请根据具体应用场景修改验证。
