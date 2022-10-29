import service from '@/utils/request'

// @Tags EzyCustomer
// @Summary Create EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyCustomer true "Create EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyCustomer/createEzyCustomer [post]
export const createEzyCustomer = (data) => {
  return service({
    url: '/ezyCustomer/createEzyCustomer',
    method: 'post',
    data
  })
}

// @Tags EzyCustomer
// @Summary Delete EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyCustomer true "Delete EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyCustomer/deleteEzyCustomer [delete]
export const deleteEzyCustomer = (data) => {
  return service({
    url: '/ezyCustomer/deleteEzyCustomer',
    method: 'delete',
    data
  })
}

// @Tags EzyCustomer
// @Summary Delete EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyCustomer/deleteEzyCustomer [delete]
export const deleteEzyCustomerByIds = (data) => {
  return service({
    url: '/ezyCustomer/deleteEzyCustomerByIds',
    method: 'delete',
    data
  })
}

// @Tags EzyCustomer
// @Summary 更新EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyCustomer true "更新EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyCustomer/updateEzyCustomer [put]
export const updateEzyCustomer = (data) => {
  return service({
    url: '/ezyCustomer/updateEzyCustomer',
    method: 'put',
    data
  })
}

// @Tags EzyCustomer
// @Summary 用id查询EzyCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EzyCustomer true "用id查询EzyCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyCustomer/findEzyCustomer [get]
export const findEzyCustomer = (params) => {
  return service({
    url: '/ezyCustomer/findEzyCustomer',
    method: 'get',
    params
  })
}

// @Tags EzyCustomer
// @Summary 分页获取EzyCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EzyCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyCustomer/getEzyCustomerList [get]
export const getEzyCustomerList = (params) => {
  return service({
    url: '/ezyCustomer/getEzyCustomerList',
    method: 'get',
    params
  })
}
