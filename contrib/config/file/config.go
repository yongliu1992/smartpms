package file

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

var co Configs
var synConf sync.Once

func GetConf() Configs {
	synConf.Do(
		func() {
			data, err := ioutil.ReadFile("/Users/kok/go/src/smartpms/contrib/config/file/config.yaml")
			if err != nil {
				panic(err)
			}
			err = yaml.Unmarshal(data, &co)
			if err != nil {
				panic(err)
			}
			fmt.Println(co.Property.Mysql.Host)
		})
	return co
}

type ConfigMysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"db_name"`
}

type ConfigRedis struct {
	Password string `yaml:"password"`
	Address  string `yaml:"address"` //host:port
	Db       int    `yaml:"db"`
}

type Configs struct {
	Property Property `yaml:"property"`
}
type Property struct {
	Mysql ConfigMysql `yaml:"mysql"`
	Redis ConfigRedis `yaml:"redis"`
}
