package main

import (
	"flag"
	"layout_template/internal/conf/template_config"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

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
)

const (
	version            = "v1.1.20240523.1"
	defaultCfgFilePath = "../../configs/template_config_file/config.yaml"
	cfgFlagUsage       = "config path, eg: -conf config.yaml"
)

func init() {
	flag.StringVar(&flagConf, "conf", defaultCfgFilePath, cfgFlagUsage)
}

func newApp(loggerCfg *template_config.LoggerConfig, gs *grpc.Server, hs *http.Server) (*kratos.App, error) {
	_logger, err := template_config.NewKratosLogger(loggerCfg, "kratos", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(GitVersion+version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(_logger),
		kratos.Server(
			gs,
			hs,
		),
	), nil
}

func main() {
	flag.Parse()

	cfgSource := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	defer func() { _ = cfgSource.Close() }()

	if err := cfgSource.Load(); err != nil {
		panic(err)
	}

	var cfg template_config.Config
	if err := cfgSource.Scan(&cfg); err != nil {
		panic(err)
	}

	if cfg.LoggerCfg == nil {
		panic("cfg.LoggerCfg == nil")
	}

	app, cleanup, err := wireApp(cfg.ServerCfg, cfg.DataSourceCfg, cfg.LoggerCfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
