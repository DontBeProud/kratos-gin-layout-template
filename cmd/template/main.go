package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"layout_template/internal/conf/template_config"
	"os"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.GitVersion=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// GitVersion is the version of the compiled software.
	GitVersion string
	// flagConf is the config flag.
	flagConf string

	id, _ = os.Hostname()

	Version = GitVersion + version
)

const (
	version            = "v1.1.20240523.1"
	defaultCfgFilePath = "../../configs/template_config_file/config.yaml"
	cfgFlagUsage       = "config path, eg: -conf config.yaml"
)

func main() {
	run()
}

func newApp(loggerCfg *template_config.LoggerConfig, gs *grpc.Server, hs *http.Server) (*kratos.App, error) {
	_logger, err := template_config.NewKratosLogger(loggerCfg, "kratos", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(_logger),
		kratos.Server(
			gs,
			hs,
		),
	), nil
}
