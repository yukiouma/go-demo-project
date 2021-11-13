# A Bookstore Demo Project 

基于Gin和GRPC实现一个书店项目的demo，包括两个服务，book（书本管理）和customer（顾客管理），每个服务的http服务使用gin来实现，同时提供对应的RPC服务使用GRPC框架实现。book服务会调用customer的GRPC服务来验证和记录购买书籍的用户信息



## 需求

按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。 



## 框架使用

* Gin
* GRPC



## 使用说明



启动book服务

```bash
make book
```

控制台出现以下内容表示启动成功

```bash
go run ./app/book/cmd
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/book/v1/find/:id     --> frame/api/book/v1.FindBookTransfer.func1 (3 handlers)
[GIN-debug] POST   /api/book/v1/new          --> frame/api/book/v1.NewBookTransfer.func1 (3 handlers)
[GIN-debug] POST   /api/book/v1/sale         --> frame/api/book/v1.SaleBookTransfer.func1 (3 handlers)
[GIN-debug] DELETE /api/book/v1/delete/:id   --> frame/api/book/v1.DeleteBookTransfer.func1 (3 handlers)
```



启动customer服务

```bash
make customer
```

控制台出现以下内容表示启动成功

```bash
go run ./app/customer/cmd
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/customer/v1/find/:id --> frame/api/customer/v1.FindCustomerTransfer.func1 (3 handlers)
[GIN-debug] POST   /api/customer/v1/register --> frame/api/customer/v1.RegisterCustomerTransfer.func1 (3 handlers)
[GIN-debug] POST   /api/customer/v1/update   --> frame/api/customer/v1.UpdateCustomerTransfer.func1 (3 handlers)
[GIN-debug] DELETE /api/customer/v1/remove/:id --> frame/api/customer/v1.RemoveCustomerTransfer.func1 (3 handlers)
```



运行测试脚本

```bash
make testcase
```

```bash
[root@playground go-demo-project]# make testcase
sh ./test/script.sh
creating customer...
{"data":{"id":1,"name":"yuki"},"message":"Register customer successfully"}
creating book...
{"data":{"id":1,"name":"golang","saleInfo":{}},"message":"Putting a book on the shelf successfully"}
finding book...
{"data":{"id":1,"name":"golang","saleInfo":{"saledAt":"0001-01-01 00:00:00 +0000 UTC"}},"message":"Getting book successfully"}
finding customer...
{"data":{"id":1,"name":"yuki"},"message":"Getting customer successfully"}
saling book...
{"data":{"id":1,"name":"golang","saleInfo":{"saledAt":"2021-11-13 13:56:56.434336579 +0800 CST m=+7.387403101","customerId":1,"customerName":"yuki"}},"message":"Saling book successfully"}
ending test...
```

测试脚本执行的内容为

1. 新建一个顾客信息
2. 新建一本书籍信息
3. 查找id为1的书籍信息
4. 查找id为1的顾客信息
5. id为1的顾客购买了id为1的书籍，id为1的书籍信息得到更新，追加上了购买时间与购买的顾客的信息



## 项目结构

最大程度参考了kratos的标准结构

