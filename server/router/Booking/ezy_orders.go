package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EzyOrdersRouter struct {
}

// InitEzyOrdersRouter 初始化 EzyOrders 路由信息
func (s *EzyOrdersRouter) InitEzyOrdersRouter(Router *gin.RouterGroup) {
	ezyOrdersRouter := Router.Group("ezyOrders").Use(middleware.OperationRecord())
	ezyOrdersRouterWithoutRecord := Router.Group("ezyOrders")
	var ezyOrdersApi = v1.ApiGroupApp.BookingApiGroup.EzyOrdersApi
	{
		ezyOrdersRouter.POST("createEzyOrders", ezyOrdersApi.CreateEzyOrders)   // 新建EzyOrders
		ezyOrdersRouter.DELETE("deleteEzyOrders", ezyOrdersApi.DeleteEzyOrders) // Delete EzyOrders
		ezyOrdersRouter.DELETE("deleteEzyOrdersByIds", ezyOrdersApi.DeleteEzyOrdersByIds) // 批量Delete EzyOrders
		ezyOrdersRouter.PUT("updateEzyOrders", ezyOrdersApi.UpdateEzyOrders)    // 更新EzyOrders
	}
	{
		ezyOrdersRouterWithoutRecord.GET("findEzyOrders", ezyOrdersApi.FindEzyOrders)        // 根据ID获取EzyOrders
		ezyOrdersRouterWithoutRecord.GET("getEzyOrdersList", ezyOrdersApi.GetEzyOrdersList)  // 获取EzyOrders列表
	}
}
