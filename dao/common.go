package dao

import (
	"database/sql"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MAX_IDLE       = 30
	MAX_CONN       = 30
	DATABASE_ALIAS = "default"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(DATABASE_ALIAS, "mysql", "root:a1234567@tcp(127.0.0.1:3306)/datastruct?charset=utf8")
	//  根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns(DATABASE_ALIAS, MAX_IDLE)
	// 根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
	orm.SetMaxOpenConns(DATABASE_ALIAS, MAX_CONN)
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 设置查询日志
	orm.Debug = true
}

func GetDB() (*sql.DB, error) {
	db, err := orm.GetDB(DATABASE_ALIAS)
	if err != nil {
		return nil, err
	}
	return db, err
}
