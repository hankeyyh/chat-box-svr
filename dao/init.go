package dao

import (
	"github.com/hankeyyh/chat-box-svr/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// open db
	var err error
	mysqlConf := conf.DefaultConf.MysqlConf
	if db, err = gorm.Open(mysql.Open(mysqlConf.GetDsn())); err != nil {
		panic(err)
	}
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(mysqlConf.MaxIdleConn)
	sqldb.SetMaxOpenConns(mysqlConf.MaxOpenConn)

	// set default model
	SetDefault(db)
}