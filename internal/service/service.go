package service

import (
	"github.com/google/wire"
	"layout_template/internal/service/template_service"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(template_service.NewTemplateService)
