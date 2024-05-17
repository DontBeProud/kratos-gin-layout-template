package data

import (
	"github.com/google/wire"
	"layout_template/internal/data/base"
	"layout_template/internal/data/template_data"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(base.NewData, template_data.NewTemplateRepo)
