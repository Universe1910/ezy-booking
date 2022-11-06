package Booking

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Booking"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    BookingReq "github.com/flipped-aurora/gin-vue-admin/server/model/Booking/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EzyOrdersApi struct {
}

var ezyOrdersService = service.ServiceGroupApp.BookingServiceGroup.EzyOrdersService


// CreateEzyOrders Create EzyOrders
// @Tags EzyOrders
// @Summary Create EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyOrders true "Create EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyOrders/createEzyOrders [post]
func (ezyOrdersApi *EzyOrdersApi) CreateEzyOrders(c *gin.Context) {
	var ezyOrders Booking.EzyOrders
	err := c.ShouldBindJSON(&ezyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyOrdersService.CreateEzyOrders(ezyOrders); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteEzyOrders Delete EzyOrders
// @Tags EzyOrders
// @Summary Delete EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyOrders true "Delete EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyOrders/deleteEzyOrders [delete]
func (ezyOrdersApi *EzyOrdersApi) DeleteEzyOrders(c *gin.Context) {
	var ezyOrders Booking.EzyOrders
	err := c.ShouldBindJSON(&ezyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyOrdersService.DeleteEzyOrders(ezyOrders); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteEzyOrdersByIds Delete EzyOrders
// @Tags EzyOrders
// @Summary Delete EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyOrders/deleteEzyOrdersByIds [delete]
func (ezyOrdersApi *EzyOrdersApi) DeleteEzyOrdersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyOrdersService.DeleteEzyOrdersByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateEzyOrders 更新EzyOrders
// @Tags EzyOrders
// @Summary 更新EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyOrders true "更新EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyOrders/updateEzyOrders [put]
func (ezyOrdersApi *EzyOrdersApi) UpdateEzyOrders(c *gin.Context) {
	var ezyOrders Booking.EzyOrders
	err := c.ShouldBindJSON(&ezyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyOrdersService.UpdateEzyOrders(ezyOrders); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindEzyOrders 用id查询EzyOrders
// @Tags EzyOrders
// @Summary 用id查询EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Booking.EzyOrders true "用id查询EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyOrders/findEzyOrders [get]
func (ezyOrdersApi *EzyOrdersApi) FindEzyOrders(c *gin.Context) {
	var ezyOrders Booking.EzyOrders
	err := c.ShouldBindQuery(&ezyOrders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reezyOrders, err := ezyOrdersService.GetEzyOrders(ezyOrders.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"reezyOrders": reezyOrders}, c)
	}
}

// GetEzyOrdersList 分页获取EzyOrders列表
// @Tags EzyOrders
// @Summary 分页获取EzyOrders列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BookingReq.EzyOrdersSearch true "分页获取EzyOrders列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyOrders/getEzyOrdersList [get]
func (ezyOrdersApi *EzyOrdersApi) GetEzyOrdersList(c *gin.Context) {
	var pageInfo BookingReq.EzyOrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ezyOrdersService.GetEzyOrdersInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("Fail!", zap.Error(err))
        response.FailWithMessage("Fail", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "Successful", c)
    }
}
