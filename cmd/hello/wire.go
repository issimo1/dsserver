//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/issimo1/dsserver/internal/biz"
	"github.com/issimo1/dsserver/internal/conf"
	"github.com/issimo1/dsserver/internal/data"
	"github.com/issimo1/dsserver/internal/server"
	"github.com/issimo1/dsserver/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
