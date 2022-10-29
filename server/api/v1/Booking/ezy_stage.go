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

type EzyStageApi struct {
}

var ezyStageService = service.ServiceGroupApp.BookingServiceGroup.EzyStageService


// CreateEzyStage Create EzyStage
// @Tags EzyStage
// @Summary Create EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyStage true "Create EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyStage/createEzyStage [post]
func (ezyStageApi *EzyStageApi) CreateEzyStage(c *gin.Context) {
	var ezyStage Booking.EzyStage
	err := c.ShouldBindJSON(&ezyStage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Address":{utils.NotEmpty()},
    }
	if err := utils.Verify(ezyStage, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ezyStageService.CreateEzyStage(ezyStage); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteEzyStage Delete EzyStage
// @Tags EzyStage
// @Summary Delete EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyStage true "Delete EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyStage/deleteEzyStage [delete]
func (ezyStageApi *EzyStageApi) DeleteEzyStage(c *gin.Context) {
	var ezyStage Booking.EzyStage
	err := c.ShouldBindJSON(&ezyStage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyStageService.DeleteEzyStage(ezyStage); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteEzyStageByIds Delete EzyStage
// @Tags EzyStage
// @Summary Delete EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyStage/deleteEzyStageByIds [delete]
func (ezyStageApi *EzyStageApi) DeleteEzyStageByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyStageService.DeleteEzyStageByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateEzyStage 更新EzyStage
// @Tags EzyStage
// @Summary 更新EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyStage true "更新EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyStage/updateEzyStage [put]
func (ezyStageApi *EzyStageApi) UpdateEzyStage(c *gin.Context) {
	var ezyStage Booking.EzyStage
	err := c.ShouldBindJSON(&ezyStage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Address":{utils.NotEmpty()},
      }
    if err := utils.Verify(ezyStage, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ezyStageService.UpdateEzyStage(ezyStage); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindEzyStage 用id查询EzyStage
// @Tags EzyStage
// @Summary 用id查询EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Booking.EzyStage true "用id查询EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyStage/findEzyStage [get]
func (ezyStageApi *EzyStageApi) FindEzyStage(c *gin.Context) {
	var ezyStage Booking.EzyStage
	err := c.ShouldBindQuery(&ezyStage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reezyStage, err := ezyStageService.GetEzyStage(ezyStage.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"reezyStage": reezyStage}, c)
	}
}

// GetEzyStageList 分页获取EzyStage列表
// @Tags EzyStage
// @Summary 分页获取EzyStage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BookingReq.EzyStageSearch true "分页获取EzyStage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyStage/getEzyStageList [get]
func (ezyStageApi *EzyStageApi) GetEzyStageList(c *gin.Context) {
	var pageInfo BookingReq.EzyStageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ezyStageService.GetEzyStageInfoList(pageInfo); err != nil {
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
