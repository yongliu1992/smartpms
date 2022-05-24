package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/yongliu1992/smartpms/api"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data"
	"net/url"
	"strconv"
	"time"
)
import "github.com/smallnest/rpcx/serverplugin"
var db *data.Data
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
	dt,_,err := data.InitData()
	db = dt
	if err !=nil {
		panic(err )
	}
	//defer clean()
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
func (p *Property) Shops(ctx context.Context, req Param, resp *api.Resp) error {
	pageSize ,_ :=strconv.Atoi(req.Url.Get("pageSize"))
	data,err := data.NewShopRepo(db).ListShop(ctx,0,pageSize,1)
	fmt.Println(err)
	resp.Data = data
	resp.Code = 200
	resp.Message = "ok"
	return nil
}

type Param struct {
	Url  url.Values
	Post *url.Values
}
