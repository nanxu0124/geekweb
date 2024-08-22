package ioc

import (
	"fmt"
	"geekweb/internal/repository/dao"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	type config struct {
		Dsn string `yaml:"dsn"`
	}
	var c config
	err := viper.UnmarshalKey("mysql", &c)
	if err != nil {
		panic(fmt.Errorf("初始化配置失败 %v", err))
	}

	db, err := gorm.Open(mysql.Open(c.Dsn))
	if err != nil {
		panic(err)
	}
	err = dao.INitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
