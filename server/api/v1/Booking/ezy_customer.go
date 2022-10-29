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

type EzyCustomerApi struct {
}

var ezyCustomerService = service.ServiceGroupApp.BookingServiceGroup.EzyCustomerService


// CreateEzyCustomer Create EzyCustomer
// @Tags EzyCustomer
// @Summary Create EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyCustomer true "Create EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyCustomer/createEzyCustomer [post]
func (ezyCustomerApi *EzyCustomerApi) CreateEzyCustomer(c *gin.Context) {
	var ezyCustomer Booking.EzyCustomer
	err := c.ShouldBindJSON(&ezyCustomer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Phone":{utils.NotEmpty()},
        "Email":{utils.NotEmpty()},
    }
	if err := utils.Verify(ezyCustomer, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ezyCustomerService.CreateEzyCustomer(ezyCustomer); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteEzyCustomer Delete EzyCustomer
// @Tags EzyCustomer
// @Summary Delete EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyCustomer true "Delete EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyCustomer/deleteEzyCustomer [delete]
func (ezyCustomerApi *EzyCustomerApi) DeleteEzyCustomer(c *gin.Context) {
	var ezyCustomer Booking.EzyCustomer
	err := c.ShouldBindJSON(&ezyCustomer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyCustomerService.DeleteEzyCustomer(ezyCustomer); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteEzyCustomerByIds Delete EzyCustomer
// @Tags EzyCustomer
// @Summary Delete EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyCustomer/deleteEzyCustomerByIds [delete]
func (ezyCustomerApi *EzyCustomerApi) DeleteEzyCustomerByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyCustomerService.DeleteEzyCustomerByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateEzyCustomer 更新EzyCustomer
// @Tags EzyCustomer
// @Summary 更新EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyCustomer true "更新EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyCustomer/updateEzyCustomer [put]
func (ezyCustomerApi *EzyCustomerApi) UpdateEzyCustomer(c *gin.Context) {
	var ezyCustomer Booking.EzyCustomer
	err := c.ShouldBindJSON(&ezyCustomer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Phone":{utils.NotEmpty()},
          "Email":{utils.NotEmpty()},
      }
    if err := utils.Verify(ezyCustomer, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ezyCustomerService.UpdateEzyCustomer(ezyCustomer); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindEzyCustomer 用id查询EzyCustomer
// @Tags EzyCustomer
// @Summary 用id查询EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Booking.EzyCustomer true "用id查询EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyCustomer/findEzyCustomer [get]
func (ezyCustomerApi *EzyCustomerApi) FindEzyCustomer(c *gin.Context) {
	var ezyCustomer Booking.EzyCustomer
	err := c.ShouldBindQuery(&ezyCustomer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reezyCustomer, err := ezyCustomerService.GetEzyCustomer(ezyCustomer.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"reezyCustomer": reezyCustomer}, c)
	}
}

// GetEzyCustomerList 分页获取EzyCustomer列表
// @Tags EzyCustomer
// @Summary 分页获取EzyCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BookingReq.EzyCustomerSearch true "分页获取EzyCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyCustomer/getEzyCustomerList [get]
func (ezyCustomerApi *EzyCustomerApi) GetEzyCustomerList(c *gin.Context) {
	var pageInfo BookingReq.EzyCustomerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ezyCustomerService.GetEzyCustomerInfoList(pageInfo); err != nil {
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
