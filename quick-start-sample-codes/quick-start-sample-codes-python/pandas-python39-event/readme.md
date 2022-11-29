# 应用开发说明

<p align="center"><b> 中文 | <a href="./readme_en.md"> English </a>  </b></p>


> Serverless Devs 应用开发需要严格遵守 [Serverless Package Model](../../spec/zh/0.0.2/serverless_package_model/readme.md) 中的 [应用模型规范](../../spec/zh/0.0.2/serverless_package_model/3.package_model.md#应用模型规范)。在[应用模型规范](../../spec/zh/0.0.2/serverless_package_model/3.package_model.md#应用模型规范)中有关于[应用模型元数据](../../spec/zh/0.0.2/serverless_package_model/3.package_model.md#应用模型元数据)的说明。

Serverless Devs的组件开发案例已经被集成到Serverless Devs命令行工具中，通过对Serverless Devs的命令行工具，可以进行空白应用项目的初始化，开发者只需要执行`s init`即可看到：

```shell script

🚀 Serverless Awesome: https://github.com/Serverless-Devs/package-awesome

? Hello Serverless for Cloud Vendors (Use arrow keys or type to search)
❯ Alibaba Cloud Serverless 
  AWS Cloud Serverless 
  Tencent Cloud Serverless 
  Baidu Cloud Serverless 
  Dev Template for Serverless Devs 
```

此时，选择最后的`Dev Template for Serverless Devs`，并按回车：

```shell script
$ s init

🚀 Serverless Awesome: https://github.com/Serverless-Devs/package-awesome

? Hello Serverless for Cloud Vendors Dev Template for Serverless Devs
? Please select an Serverless-Devs Application (Use arrow keys or type to search)
❯ Application Scaffolding 
  Component Scaffolding 
```

此时，选择`Application Scaffolding`，并按回车，即可完成一个完整的Serverless Devs的Application项目的初始化，可以通过命令查看文件树：

```shell script
$ find . -print | sed -e 's;[^/]*/;|____;g;s;____|; |;g'
.
|____readme.md
|____version.md
|____publish.yaml
|____src
| |____s.yaml
| |____index.js
```

这其中：

| 目录 | 含义 |
| --- | --- | 
| readme.md | 对该组件的描述，或帮助文档信息 | 
| version.md | 版本的描述，例如当前版本的更新内容等 |  
| publish.yaml | 项目所必须的文件，Serverless Devs Package的开发识别文档 |
| src | 应用所在目录，需要包括`s.yaml`和相关的应用代码等 | 


此时，开发者可以在src下完成应用的开发，并对项目进行`publish.yaml`文件的编写。完成之后，即可将项目发不到不同的源，以Github Registry为例，可以在Github创建一个`Public`的仓库，并将编译后的代码放到仓库，并发布一个版本。此时，就可以通过客户端获取到该应用。


