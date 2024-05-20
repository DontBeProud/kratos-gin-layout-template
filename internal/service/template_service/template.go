package template_service

import (
	"context"
	"fmt"
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

func (s *TemplateService) QueryTemplate(ctx context.Context, in *template.QueryTemplateRequest) (*template.QueryTemplateReply, error) {
	g, err := s.uc.QueryTemplate(ctx, &template_biz.Template{Hello: in.Name, Sid: in.Sid})
	if err != nil {
		return nil, err
	}
	return &template.QueryTemplateReply{Message: fmt.Sprintf("Hello %s; your service id is %s", g.Hello, g.Sid)}, nil
}

func (s *TemplateService) CreateTemplate(ctx context.Context, in *template.CreateTemplateRequest) (*template.CreateTemplateReply, error) {
	g, err := s.uc.CreateTemplate(ctx, &template_biz.Template{Hello: in.RealName.Name, Sid: in.Sid})
	if err != nil {
		return nil, err
	}
	return &template.CreateTemplateReply{Message: fmt.Sprintf("Hello %s; your service id is %s", g.Hello, g.Sid)}, nil
}
