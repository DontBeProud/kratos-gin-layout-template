package template_data

import (
	"go.uber.org/zap"
	"layout_template/internal/conf/template_config"
)

// Data .
type Data struct {
	logger *zap.Logger
}

// NewData .
func NewData(c *template_config.DataSourceCfg, loggerCfg *template_config.LoggerConfig) (*Data, func(), error) {
	_logger, err := template_config.NewStandardSystemLogger(loggerCfg, "data", nil, nil, nil)
	cleanup := func() {
		if _logger != nil {
			_logger.Info("closing the data resources")
		}
	}
	if err != nil {
		return nil, cleanup, err
	}

	return &Data{}, cleanup, nil
}
