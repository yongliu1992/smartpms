//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data"
)

func initApp() (*data.Data, func(), error) {
	panic(wire.Build(data.ProviderSet))
}
