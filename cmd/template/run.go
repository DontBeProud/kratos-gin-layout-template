package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/DontBeProud/go-kits/standard_logger"
	pb "github.com/DontBeProud/go-kits/standard_logger/standard_logger_pb"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"go.uber.org/zap/zapcore"
	"layout_template/internal/conf/template_config"
	"time"
)

var (
	// 紧急情况下使用的日志对象。当配置文件出错，导致无法正确初始化日志对象时使用
	emergencyLogger, _ = standard_logger.NewStandardLoggerWithPb(&pb.LoggerConfig{}, "emergency", nil, nil, nil)
)

func init() {
	flag.StringVar(&flagConf, "conf", defaultCfgFilePath, cfgFlagUsage)
}

func run() {
	// 从配置文件中读取配置
	cfg := getConfig()

	// 初始化主线程的日志对象
	mainLogger, err := newMonitorLogger(cfg, "main")
	if err != nil {
		emergencyLogger.Panic(err.Error())
	}
	defer func() {
		if err != nil {
			mainLogger.GetLogger().Panic(err.Error())
		} else {
			mainLogger.GetLogger().Info("process exit")
		}
	}()
	mainLogger.GetLogger().Info(fmt.Sprintf("Process Run: %s", Version))
	mainLogger.GetLogger().Info(fmt.Sprintf("Load Config: %s", cfg.String()))

	// 开启进程状态监控
	ctx := context.Background()
	go func() { mainLogger.ExecuteGlobalMonitorTask(ctx, time.Minute, nil) }()

	// init app
	app, cleanup, err := wireApp(cfg.ServerCfg, cfg.DataSourceCfg, cfg.LoggerCfg)
	if err != nil {
		mainLogger.GetLogger().Panic(err.Error())
	}
	defer cleanup()

	// run
	if err = app.Run(); err != nil {
		mainLogger.GetLogger().Panic(err.Error())
	}
}

func newMonitorLogger(cfg *template_config.Config, serviceName string) (standard_logger.GlobalMonitorLogger, error) {
	sysCfg, err := template_config.GetSystemConfig(cfg.LoggerCfg)
	if err != nil {
		return nil, err
	}
	if sysCfg == nil {
		return nil, errors.New("sysCfg == nil")
	}
	rotation := 24 * time.Hour
	return standard_logger.NewMonitorLogger(&standard_logger.StandardLoggerConfig{
		RootDir:         sysCfg.RootDir,
		Level:           zapcore.InfoLevel,
		StackTraceLevel: zapcore.WarnLevel,
		DirName:         sysCfg.DirName,
		RotationTime:    &rotation,
		MaxAge:          nil,
	}, serviceName)
}

func getConfig() *template_config.Config {
	flag.Parse()

	cfgSource := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	defer func() { _ = cfgSource.Close() }()

	if err := cfgSource.Load(); err != nil {
		emergencyLogger.Panic(err.Error())
	}

	var cfg template_config.Config
	if err := cfgSource.Scan(&cfg); err != nil {
		emergencyLogger.Panic(err.Error())
	}

	return &cfg
}
