package main

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data"
	"github.com/yongliu1992/smartpms/log"
	"time"
)
import "github.com/smallnest/rpcx/serverplugin"

var repo *data.Data

func main() {
	s := server.NewServer()
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ZooKeeperServers: []string{"127.0.0.1:2181"},
		ServiceAddress:   "tcp@127.0.0.1:8088",
		UpdateInterval:   time.Minute,
		BasePath:         "smartPms",
	}
	r.Start()
	s.Plugins.Add(r)
	appName := "property"
	log.InitZap(appName)
	sp, err := initApp()
	if err != nil {
		fmt.Println(err)
	}
	s.RegisterName(appName, sp, "")
	//defer clean()
	err = s.Serve("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
}
