
# go-gin-layout

## 项目简介
`go-gin-layout` 是一个基于 [Gin](https://github.com/gin-gonic/gin) 框架的 Go 项目模板，集成了常用的功能和工具，旨在帮助开发者快速构建高效、可维护的 Web 应用程序。

## 功能特性
- **配置管理**：使用 [Viper](https://github.com/spf13/viper) 进行灵活的配置管理。
- **日志记录**：集成 [Zap](https://github.com/uber-go/zap) 提供高性能日志记录。
- **数据库支持**：使用 [GORM](https://gorm.io/) 和 MySQL 驱动进行数据库操作。
- **Redis 集成**：支持 Redis 缓存，使用 [go-redis](https://github.com/redis/go-redis)。
- **JWT 认证**：集成 [golang-jwt](https://github.com/golang-jwt/jwt) 实现用户认证。
- **定时任务**：支持基于 [Cron](https://github.com/robfig/cron) 的定时任务调度。
- **Swagger 文档**：集成 [Swaggo](https://github.com/swaggo/gin-swagger) 自动生成 API 文档。
- **Prometheus 监控**：支持 [Prometheus](https://github.com/prometheus/client_golang) 指标监控。

## 快速开始

### 环境要求
- Go 1.23 或更高版本
- MySQL 数据库
- Redis 缓存

### 安装依赖
```bash
go mod tidy
```

### 配置文件
修改 `config.yaml` 文件，根据实际需求调整以下配置：
```yaml
server:
  port: 8080

database:
  dsn: root:password@tcp(127.0.0.1:3306)/your_database?charset=utf8mb4&parseTime=True&loc=Local

redis:
  addr: localhost:6379
  password: ""
  db: 0

jwt:
  secret: your_jwt_secret
  expire: 3600
```

### 启动项目
```bash
go run cmd/main.go
```

启动后，服务将运行在 `http://localhost:8080`。

### Swagger 文档
访问 `http://localhost:8080/swagger/index.html` 查看自动生成的 API 文档。

### Prometheus 指标
访问 `http://localhost:8080/metrics` 查看 Prometheus 指标。

## 项目结构
```
go-gin-layout/
├── cmd/                # 主程序入口
├── internal/           # 内部模块
│   ├── config/         # 配置管理
│   ├── handler/        # 路由处理
│   ├── middleware/     # 中间件
│   ├── model/          # 数据模型
│   ├── service/        # 业务逻辑
│   └── util/           # 工具函数
├── config.yaml         # 配置文件
├── go.mod              # 依赖管理
├── README.md           # 项目说明
```

## 依赖列表
主要依赖：
- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Zap](https://github.com/uber-go/zap)
- [GORM](https://gorm.io/)
- [go-redis](https://github.com/redis/go-redis)
- [golang-jwt](https://github.com/golang-jwt/jwt)
- [Swaggo](https://github.com/swaggo/gin-swagger)
- [Prometheus](https://github.com/prometheus/client_golang)

## 贡献
欢迎提交 Issue 和 Pull Request 来改进本项目。

## 许可证
本项目基于 [MIT License](LICENSE) 开源。
```