import service from '@/utils/request'

// @Tags EzyOrders
// @Summary Create EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyOrders true "Create EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyOrders/createEzyOrders [post]
export const createEzyOrders = (data) => {
  return service({
    url: '/ezyOrders/createEzyOrders',
    method: 'post',
    data
  })
}

// @Tags EzyOrders
// @Summary Delete EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyOrders true "Delete EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyOrders/deleteEzyOrders [delete]
export const deleteEzyOrders = (data) => {
  return service({
    url: '/ezyOrders/deleteEzyOrders',
    method: 'delete',
    data
  })
}

// @Tags EzyOrders
// @Summary Delete EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyOrders/deleteEzyOrders [delete]
export const deleteEzyOrdersByIds = (data) => {
  return service({
    url: '/ezyOrders/deleteEzyOrdersByIds',
    method: 'delete',
    data
  })
}

// @Tags EzyOrders
// @Summary 更新EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyOrders true "更新EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyOrders/updateEzyOrders [put]
export const updateEzyOrders = (data) => {
  return service({
    url: '/ezyOrders/updateEzyOrders',
    method: 'put',
    data
  })
}

// @Tags EzyOrders
// @Summary 用id查询EzyOrders
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EzyOrders true "用id查询EzyOrders"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyOrders/findEzyOrders [get]
export const findEzyOrders = (params) => {
  return service({
    url: '/ezyOrders/findEzyOrders',
    method: 'get',
    params
  })
}

// @Tags EzyOrders
// @Summary 分页获取EzyOrders列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EzyOrders列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyOrders/getEzyOrdersList [get]
export const getEzyOrdersList = (params) => {
  return service({
    url: '/ezyOrders/getEzyOrdersList',
    method: 'get',
    params
  })
}
