package template_data

import (
	"context"
	"layout_template/internal/biz/template_biz"
)

type templateRepo struct {
	data *Data
}

// NewTemplateRepo .
func NewTemplateRepo(data *Data) template_biz.TemplateRepo {

	return &templateRepo{data: data}
}

func (r *templateRepo) CreateTemplate(ctx context.Context, g *template_biz.Template) (*template_biz.Template, error) {
	return g, nil
}

func (r *templateRepo) QueryTemplate(ctx context.Context, g *template_biz.Template) (*template_biz.Template, error) {
	return g, nil
}
