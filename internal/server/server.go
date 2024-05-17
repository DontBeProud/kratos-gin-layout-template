package server

import (
	"github.com/google/wire"
	"layout_template/internal/server/grpc_server"
	"layout_template/internal/server/http_server"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(grpc_server.NewGRPCServer, http_server.NewHTTPServer)
