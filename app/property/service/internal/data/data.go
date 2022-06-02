package data

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent"
	"github.com/yongliu1992/smartpms/contrib/config/file"
	"log"
)

type Data struct {
	db *ent.Client
	r  *redis.Client
}

var ProviderSet = wire.NewSet(NewEntClient, NewRedisClient, NewProvideData, NewShopRepo, NewCommunityRepo)

func NewEntClient() *ent.Client {
	c := file.GetConf()
	m := c.Property.Mysql

	mysqlAddress := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DbName + "?parseTime=true"
	client, err := ent.Open(
		"mysql",
		mysqlAddress,
	)
	if err != nil {
		log.Fatal("连接数据库错误", err)
	}
	return client
}
func NewRedisClient() (*redis.Client, error) {
	c := file.GetConf()
	r := c.Property.Redis
	rc := redis.NewClient(&redis.Options{Addr: r.Address, Password: r.Password, DB: r.Db})
	err := rc.Ping(context.Background()).Err()
	return rc, err
}
func NewProvideData(entClient *ent.Client, rc *redis.Client) *Data {
	d := &Data{db: entClient, r: rc}
	return d
}
