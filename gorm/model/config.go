package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	return
}

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func Gorm() *gorm.DB {
	var m Mysql
	m.Username = "root"
	m.Password = "123456"
	m.Dbname = "testgorm"
	m.Config = "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
