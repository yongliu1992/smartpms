package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/yongliu1992/smartpms/api"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"net/url"
	"time"
)
import "github.com/smallnest/rpcx/serverplugin"

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
	s.RegisterName("property", new(Property), "")
	app, cleanup, err := initApp()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(app)
	defer cleanup()
	err = s.Serve("tcp", "127.0.0.1:8088")
	if err != nil {
		panic(err)
	}
}

type Property struct{}

func (p *Property) ListUnits(ctx context.Context, req biz.ListUnitsReq, resp *api.Resp) error {
	resp.Data = biz.ListUnitsResp{}.Data
	resp.Code = 200
	resp.Message = "ok"
	return nil
}
func (p *Property) ListShops(ctx context.Context, req Param, resp *api.Resp) error {
	fmt.Println("ccc", req.Url.Get("a"))
	resp.Data = "cccsadas"
	resp.Code = 200
	resp.Message = "ok"
	return nil
}

type Param struct {
	Url  url.Values
	Post *url.Values
}
