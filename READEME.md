```shell
your_project/
├── bin/               # 编译后的二进制文件
├── cmd/               # 主应用程序的入口点
│   └── your_app/      # 应用的主程序
│       └── main.go    # main 函数
├── config/            # 配置文件
│   └── config.yaml    # 项目的配置文件
├── docs/              # 项目文档
├── internal/          # 私有应用程序逻辑（不可导出的包）
│   ├── controller/    # 控制器层，用于处理 HTTP 请求
│   ├── model/         # 数据库模型
│   ├── service/       # 服务层，封装业务逻辑
│   └── repository/    # 数据库操作层
├── pkg/               # 可导出的库和工具包
├── routes/            # 路由定义
│   └── router.go      # 路由配置
├── scripts/           # 启动、部署等脚本
├── static/            # 静态资源（HTML、JS、CSS等）
├── test/              # 单元测试
│   └── controller_test.go  # 测试文件
├── go.mod             # 依赖管理
└── README.md          # 项目说明文件
```