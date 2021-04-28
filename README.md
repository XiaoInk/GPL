# GPL - Go/Gin Project Layout

这是一个 **Go/Gin** 的项目结构模版，实际项目开发中可以使用这种代码组织结构。

## 项目结构

```json
.
├── LICENSE
├── go.mod
├── go.sum
├── README.md                   // 项目自述文档
├── docker                      // 存放部署相关代码
├── bin                         // 存放编译后的可执行文件
│   └── main
├── cmd                         // 存放项目入口程序，不应有业务逻辑出现  
│   └── main.go
├── dao                         // 存放与数据库交互的代码，不应有业务逻辑出现
│   ├── token.go
│   └── user.go
├── doc                         // 存放项目相关文档
│   └── api.md
├── handler                     // 业务接口
│   ├── base.go
│   ├── login.go
│   └── logout.go
├── logic                       // 业务逻辑
│   ├── login.go
│   └── logout.go
├── middleware                  // 中间件
│   ├── auth.go
│   └── cors.go
├── model                       // 存放模型相关
│   ├── config.go
│   ├── mysql.go
│   └── table                   // 存放数据表
│       └── user.go
├── router                      // 存放路由配置
│   └── router.go
├── service                     // 存放任何业务都可调用的服务
│   └── domain.go
└── util                        // 存放一些小工具
    └── base.go
```
