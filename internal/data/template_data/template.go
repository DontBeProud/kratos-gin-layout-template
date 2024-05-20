package template_data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"layout_template/internal/biz/template_biz"
)

type templateRepo struct {
	data *Data
	log  *log.Helper
}

// NewTemplateRepo .
func NewTemplateRepo(data *Data, logger log.Logger) template_biz.TemplateRepo {
	return &templateRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *templateRepo) CreateTemplate(ctx context.Context, g *template_biz.Template) (*template_biz.Template, error) {
	return g, nil
}

func (r *templateRepo) QueryTemplate(ctx context.Context, g *template_biz.Template) (*template_biz.Template, error) {
	return g, nil
}
