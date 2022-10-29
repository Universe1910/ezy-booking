package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EzyBranchRouter struct {
}

// InitEzyBranchRouter 初始化 EzyBranch 路由信息
func (s *EzyBranchRouter) InitEzyBranchRouter(Router *gin.RouterGroup) {
	ezyBranchRouter := Router.Group("ezyBranch").Use(middleware.OperationRecord())
	ezyBranchRouterWithoutRecord := Router.Group("ezyBranch")
	var ezyBranchApi = v1.ApiGroupApp.BookingApiGroup.EzyBranchApi
	{
		ezyBranchRouter.POST("createEzyBranch", ezyBranchApi.CreateEzyBranch)   // 新建EzyBranch
		ezyBranchRouter.DELETE("deleteEzyBranch", ezyBranchApi.DeleteEzyBranch) // Delete EzyBranch
		ezyBranchRouter.DELETE("deleteEzyBranchByIds", ezyBranchApi.DeleteEzyBranchByIds) // 批量Delete EzyBranch
		ezyBranchRouter.PUT("updateEzyBranch", ezyBranchApi.UpdateEzyBranch)    // 更新EzyBranch
	}
	{
		ezyBranchRouterWithoutRecord.GET("findEzyBranch", ezyBranchApi.FindEzyBranch)        // 根据ID获取EzyBranch
		ezyBranchRouterWithoutRecord.GET("getEzyBranchList", ezyBranchApi.GetEzyBranchList)  // 获取EzyBranch列表
	}
}
