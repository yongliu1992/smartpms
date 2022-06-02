// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/yongliu1992/smartpms/app/property/service/internal/biz"
	"github.com/yongliu1992/smartpms/app/property/service/internal/data"
	"github.com/yongliu1992/smartpms/app/property/service/internal/service"
)

// Injectors from wire.go:

func initApp() (*service.PropertyService, error) {
	client := data.NewEntClient()
	redisClient, err := data.NewRedisClient()
	if err != nil {
		return nil, err
	}
	dataData := data.NewProvideData(client, redisClient)
	communityRepo := data.NewCommunityRepo(dataData)
	communityUseCase := biz.NewCommunityUseCase(communityRepo)
	shopsRepo := data.NewShopRepo(dataData)
	shopUseCase := biz.NewShopUseCase(shopsRepo)
	propertyService := service.NewPropertyService(communityUseCase, shopUseCase)
	return propertyService, nil
}