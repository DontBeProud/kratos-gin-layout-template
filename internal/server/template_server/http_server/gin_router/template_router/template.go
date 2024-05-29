package template_router

import (
	"github.com/gin-gonic/gin"
	"layout_template/api/common/common_pb"
	"layout_template/api/template"
	"layout_template/internal/service/template_service"
	"strconv"
)

// RegisterTemplateRouter 配置路由
// 路由配置应与proto中的配置保持一致
func RegisterTemplateRouter(engine *gin.Engine, svc *template_service.TemplateService) {

	groupTp := engine.Group("template")
	groupTp.GET("/:base.gameId/query", func(c *gin.Context) {
		gId, _ := strconv.Atoi(c.Param("base.gameId"))
		req := &template.QueryTemplateRequest{
			Base: &common_pb.StandardRequestBase{
				GameId: uint32(gId),
			},
			Account: &common_pb.StandardAccount{
				AccountUId:  c.Query("account.accountUId"),
				AccountName: c.Query("account.accountName"),
			},
			Action: c.Query("action"),
		}
		resp, err := svc.QueryTemplate(c.Request.Context(), req)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		c.JSON(200, resp)
	})
	groupTp.POST("/:base.gameId/create", func(c *gin.Context) {
		gId, _ := strconv.Atoi(c.Param("base.gameId"))
		req := &template.CreateTemplateRequest{
			Base: &common_pb.StandardRequestBase{
				GameId: uint32(gId),
			},
			Body: &template.CreateTemplateRequestBody{},
		}
		if err := c.BindJSON(req.Body); err != nil {
			c.JSON(500, err.Error())
			return
		}
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
