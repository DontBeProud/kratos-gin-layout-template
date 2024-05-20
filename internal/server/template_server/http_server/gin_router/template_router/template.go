package template_router

import (
	"github.com/gin-gonic/gin"
	"layout_template/api/template"
	"layout_template/internal/service/template_service"
)

// RegisterTemplateRouter 配置路由
// 路由配置应与proto中的配置保持一致
func RegisterTemplateRouter(engine *gin.Engine, svc *template_service.TemplateService) {

	groupTp := engine.Group("template")
	groupTp.GET("/:sid/query", func(c *gin.Context) {
		req := &template.QueryTemplateRequest{
			Sid:  c.Param("sid"),
			Name: c.DefaultQuery("name", ""),
		}
		resp, err := svc.QueryTemplate(c.Request.Context(), req)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, resp)
	})
	groupTp.POST("/:sid/create", func(c *gin.Context) {
		req := &template.CreateTemplateRequest{
			RealName: &template.CreateTemplateRequestBody{},
		}
		if err := c.BindJSON(req.RealName); err != nil {
			c.JSON(500, err.Error())
			return
		}
		req.Sid = c.Param("sid")
		resp, err := svc.CreateTemplate(c.Request.Context(), req)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, resp)
	})
}

func wrappedGinHandler() func(*gin.Context) {
	return func(c *gin.Context) {

	}
}
