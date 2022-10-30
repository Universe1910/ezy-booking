import service from '@/utils/request'

// @Tags EzyAppointment
// @Summary Create EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyAppointment true "Create EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyAppointment/createEzyAppointment [post]
export const createEzyAppointment = (data) => {
  return service({
    url: '/ezyAppointment/createEzyAppointment',
    method: 'post',
    data
  })
}

// @Tags EzyAppointment
// @Summary Delete EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyAppointment true "Delete EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyAppointment/deleteEzyAppointment [delete]
export const deleteEzyAppointment = (data) => {
  return service({
    url: '/ezyAppointment/deleteEzyAppointment',
    method: 'delete',
    data
  })
}

// @Tags EzyAppointment
// @Summary Delete EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量Delete EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"successfully deleted"}"
// @Router /ezyAppointment/deleteEzyAppointment [delete]
export const deleteEzyAppointmentByIds = (data) => {
  return service({
    url: '/ezyAppointment/deleteEzyAppointmentByIds',
    method: 'delete',
    data
  })
}

// @Tags EzyAppointment
// @Summary 更新EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EzyAppointment true "更新EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"update completed"}"
// @Router /ezyAppointment/updateEzyAppointment [put]
export const updateEzyAppointment = (data) => {
  return service({
    url: '/ezyAppointment/updateEzyAppointment',
    method: 'put',
    data
  })
}

// @Tags EzyAppointment
// @Summary 用id查询EzyAppointment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EzyAppointment true "用id查询EzyAppointment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Find Successfully"}"
// @Router /ezyAppointment/findEzyAppointment [get]
export const findEzyAppointment = (params) => {
  return service({
    url: '/ezyAppointment/findEzyAppointment',
    method: 'get',
    params
  })
}

// @Tags EzyAppointment
// @Summary 分页获取EzyAppointment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EzyAppointment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Successful"}"
// @Router /ezyAppointment/getEzyAppointmentList [get]
export const getEzyAppointmentList = (params) => {
  return service({
    url: '/ezyAppointment/getEzyAppointmentList',
    method: 'get',
    params
  })
}
