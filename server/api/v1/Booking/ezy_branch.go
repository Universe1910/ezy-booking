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

type EzyBranchApi struct {
}

var ezyBranchService = service.ServiceGroupApp.BookingServiceGroup.EzyBranchService


// CreateEzyBranch Create EzyBranch
// @Tags EzyBranch
// @Summary Create EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyBranch true "Create EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyBranch/createEzyBranch [post]
func (ezyBranchApi *EzyBranchApi) CreateEzyBranch(c *gin.Context) {
	var ezyBranch Booking.EzyBranch
	err := c.ShouldBindJSON(&ezyBranch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
    }
	if err := utils.Verify(ezyBranch, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ezyBranchService.CreateEzyBranch(ezyBranch); err != nil {
        global.GVA_LOG.Error("Failed to create!", zap.Error(err))
		response.FailWithMessage("Failed to create", c)
	} else {
		response.OkWithMessage("Successful creation", c)
	}
}

// DeleteEzyBranch Delete EzyBranch
// @Tags EzyBranch
// @Summary Delete EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyBranch true "Delete EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyBranch/deleteEzyBranch [delete]
func (ezyBranchApi *EzyBranchApi) DeleteEzyBranch(c *gin.Context) {
	var ezyBranch Booking.EzyBranch
	err := c.ShouldBindJSON(&ezyBranch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyBranchService.DeleteEzyBranch(ezyBranch); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// DeleteEzyBranchByIds Delete EzyBranch
// @Tags EzyBranch
// @Summary Delete EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "Delete EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyBranch/deleteEzyBranchByIds [delete]
func (ezyBranchApi *EzyBranchApi) DeleteEzyBranchByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ezyBranchService.DeleteEzyBranchByIds(IDS); err != nil {
        global.GVA_LOG.Error("failed to delete!", zap.Error(err))
		response.FailWithMessage("failed to delete", c)
	} else {
		response.OkWithMessage("successfully deleted", c)
	}
}

// UpdateEzyBranch 更新EzyBranch
// @Tags EzyBranch
// @Summary 更新EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Booking.EzyBranch true "更新EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyBranch/updateEzyBranch [put]
func (ezyBranchApi *EzyBranchApi) UpdateEzyBranch(c *gin.Context) {
	var ezyBranch Booking.EzyBranch
	err := c.ShouldBindJSON(&ezyBranch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
      }
    if err := utils.Verify(ezyBranch, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ezyBranchService.UpdateEzyBranch(ezyBranch); err != nil {
        global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
	} else {
		response.OkWithMessage("update completed", c)
	}
}

// FindEzyBranch 用id查询EzyBranch
// @Tags EzyBranch
// @Summary 用id查询EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Booking.EzyBranch true "用id查询EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyBranch/findEzyBranch [get]
func (ezyBranchApi *EzyBranchApi) FindEzyBranch(c *gin.Context) {
	var ezyBranch Booking.EzyBranch
	err := c.ShouldBindQuery(&ezyBranch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reezyBranch, err := ezyBranchService.GetEzyBranch(ezyBranch.ID); err != nil {
        global.GVA_LOG.Error("Query Failed", zap.Error(err))
		response.FailWithMessage("Query Failed", c)
	} else {
		response.OkWithData(gin.H{"reezyBranch": reezyBranch}, c)
	}
}

// GetEzyBranchList 分页获取EzyBranch列表
// @Tags EzyBranch
// @Summary 分页获取EzyBranch列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query BookingReq.EzyBranchSearch true "分页获取EzyBranch列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyBranch/getEzyBranchList [get]
func (ezyBranchApi *EzyBranchApi) GetEzyBranchList(c *gin.Context) {
	var pageInfo BookingReq.EzyBranchSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ezyBranchService.GetEzyBranchInfoList(pageInfo); err != nil {
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
