package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/timchine/jxc/pkg/app"
	log "github.com/timchine/jxc/pkg/log"
	"github.com/timchine/jxc/router"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// @title			进销存系统
// @version		1.0
// @description	以实现无纸化办公为目标
// @host
// @BasePath	/api/jxc
func main() {
	var (
		db *gorm.DB
	)
	a, err := app.NewApp("jxc", log.ModeDev, zapcore.InfoLevel)
	if err != nil {
		log.Logger().Error(err.Error())
		return
	}
	defer a.Close()
	err = a.AddStageFunc(initConfig).Run()
	if err != nil {
		return
	}
	err = a.AddStageFunc(initDatabase(&db)).Run()
	if err != nil {
		return
	}
	err = a.AddDaemonFunc(router.Run(db)).Run()
	if err != nil {
		return
	}
}

func initDatabase(db **gorm.DB) app.StageFunc {
	return func(ctx context.Context) (app.CleanFunc, error) {
		//连接数据库
		var (
			user     = viper.GetString("mysql.username")
			password = viper.GetString("mysql.password")
			host     = viper.GetString("mysql.host")
			port     = viper.GetString("mysql.port")
			jxc      = viper.GetString("mysql.dbname.jxc")
			maxIdle  = viper.GetInt("mysql.maxIdleConn")
			maxOpen  = viper.GetInt("mysql.maxOpenConn")
			err      error
		)
		newLogger := logger.New(
			log.NewOrmLog(),
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			},
		)
		*db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, port, jxc)), &gorm.Config{NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
			Logger: newLogger})
		if err != nil {
			return nil, err
		}
		sqlDb, err := (*db).DB()
		if err != nil {
			return nil, err
		}
		sqlDb.SetMaxIdleConns(maxIdle)
		sqlDb.SetMaxOpenConns(maxOpen)
		return func() error {
			return sqlDb.Close()
		}, nil
	}
}

func initConfig(ctx context.Context) (app.CleanFunc, error) {
	var (
		err error
	)
	// todo 创建文件
	//初始化配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
