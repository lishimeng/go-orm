package persistence

import (
	"github.com/beego/beego/v2/client/orm"
)

// DriverType RegisterModel
type DriverType orm.DriverType

type Driver struct {
	name string
	t    orm.DriverType
}

type BaseConfig struct {
	initDb     bool
	aliasName  string
	driver     Driver
	dataSource string
	params     []orm.DBOption
	models     []interface{}
}

func (b *BaseConfig) MaxIdle(n int) {
	if n > 0 {
		b.params = append(b.params, orm.MaxIdleConnections(n))
	}
}

func (b *BaseConfig) MaxConn(n int) {
	if n > 0 {
		b.params = append(b.params, orm.MaxOpenConnections(n))
	}
}

var DriverMysql = Driver{"mysql", orm.DRMySQL}
var DriverSqlite = Driver{"sqlite3", orm.DRSqlite}
var DriverOracle = Driver{"oracle", orm.DROracle}
var DriverPostgres = Driver{"postgres", orm.DRPostgres}
var DriverTiDB = Driver{"tidb", orm.DRTiDB}

func (b *BaseConfig) RegisterModel(models ...interface{}) {
	b.models = append(b.models, models...)
}

func InitOrm(config BaseConfig) (err error) {

	err = orm.RegisterDriver(config.driver.name, config.driver.t)
	if err == nil {
		if len(config.models) > 0 {
			orm.RegisterModel(config.models...)
		}
		err = orm.RegisterDataBase(config.aliasName, config.driver.name, config.dataSource, config.params...)
		if err == nil {
			if config.initDb {
				err = orm.RunSyncdb(config.aliasName, false, true)
			}
		}
	}

	return
}
