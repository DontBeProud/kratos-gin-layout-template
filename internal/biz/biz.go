package biz

import (
	"github.com/google/wire"
	"layout_template/internal/biz/template_biz"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(template_biz.NewTemplateUseCase)
