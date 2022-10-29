package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EzyCustomerRouter struct {
}

// InitEzyCustomerRouter 初始化 EzyCustomer 路由信息
func (s *EzyCustomerRouter) InitEzyCustomerRouter(Router *gin.RouterGroup) {
	ezyCustomerRouter := Router.Group("ezyCustomer").Use(middleware.OperationRecord())
	ezyCustomerRouterWithoutRecord := Router.Group("ezyCustomer")
	var ezyCustomerApi = v1.ApiGroupApp.BookingApiGroup.EzyCustomerApi
	{
		ezyCustomerRouter.POST("createEzyCustomer", ezyCustomerApi.CreateEzyCustomer)   // 新建EzyCustomer
		ezyCustomerRouter.DELETE("deleteEzyCustomer", ezyCustomerApi.DeleteEzyCustomer) // Delete EzyCustomer
		ezyCustomerRouter.DELETE("deleteEzyCustomerByIds", ezyCustomerApi.DeleteEzyCustomerByIds) // 批量Delete EzyCustomer
		ezyCustomerRouter.PUT("updateEzyCustomer", ezyCustomerApi.UpdateEzyCustomer)    // 更新EzyCustomer
	}
	{
		ezyCustomerRouterWithoutRecord.GET("findEzyCustomer", ezyCustomerApi.FindEzyCustomer)        // 根据ID获取EzyCustomer
		ezyCustomerRouterWithoutRecord.GET("getEzyCustomerList", ezyCustomerApi.GetEzyCustomerList)  // 获取EzyCustomer列表
	}
}
