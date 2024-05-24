package template_biz

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"layout_template/internal/conf/template_config"
)

type Template struct {
	Hello string
	Sid   string
}

type TemplateRepo interface {
	CreateTemplate(context.Context, *Template) (*Template, error)
	QueryTemplate(context.Context, *Template) (*Template, error)
}

type TemplateUseCase struct {
	repo   TemplateRepo
	logger *zap.Logger
}

func NewTemplateUseCase(repo TemplateRepo, loggerCfg *template_config.LoggerConfig) (*TemplateUseCase, error) {
	_logger, err := template_config.NewStandardSystemLogger(loggerCfg, "biz", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TemplateUseCase{repo: repo, logger: _logger}, nil
}

func (uc *TemplateUseCase) CreateTemplate(ctx context.Context, g *Template) (*Template, error) {
	uc.logger.Info(fmt.Sprintf("CreateTemplate Template: %v", g.Hello))
	return uc.repo.CreateTemplate(ctx, g)
}

func (uc *TemplateUseCase) QueryTemplate(ctx context.Context, g *Template) (*Template, error) {
	uc.logger.Info(fmt.Sprintf("QueryTemplate Template: %v", g.Hello))
	return uc.repo.QueryTemplate(ctx, g)
}
