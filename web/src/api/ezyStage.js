import service from '@/utils/request'

// @Tags EzyStage
// @Summary Create EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyStage true "Create EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyStage/createEzyStage [post]
export const createEzyStage = (data) => {
  return service({
    url: '/ezyStage/createEzyStage',
    method: 'post',
    data
  })
}

// @Tags EzyStage
// @Summary Delete EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyStage true "Delete EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyStage/deleteEzyStage [delete]
export const deleteEzyStage = (data) => {
  return service({
    url: '/ezyStage/deleteEzyStage',
    method: 'delete',
    data
  })
}

// @Tags EzyStage
// @Summary Delete EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyStage/deleteEzyStage [delete]
export const deleteEzyStageByIds = (data) => {
  return service({
    url: '/ezyStage/deleteEzyStageByIds',
    method: 'delete',
    data
  })
}

// @Tags EzyStage
// @Summary 更新EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyStage true "更新EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyStage/updateEzyStage [put]
export const updateEzyStage = (data) => {
  return service({
    url: '/ezyStage/updateEzyStage',
    method: 'put',
    data
  })
}

// @Tags EzyStage
// @Summary 用id查询EzyStage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EzyStage true "用id查询EzyStage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyStage/findEzyStage [get]
export const findEzyStage = (params) => {
  return service({
    url: '/ezyStage/findEzyStage',
    method: 'get',
    params
  })
}

// @Tags EzyStage
// @Summary 分页获取EzyStage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EzyStage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyStage/getEzyStageList [get]
export const getEzyStageList = (params) => {
  return service({
    url: '/ezyStage/getEzyStageList',
    method: 'get',
    params
  })
}
