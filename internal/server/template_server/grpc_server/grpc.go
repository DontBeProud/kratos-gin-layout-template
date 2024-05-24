package grpc_server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"layout_template/api/template"
	"layout_template/internal/conf/template_config"
	"layout_template/internal/service/template_service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *template_config.ServerConfig, svc *template_service.TemplateService, loggerCfg *template_config.LoggerConfig) (*grpc.Server, error) {
	_logger, err := template_config.NewKratosLogger(loggerCfg, "grpc", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(_logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	template.RegisterTemplateServer(srv, svc)
	return srv, nil
}
