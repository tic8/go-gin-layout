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
- Go 1.22 或更高版本
- MySQL 数据库
- Redis 缓存

### 安装依赖
```bash
go mod tidy