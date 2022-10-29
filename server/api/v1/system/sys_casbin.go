package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// UpdateCasbin
// @Tags      Casbin
// @Summary   更新角色api权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CasbinInReceive        true  "权限id, 权限模型列表"
// @Success   200   {object}  response.Response{msg=string}  "更新角色api权限"
// @Router    /casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(cmr, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		global.GVA_LOG.Error("Update failure!", zap.Error(err))
		response.FailWithMessage("Update failure", c)
		return
	}
	response.OkWithMessage("update completed", c)
}

// GetPolicyPathByAuthorityId
// @Tags      Casbin
// @Summary   获取权限列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CasbinInReceive                                          true  "权限id, 权限模型列表"
// @Success   200   {object}  response.Response{data=systemRes.PolicyPathResponse,msg=string}  "获取权限列表,返回包括casbin详情列表"
// @Router    /casbin/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	err := c.ShouldBindJSON(&casbin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(casbin, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "Successful", c)
}
