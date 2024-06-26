package http_server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"layout_template/api/template"
	"layout_template/internal/conf/template_config"
	"layout_template/internal/service/template_service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *template_config.ServerConfig, svc *template_service.TemplateService, loggerCfg *template_config.LoggerConfig) (*http.Server, error) {
	_logger, err := template_config.NewKratosLogger(loggerCfg, "http", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(_logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// ---------------------------------------------------------
	// ---------------------------------------------------------
	// 使用grpc原生生成的路由配置，则使用对应的原生register函数
	template.RegisterTemplateHTTPServer(srv, svc)
	// ---------------------------------------------------------
	// ---------------------------------------------------------

	// ---------------------------------------------------------
	// ---------------------------------------------------------
	//// // 集成gin，可参照下列代码
	//g := gin.Default()
	//template_router.RegisterTemplateRouter(g, svc)
	//srv.HandlePrefix("/", g)
	// ---------------------------------------------------------
	// ---------------------------------------------------------

	return srv, nil
}
