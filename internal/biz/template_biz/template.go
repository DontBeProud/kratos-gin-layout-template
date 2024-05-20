package template_biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
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
	repo TemplateRepo
	log  *log.Helper
}

func NewTemplateUseCase(repo TemplateRepo, logger log.Logger) *TemplateUseCase {
	return &TemplateUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TemplateUseCase) CreateTemplate(ctx context.Context, g *Template) (*Template, error) {
	uc.log.WithContext(ctx).Infof("CreateTemplate Template: %v", g.Hello)
	return uc.repo.CreateTemplate(ctx, g)
}

func (uc *TemplateUseCase) QueryTemplate(ctx context.Context, g *Template) (*Template, error) {
	uc.log.WithContext(ctx).Infof("QueryTemplate Template: %v", g.Hello)
	return uc.repo.QueryTemplate(ctx, g)
}
