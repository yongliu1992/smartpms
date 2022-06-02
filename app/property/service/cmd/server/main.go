package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
	"github.com/yongliu1992/smartpms/api"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data"
	"github.com/yongliu1992/smartpms/log"
	"net/url"
	"strconv"
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
	s.RegisterName(appName, new(Property), "")
	rp, _, err := data.InitData()
	repo = rp
	if err != nil {
		panic(err)
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
	pageSize, _ := strconv.Atoi(req.Url.Get("pageSize"))
	data, err := biz.NewShopUseCase(data.NewShopRepo(repo)).ListShop(ctx, 0, pageSize, 1)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		resp.Data = nil
		return nil
	}
	resp.Data = data
	resp.Code = 200
	resp.Message = "ok"
	return nil
}

func (p *Property) ListComm(ctx context.Context, req Param, resp *api.Resp) error {
	req2 := biz.ListCommunityReq{}
	data, err := biz.NewCommunityUseCase(data.NewCommunityRepo(repo)).List(ctx, req2)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		resp.Data = nil
		return nil
	}
	resp.Data = data
	resp.Code = 200
	resp.Message = "ok"
	return nil
}
func (p *Property) CreateCommunity(ctx context.Context, req Param, resp *api.Resp) error {

	var reqData biz.Community
	reqData.Name = req.Post.Get("name")
	data, err := biz.NewCommunityUseCase(data.NewCommunityRepo(repo)).Add(ctx, reqData)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		resp.Data = nil
		return nil
	}
	resp.Data = data
	resp.Code = 200
	resp.Message = "ok"
	return nil
}

type Param struct {
	Url  url.Values
	Post *url.Values
}
