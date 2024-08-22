### GeekWeb
这是一个基于 Geek Web 框架开发的 Web 项目，主要实现了用户注册功能。项目使用了 MySQL 作为数据库，并集成了 Wire 进行依赖注入，Viper 进行配置管理，以及 Docker 用于启动和管理 MySQL 容器。

### 功能特性
- 用户注册：实现了基本的用户注册功能，支持用户信息的存储和管理。
- 依赖注入：使用 Wire 进行依赖注入，简化代码结构和依赖管理。
- 配置管理：通过 Viper 进行应用配置的加载和管理，支持多环境配置。
- 容器化：使用 Docker 启动 MySQL 数据库，确保开发和生产环境的一致性。

### 项目结构
~~~text
├── config
│   └── dev.yaml          # 配置文件
├── internal
│   ├── domain            # 领域对象
│   ├── repository        # 数据访问层
│   ├── service           # 业务逻辑层
│   └── web               # HTTP 处理器
│
├── ioc                   # 依赖注入相关代码
├── wire                  # 依赖注入相关代码
├── wire_gen              # 依赖注入相关代码
├── main                  # 程序入口
├── script
│   └── mysql             # 数据库初始化
└── docker-compose.yml    # Docker Compose 配置
~~~