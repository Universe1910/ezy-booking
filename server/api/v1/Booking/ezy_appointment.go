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
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type EzyAppointmentApi struct {
}

var ezyAppointmentService = service.ServiceGroupApp.BookingServiceGroup.EzyAppointmentService


// CreateEzyAppointment Create EzyAppointment
// @Tags EzyAppointment
// @Summary Create EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyAppointment true "Create EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyAppointment/createEzyAppointment [post]
func (ezyAppointmentApi *EzyAppointmentApi) CreateEzyAppointment(c *gin.Context) {
	var ezyAppointment Booking.EzyAppointment
	err := c.ShouldBindJSON(&ezyAppointment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "AppointmentName":{utils.NotEmpty()},
        "Singer":{utils.NotEmpty()},
        "AppointmentDate":{utils.NotEmpty()},
        "StartAt":{utils.NotEmpty()},
        "EndAt":{utils.NotEmpty()},
        "Stage":{utils.NotEmpty()},
        "StageMap":{utils.NotEmpty()},
        "StageArea":{utils.NotEmpty()},
        "Branch":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(ezyAppointment, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ezyAppointmentService.CreateEzyAppointment(ezyAppointment); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteEzyAppointment Delete EzyAppointment
// @Tags EzyAppointment
// @Summary Delete EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyAppointment true "Delete EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyAppointment/deleteEzyAppointment [delete]
func (ezyAppointmentApi *EzyAppointmentApi) DeleteEzyAppointment(c *gin.Context) {
	var ezyAppointment Booking.EzyAppointment
	err := c.ShouldBindJSON(&ezyAppointment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyAppointmentService.DeleteEzyAppointment(ezyAppointment); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteEzyAppointmentByIds Delete EzyAppointment
// @Tags EzyAppointment
// @Summary Delete EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyAppointment/deleteEzyAppointmentByIds [delete]
func (ezyAppointmentApi *EzyAppointmentApi) DeleteEzyAppointmentByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyAppointmentService.DeleteEzyAppointmentByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateEzyAppointment 更新EzyAppointment
// @Tags EzyAppointment
// @Summary 更新EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyAppointment true "更新EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyAppointment/updateEzyAppointment [put]
func (ezyAppointmentApi *EzyAppointmentApi) UpdateEzyAppointment(c *gin.Context) {
	var ezyAppointment Booking.EzyAppointment
	err := c.ShouldBindJSON(&ezyAppointment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "AppointmentName":{utils.NotEmpty()},
          "Singer":{utils.NotEmpty()},
          "AppointmentDate":{utils.NotEmpty()},
          "StartAt":{utils.NotEmpty()},
          "EndAt":{utils.NotEmpty()},
          "Stage":{utils.NotEmpty()},
          "StageMap":{utils.NotEmpty()},
          "StageArea":{utils.NotEmpty()},
          "Branch":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(ezyAppointment, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ezyAppointmentService.UpdateEzyAppointment(ezyAppointment); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindEzyAppointment 用id查询EzyAppointment
// @Tags EzyAppointment
// @Summary 用id查询EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Booking.EzyAppointment true "用id查询EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyAppointment/findEzyAppointment [get]
func (ezyAppointmentApi *EzyAppointmentApi) FindEzyAppointment(c *gin.Context) {
	var ezyAppointment Booking.EzyAppointment
	err := c.ShouldBindQuery(&ezyAppointment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reezyAppointment, err := ezyAppointmentService.GetEzyAppointment(ezyAppointment.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"reezyAppointment": reezyAppointment}, c)
	}
}

// GetEzyAppointmentList 分页获取EzyAppointment列表
// @Tags EzyAppointment
// @Summary 分页获取EzyAppointment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BookingReq.EzyAppointmentSearch true "分页获取EzyAppointment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyAppointment/getEzyAppointmentList [get]
func (ezyAppointmentApi *EzyAppointmentApi) GetEzyAppointmentList(c *gin.Context) {
	var pageInfo BookingReq.EzyAppointmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ezyAppointmentService.GetEzyAppointmentInfoList(pageInfo); err != nil {
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
