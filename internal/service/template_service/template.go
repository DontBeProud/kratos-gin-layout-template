package template_service

import (
	"context"
	"layout_template/api/template"
	"layout_template/internal/biz/template_biz"
)

// TemplateService is a Template service.
type TemplateService struct {
	template.UnimplementedTemplateServer
	uc *template_biz.TemplateUseCase
}

// NewTemplateService new a Template service.
func NewTemplateService(uc *template_biz.TemplateUseCase) *TemplateService {
	return &TemplateService{uc: uc}
}

func (s *TemplateService) GetTemplate(ctx context.Context, in *template.GetTemplateRequest) (*template.GetTemplateReply, error) {
	g, err := s.uc.GetTemplate(ctx, &template_biz.Template{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &template.GetTemplateReply{Message: "Hello " + g.Hello}, nil
}

func (s *TemplateService) CreateTemplate(ctx context.Context, in *template.CreateTemplateRequest) (*template.CreateTemplateReply, error) {
	g, err := s.uc.CreateTemplate(ctx, &template_biz.Template{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &template.CreateTemplateReply{Message: "Hello " + g.Hello}, nil
}
