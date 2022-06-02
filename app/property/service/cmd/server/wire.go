//go:build wireinject
// +build wireinject

package main

func initApp() (*service.PropertyService, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
