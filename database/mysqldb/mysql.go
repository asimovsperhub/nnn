package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/dnsdao/nft.pass/config"
	"sync"
	// mysql 驱动
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	Mysql *sql.DB
}

var (
	mdb     *MysqlDb
	mdbOnce sync.Once
)

func GetMdb() *MysqlDb {
	if mdb == nil {
		mdbOnce.Do(func() {
			d := config.GetRConf().GetMysqlDb()
			mysqldns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
				d.User, d.Passwd, d.Host, d.DbName)
			fmt.Println(mysqldns)
			db, err := sql.Open(d.Driver, mysqldns)
			if err != nil {
				panic(err)
			}
			mdb = &MysqlDb{
				Mysql: db,
			}

		})
	}

	return mdb
}

//// 读配置文件
//func NewMysqlDb() *DbConn {
//
//	dbconf := config.GetRConf().GetMysqlDb()
//	return &DbConn{user: dbconf.User,
//		passwd: dbconf.Passwd,
//		host:   dbconf.Host,
//		dbName: dbconf.DbName,
//		driver: dbconf.Driver}
//}

// 传值
//func NewMysqlDbPrams(conf *config.DatabaseConf) *DbConn {
//	return &DbConn{user: conf.User,
//		passwd: conf.Passwd,
//		host:   conf.Host,
//		dbName: conf.DbName,
//		driver: conf.Driver}
//}

//func (d *DbConn) Connect() error {
//	mysqldns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
//		d.user, d.passwd, d.host, d.dbName)
//	fmt.Println(mysqldns)
//	db, err := sql.Open(d.driver, mysqldns)
//	if err != nil {
//		return err
//	} else {
//		d.DB = db
//		return nil
//	}
//}

func (d *MysqlDb) Close() error {
	return d.Close()
}
