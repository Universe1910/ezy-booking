import service from '@/utils/request'

// @Tags EzyBranch
// @Summary Create EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyBranch true "Create EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyBranch/createEzyBranch [post]
export const createEzyBranch = (data) => {
  return service({
    url: '/ezyBranch/createEzyBranch',
    method: 'post',
    data
  })
}

// @Tags EzyBranch
// @Summary Delete EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyBranch true "Delete EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyBranch/deleteEzyBranch [delete]
export const deleteEzyBranch = (data) => {
  return service({
    url: '/ezyBranch/deleteEzyBranch',
    method: 'delete',
    data
  })
}

// @Tags EzyBranch
// @Summary Delete EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyBranch/deleteEzyBranch [delete]
export const deleteEzyBranchByIds = (data) => {
  return service({
    url: '/ezyBranch/deleteEzyBranchByIds',
    method: 'delete',
    data
  })
}

// @Tags EzyBranch
// @Summary 更新EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyBranch true "更新EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyBranch/updateEzyBranch [put]
export const updateEzyBranch = (data) => {
  return service({
    url: '/ezyBranch/updateEzyBranch',
    method: 'put',
    data
  })
}

// @Tags EzyBranch
// @Summary 用id查询EzyBranch
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EzyBranch true "用id查询EzyBranch"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyBranch/findEzyBranch [get]
export const findEzyBranch = (params) => {
  return service({
    url: '/ezyBranch/findEzyBranch',
    method: 'get',
    params
  })
}

// @Tags EzyBranch
// @Summary 分页获取EzyBranch列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EzyBranch列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyBranch/getEzyBranchList [get]
export const getEzyBranchList = (params) => {
  return service({
    url: '/ezyBranch/getEzyBranchList',
    method: 'get',
    params
  })
}
