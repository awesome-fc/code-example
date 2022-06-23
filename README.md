
# 简介
本仓库是函数计算各 runtime 相关的示例和最佳实践建议。
- Python runtime 代码开发文档: [Python runtime](https://help.aliyun.com/document_detail/74753.html) 。
- Nodejs runtime 代码开发文档: [Nodejs runtime](https://help.aliyun.com/document_detail/74754.html) 。
- Java runtime 代码开发文档: [Java runtime](https://help.aliyun.com/document_detail/74755.html) 。
- Golang runtime 代码开发文档: [Golang runtime](https://help.aliyun.com/document_detail/323505.html) 。
- Custom runtime 代码开发文档: [Custom runtime](https://help.aliyun.com/document_detail/132042.html) 。

## FC Runtime 示例
> 注意： python/nodejs/golang runtime 与 custom python/nodejs/golang runtime 是不同的运行时，custom runtime 文档请参考 [Custom runtime](https://help.aliyun.com/document_detail/132044.html) 。

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
