package template_service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"layout_template/api/common/common_pb"
	"layout_template/api/template"
	"layout_template/internal/biz/template_biz"
	"layout_template/internal/conf/template_config"
)

// TemplateService is a Template service.
type TemplateService struct {
	template.UnimplementedTemplateServer
	uc     *template_biz.TemplateUseCase
	logger *zap.Logger
}

// NewTemplateService new a Template service.
func NewTemplateService(uc *template_biz.TemplateUseCase, loggerCfg *template_config.LoggerConfig) (*TemplateService, error) {
	_logger, err := template_config.NewStandardSystemLogger(loggerCfg, "svc", nil, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TemplateService{uc: uc, logger: _logger}, nil
}

func (s *TemplateService) QueryTemplate(ctx context.Context, in *template.QueryTemplateRequest) (*template.QueryTemplateReply, error) {
	data := &template_biz.Template{
		Action: in.Action,
	}
	if in.Base != nil {
		data.GameId = in.Base.GameId
	}
	if in.Account != nil {
		data.AccountUid = in.Account.AccountUId
		data.AccountName = in.Account.AccountName
	}
	g, err := s.uc.QueryTemplate(ctx, data)
	if err != nil {
		return nil, err
	}
	return &template.QueryTemplateReply{
		Base: &common_pb.StandardReplyBase{
			Code: 1,
		},
		Message: fmt.Sprintf("game-id: %d, account: %s(%s), action: %s", g.GameId, g.AccountUid, g.AccountName, g.Action),
	}, nil
}

func (s *TemplateService) CreateTemplate(ctx context.Context, in *template.CreateTemplateRequest) (*template.CreateTemplateReply, error) {
	data := &template_biz.Template{}
	if in.Base != nil {
		data.GameId = in.Base.GameId
	}
	if in.Body != nil {
		data.Action = in.Body.Action
		if in.Body.Account != nil {
			data.AccountUid = in.Body.Account.AccountUId
			data.AccountName = in.Body.Account.AccountName
		}
	}
	g, err := s.uc.QueryTemplate(ctx, data)
	if err != nil {
		return nil, err
	}
	return &template.CreateTemplateReply{
		Base: &common_pb.StandardReplyBase{
			Code: 1,
		},
		Message: fmt.Sprintf("game-id: %d, account: %s(%s), action: %s", g.GameId, g.AccountUid, g.AccountName, g.Action),
	}, nil
}
