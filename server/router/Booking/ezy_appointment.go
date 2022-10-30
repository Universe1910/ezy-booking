package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EzyAppointmentRouter struct {
}

// InitEzyAppointmentRouter 初始化 EzyAppointment 路由信息
func (s *EzyAppointmentRouter) InitEzyAppointmentRouter(Router *gin.RouterGroup) {
	ezyAppointmentRouter := Router.Group("ezyAppointment").Use(middleware.OperationRecord())
	ezyAppointmentRouterWithoutRecord := Router.Group("ezyAppointment")
	var ezyAppointmentApi = v1.ApiGroupApp.BookingApiGroup.EzyAppointmentApi
	{
		ezyAppointmentRouter.POST("createEzyAppointment", ezyAppointmentApi.CreateEzyAppointment)   // 新建EzyAppointment
		ezyAppointmentRouter.DELETE("deleteEzyAppointment", ezyAppointmentApi.DeleteEzyAppointment) // Delete EzyAppointment
		ezyAppointmentRouter.DELETE("deleteEzyAppointmentByIds", ezyAppointmentApi.DeleteEzyAppointmentByIds) // 批量Delete EzyAppointment
		ezyAppointmentRouter.PUT("updateEzyAppointment", ezyAppointmentApi.UpdateEzyAppointment)    // 更新EzyAppointment
	}
	{
		ezyAppointmentRouterWithoutRecord.GET("findEzyAppointment", ezyAppointmentApi.FindEzyAppointment)        // 根据ID获取EzyAppointment
		ezyAppointmentRouterWithoutRecord.GET("getEzyAppointmentList", ezyAppointmentApi.GetEzyAppointmentList)  // 获取EzyAppointment列表
	}
}
