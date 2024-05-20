package template_server

import (
	"github.com/google/wire"
	"layout_template/internal/server/template_server/grpc_server"
	"layout_template/internal/server/template_server/http_server"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(grpc_server.NewGRPCServer, http_server.NewHTTPServer)
