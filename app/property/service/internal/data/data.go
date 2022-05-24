package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data/ent"
	"github.com/yongliu1992/smartpms/contrib/config/file"
	"log"
)


type Data struct {
	db *ent.Client
}

func NewEntClient() *ent.Client {
	c := file.GetConf()
	m := c.Property.Mysql

	mysqlAddress := m.User + ":"+m.Password+"@tcp(" + m.Host + ":" + m.Port + ")/" + m.DbName + "?parseTime=true"
	client, err := ent.Open(
		"mysql",
		mysqlAddress,
	)
	if err != nil {
		log.Fatal("连接数据库错误",err)
	}
	return client
}
func ProvideData(entClient *ent.Client) *Data {
	d := &Data{db: entClient}
	return d
}

