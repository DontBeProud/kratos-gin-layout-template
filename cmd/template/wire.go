//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"layout_template/internal/biz/template_biz"
	"layout_template/internal/conf/template_config"
	"layout_template/internal/data/template_data"
	"layout_template/internal/server/template_server"
	"layout_template/internal/service/template_service"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*template_config.ServerConfig, *template_config.DataSourceCfg, *template_config.LoggerConfig) (*kratos.App, func(), error) {
	panic(wire.Build(template_server.ProviderSet, template_data.ProviderSet, template_biz.ProviderSet, template_service.ProviderSet, newApp))
}
