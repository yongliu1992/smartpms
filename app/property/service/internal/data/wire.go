//go:build wireinject

package data

import "github.com/google/wire"

func InitData() (*Data, func(), error) {
	panic(wire.Build(ProvideData, NewEntClient, NewRedisClient))
	return &Data{}, func() {
	}, nil
}
