package service

import (
	"context"
	"github.com/yongliu1992/smartpms/api"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
)

func (p *PropertyService) ListComm(ctx context.Context, req Param, resp *api.Resp) error {
	req2 := biz.ListCommunityReq{}
	data, err := p.common.List(ctx, req2)
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
func (p *PropertyService) CreateCommunity(ctx context.Context, req Param, resp *api.Resp) error {

	var reqData biz.Community
	reqData.Name = req.Post.Get("name")
	data, err := p.common.Add(ctx, reqData)
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
