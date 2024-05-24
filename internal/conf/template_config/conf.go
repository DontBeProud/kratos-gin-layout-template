package template_config

import (
	"github.com/DontBeProud/go-kits/error_ex"
	"github.com/DontBeProud/go-kits/standard_logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

// NewStandardSystemLogger 根据配置创建系统日志对象
func NewStandardSystemLogger(loggerCfg *LoggerConfig, serviceName string, extraCallerSkip *uint,
	encoderCfg *zapcore.EncoderConfig, options []zap.Option) (*zap.Logger, error) {
	if loggerCfg == nil {
		return nil, error_ex.NewErrorExWithFuncNamePrefix(0, "loggerCfg == nil")
	}
	return standard_logger.NewStandardLoggerWithPb(loggerCfg.SystemLoggerCfg, serviceName, extraCallerSkip, encoderCfg, options)
}

// NewKratosLogger 根据配置创建kratos日志对象
func NewKratosLogger(loggerCfg *LoggerConfig, serviceName string, extraCallerSkip *uint,
	encoderCfg *zapcore.EncoderConfig, options []zap.Option) (*standard_logger.KratosLogger, error) {
	_l, err := NewStandardSystemLogger(loggerCfg, serviceName, extraCallerSkip, encoderCfg, options)
	if err != nil {
		return nil, err
	}
	return standard_logger.NewKratosLogger(_l)
}

func NewGormTracingLogger(loggerCfg *LoggerConfig, extraCallerSkip *uint,
	encoderCfg *zapcore.EncoderConfig, options []zap.Option) (logger.Interface, error_ex.ErrorEx) {
	_l, err := NewStandardSystemLogger(loggerCfg, "kratos", extraCallerSkip, encoderCfg, options)
	if err != nil {
		return nil, err
	}
	return standard_logger.NewGormTracingLoggerWithPb(_l, loggerCfg.MysqlTracingCfg)
}
