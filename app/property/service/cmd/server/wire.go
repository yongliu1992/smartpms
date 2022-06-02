package main

import (
	"github.com/google/wire"
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data"
	"github.com/yongliu1992/smartpms/app/property/service/internal/service"
)

func initApp() (*service.PropertyService, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
