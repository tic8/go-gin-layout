package config

import (
	"github.com/axiaoxin-com/logging"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func InitMySQL() *gorm.DB {
	// 从配置中获取 DSN
	dsn := Cfg.Database.DSN

	// 配置 GORM 的选项
	gormConfig := &gorm.Config{
		Logger: logging.NewGormLogger(zap.InfoLevel, zap.InfoLevel, time.Second*3), // 设置日志级别
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名使用单数形式
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if Cfg.Database.Debug {
		db = db.Debug() // 如果配置了 Debug，则启用调试模式
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(100)                 // 设置最大连接数
	sqlDB.SetMaxIdleConns(10)                  // 设置最大空闲连接数
	sqlDB.SetConnMaxLifetime(10 * time.Minute) // 设置连接的最大生命周期

	return db
}
