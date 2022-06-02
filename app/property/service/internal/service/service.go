package service

import (
	"github.com/google/wire"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"net/url"
)

var ProviderSet = wire.NewSet(NewPropertyService)

type PropertyService struct {
	common *biz.CommunityUseCase
	shop   *biz.ShopUseCase
}

func NewPropertyService(co *biz.CommunityUseCase, shop *biz.ShopUseCase) *PropertyService {
	return &PropertyService{
		common: co,
		shop:   shop,
	}
}

type Param struct {
	Url  url.Values
	Post *url.Values
}
