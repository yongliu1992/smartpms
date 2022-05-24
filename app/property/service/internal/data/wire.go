//go:build wireinject
package data

import "github.com/google/wire"

func InitData() (*Data, func(),error) {
	panic(wire.Build(ProvideData, NewEntClient))
	return &Data{}, func() {
	} ,nil
}
