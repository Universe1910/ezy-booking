package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EzyStageRouter struct {
}

// InitEzyStageRouter 初始化 EzyStage 路由信息
func (s *EzyStageRouter) InitEzyStageRouter(Router *gin.RouterGroup) {
	ezyStageRouter := Router.Group("ezyStage").Use(middleware.OperationRecord())
	ezyStageRouterWithoutRecord := Router.Group("ezyStage")
	var ezyStageApi = v1.ApiGroupApp.BookingApiGroup.EzyStageApi
	{
		ezyStageRouter.POST("createEzyStage", ezyStageApi.CreateEzyStage)   // 新建EzyStage
		ezyStageRouter.DELETE("deleteEzyStage", ezyStageApi.DeleteEzyStage) // Delete EzyStage
		ezyStageRouter.DELETE("deleteEzyStageByIds", ezyStageApi.DeleteEzyStageByIds) // 批量Delete EzyStage
		ezyStageRouter.PUT("updateEzyStage", ezyStageApi.UpdateEzyStage)    // 更新EzyStage
	}
	{
		ezyStageRouterWithoutRecord.GET("findEzyStage", ezyStageApi.FindEzyStage)        // 根据ID获取EzyStage
		ezyStageRouterWithoutRecord.GET("getEzyStageList", ezyStageApi.GetEzyStageList)  // 获取EzyStage列表
	}
}
