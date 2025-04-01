import service from '@/utils/request'
// @Tags CpuInject
// @Summary 创建CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CpuInject true "创建CPU故障注入"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /cpuInject/createCpuInject [post]
export const createCpuInject = (data) => {
  return service({
    url: '/cpuInject/createCpuInject',
    method: 'post',
    data
  })
}

// @Tags CpuInject
// @Summary 删除CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CpuInject true "删除CPU故障注入"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpuInject/deleteCpuInject [delete]
export const deleteCpuInject = (params) => {
  return service({
    url: '/cpuInject/deleteCpuInject',
    method: 'delete',
    params
  })
}

// @Tags CpuInject
// @Summary 批量删除CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CPU故障注入"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cpuInject/deleteCpuInject [delete]
export const deleteCpuInjectByIds = (params) => {
  return service({
    url: '/cpuInject/deleteCpuInjectByIds',
    method: 'delete',
    params
  })
}

// @Tags CpuInject
// @Summary 更新CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.CpuInject true "更新CPU故障注入"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cpuInject/updateCpuInject [put]
export const updateCpuInject = (data) => {
  return service({
    url: '/cpuInject/updateCpuInject',
    method: 'put',
    data
  })
}

// @Tags CpuInject
// @Summary 用id查询CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.CpuInject true "用id查询CPU故障注入"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cpuInject/findCpuInject [get]
export const findCpuInject = (params) => {
  return service({
    url: '/cpuInject/findCpuInject',
    method: 'get',
    params
  })
}

// @Tags CpuInject
// @Summary 分页获取CPU故障注入列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CPU故障注入列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cpuInject/getCpuInjectList [get]
export const getCpuInjectList = (params) => {
  return service({
    url: '/cpuInject/getCpuInjectList',
    method: 'get',
    params
  })
}

// @Tags CpuInject
// @Summary 不需要鉴权的CPU故障注入接口
// @Accept application/json
// @Produce application/json
// @Param data query faultInjectReq.CpuInjectSearch true "分页获取CPU故障注入列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cpuInject/getCpuInjectPublic [get]
export const getCpuInjectPublic = () => {
  return service({
    url: '/cpuInject/getCpuInjectPublic',
    method: 'get',
  })
}
// CpuInjectHandler CPU故障注入
// @Tags CpuInject
// @Summary CPU故障注入
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /cpuInject/cpuInjectHandler [GET]
export const cpuInjectHandler = () => {
  return service({
    url: '/cpuInject/cpuInjectHandler',
    method: 'GET'
  })
}
