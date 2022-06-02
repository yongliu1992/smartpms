package service

import (
	"context"
	"github.com/yongliu1992/smartpms/api"
	"strconv"
)

func (p *PropertyService) Shops(ctx context.Context, req Param, resp *api.Resp) error {
	pageSize, _ := strconv.Atoi(req.Url.Get("pageSize"))
	data, err := p.shop.ListShop(ctx, 0, pageSize, 1)
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
