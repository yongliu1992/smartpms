package service

import (
	"context"
	"github.com/yongliu1992/smartpms/api"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
)

func (p *PropertyService) ListUnits(ctx context.Context, req biz.ListUnitsReq, resp *api.Resp) error {
	resp.Data = biz.ListUnitsResp{}.Data
	resp.Code = 200
	resp.Message = "ok"
	return nil
}