```bash
.
├── api
│   ├── book
│   │   └── v1
│   │       ├── book_grpc.pb.go
│   │       ├── book_http.pb.go
│   │       ├── book.pb.go
│   │       └── book.proto
│   └── customer
│       └── v1
│           ├── customer_grpc.pb.go
│           ├── customer_http.pb.go
│           ├── customer.pb.go
│           └── customer.proto
├── app
│   ├── book
│   │   ├── cmd
│   │   │   ├── main.go
│   │   │   ├── wire_gen.go
│   │   │   └── wire.go
│   │   ├── configs
│   │   │   └── config.json
│   │   └── internal
│   │       ├── biz
│   │       │   ├── biz.go
│   │       │   ├── book.go
│   │       │   └── customer.go
│   │       ├── conf
│   │       │   └── conf.go
│   │       ├── data
│   │       │   ├── book.go
│   │       │   ├── customer.go
│   │       │   └── data.go
│   │       ├── server
│   │       │   ├── grpc.go
│   │       │   ├── http.go
│   │       │   └── server.go
│   │       └── service
│   │           ├── book.go
│   │           └── service.go
│   └── customer
│       ├── cmd
│       │   ├── main.go
│       │   ├── wire_gen.go
│       │   └── wire.go
│       ├── configs
│       │   └── config.json
│       └── internal
│           ├── biz
│           │   ├── biz.go
│           │   └── customer.go
│           ├── conf
│           │   └── conf.go
│           ├── data
│           │   ├── customer.go
│           │   └── data.go
│           ├── server
│           │   ├── grpc.go
│           │   ├── http.go
│           │   └── server.go
│           └── service
│               ├── customer.go
│               └── service.go
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   └── appmanage
│       ├── app.go
│       ├── config.go
│       └── signal.go
├── README.md
└── test
    └── script.sh

```



### api/\<service_name\>

存放每个服务对外提供的接口形状

使用Proto文件进行接口定义

grpc服务的代码桩使用protoc生成

http服务的代码桩则手动编写根据\<service\>.pb.go中的定义转换为相对应的gin.HandleFunc转换函数。http服务对外的restful接口的路径则直接映射服务所在的目录位置，如`./api/book/v1/book_http.pb.go`中记载的路由方法的前缀均为`/api/book/v1`



### app/\<service_name\>

存放服务的具体实现代码，使用以下分层



#### cmd

* main.go

  项目启动入口

* wire_gen.go

  依赖注入实现

* wire.go

  依赖注入定义



#### internal

服务具体实现代码



#### configs

存放服务定义的文件，采用json格式



##### biz

存放具体业务逻辑的领域层，定义以下内容

* 领域对象
* 仓储层接口
* 用例
* 请求其它服务的客户端的接口定义

与目录同名的文件存放wire注册信息



##### data

定义以下内容：

* 仓储的具体实现

* 访问其它服务的客户端的具体实现

与目录同名的文件存放wire注册信息



##### conf

配置文件的读取方法



##### service

对api接口的具体实现，并调用biz中的领域用例来完成接口定义的具体服务

与目录同名的文件存放wire注册信息



##### server

实现了http服务与rpc服务的路由注册与启动方法

与目录同名的文件存放wire注册信息



### pkg

定义了以下的服务公用的方法

* 项目启动
* 服务信息注册
* 配置读取
* 不同类型的服务的同时退出逻辑



### test

定义了一个测试脚本

测试脚本执行的内容为

1. 新建一个顾客信息
2. 新建一本书籍信息
3. 查找id为1的书籍信息
4. 查找id为1的顾客信息
5. id为1的顾客购买了id为1的书籍，id为1的书籍信息得到更新，追加上了购买时间与购买的顾客的信息



## 实现思路



### 配置信息

使用json文件存储配置信息，读取时存放到一个map中

服务级别的配置读取方法，则基于json的读取结果，从map中获取自身所需要的信息



### api定义

使用proto文件定义具体服务的接口形状

使用protoc来获取对应的定义文件`*.pb.go`与grpc服务的文件`*_prpc.pb.go`

使用grpc服务的文件中的服务定义`type <serviceName>ServiceClient interface`，编写一个基于Gin框架的http服务定义文件`*_http.pb.go`



### 依赖注入

通过服务中的每一层使用wire.NewSet方法暴露出来的构造器函数，然后在cmd层的wire.go文件中使用暴露的构造器函数中的依赖定义，对各层的服务进行组装



### 服务管理

对每个服务中暴露出来的http和grpc服务，使用errgroup结合context进行管理，具体逻辑为当某个服务发生错误导致退出时，会激活传入服务中的context的cancel方法，使得其它服务可以被一并终止。



## Bug残留

* 还没有实现一个注册中心的项目，目前不同的服务之间进行沟通的时候，填写的域名或者ip地址时通过config.json进行暴露的
* app注册信息是假的....注册方法只是接受了入参但是并没有一个具体的效果....