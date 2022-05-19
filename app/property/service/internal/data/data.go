package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent"
	"github.com/yongliu1992/smartpms/contrib/config/file"
	"log"
)

var ProviderSet = wire.NewSet(NewData, NewEntClient, NewShopRepo)

type Data struct {
	db *ent.Client
}

func NewEntClient() *ent.Client {
	c := file.GetConf()
	m := c.Property.Mysql
	mysqlAddress := m.User + "@tcp(" + m.Host + ":" + m.Port + ")" + m.DbName + "?parseTime=true"
	client, err := ent.Open(
		"mysql",
		mysqlAddress, nil,
	)
	if err != nil {
		log.Fatal("连接数据库错误")
	}
	return client
}
func NewData(entClient *ent.Client) (*Data, func(), error) {
	d := &Data{db: entClient}
	return d, func() {
		if err := d.db.Close(); err != nil {
			fmt.Println(err)
		}
	}, nil
}
