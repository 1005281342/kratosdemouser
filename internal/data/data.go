package data

import (
	"github.com/1005281342/kratosdemouser/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	redigo "github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKratosDemoUserRepo)

// Data .
type Data struct {
	db        *gorm.DB
	redisConn redigo.Conn
	logger    *log.Helper
}

const (
	mySQL = "mysql"
)

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	var (
		logg = log.NewHelper(logger)
		err  error
		db   *gorm.DB
		data = &Data{}
		conn redigo.Conn
	)
	cleanup := func() {
		logg.Info("closing the data resources")
	}
	switch c.Database.Driver {
	case mySQL:
		if db, err = gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{}); err != nil {
			logg.Errorf("连接DB失败，%s", err.Error())
			return data, cleanup, err
		}
		data.db = db
	}
	if conn, err = redigo.Dial(c.Redis.Network, c.Redis.Addr); err != nil {
		logg.Errorf("连接Redis失败，%s", err.Error())
		return data, cleanup, err
	}
	data.redisConn = conn
	data.logger = logg
	return data, cleanup, nil
}
