package faultInject

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/faultInject"
    faultInjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/faultInject/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CpuInjectApi struct {}



// CreateCpuInject 创建CPU故障注入
// @Tags CpuInject
// @Summary 创建CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body faultInject.CpuInject true "创建CPU故障注入"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /cpuInject/createCpuInject [post]
func (cpuInjectApi *CpuInjectApi) CreateCpuInject(c *gin.Context) {
	var cpuInject faultInject.CpuInject
	err := c.ShouldBindJSON(&cpuInject)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cpuInjectService.CreateCpuInject(&cpuInject)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCpuInject 删除CPU故障注入
// @Tags CpuInject
// @Summary 删除CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body faultInject.CpuInject true "删除CPU故障注入"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /cpuInject/deleteCpuInject [delete]
func (cpuInjectApi *CpuInjectApi) DeleteCpuInject(c *gin.Context) {
	ID := c.Query("ID")
	err := cpuInjectService.DeleteCpuInject(ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCpuInjectByIds 批量删除CPU故障注入
// @Tags CpuInject
// @Summary 批量删除CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /cpuInject/deleteCpuInjectByIds [delete]
func (cpuInjectApi *CpuInjectApi) DeleteCpuInjectByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := cpuInjectService.DeleteCpuInjectByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCpuInject 更新CPU故障注入
// @Tags CpuInject
// @Summary 更新CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body faultInject.CpuInject true "更新CPU故障注入"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /cpuInject/updateCpuInject [put]
func (cpuInjectApi *CpuInjectApi) UpdateCpuInject(c *gin.Context) {
	var cpuInject faultInject.CpuInject
	err := c.ShouldBindJSON(&cpuInject)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = cpuInjectService.UpdateCpuInject(cpuInject)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCpuInject 用id查询CPU故障注入
// @Tags CpuInject
// @Summary 用id查询CPU故障注入
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询CPU故障注入"
// @Success 200 {object} response.Response{data=faultInject.CpuInject,msg=string} "查询成功"
// @Router /cpuInject/findCpuInject [get]
func (cpuInjectApi *CpuInjectApi) FindCpuInject(c *gin.Context) {
	ID := c.Query("ID")
	recpuInject, err := cpuInjectService.GetCpuInject(ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(recpuInject, c)
}
// GetCpuInjectList 分页获取CPU故障注入列表
// @Tags CpuInject
// @Summary 分页获取CPU故障注入列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query faultInjectReq.CpuInjectSearch true "分页获取CPU故障注入列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /cpuInject/getCpuInjectList [get]
func (cpuInjectApi *CpuInjectApi) GetCpuInjectList(c *gin.Context) {
	var pageInfo faultInjectReq.CpuInjectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := cpuInjectService.GetCpuInjectInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetCpuInjectPublic 不需要鉴权的CPU故障注入接口
// @Tags CpuInject
// @Summary 不需要鉴权的CPU故障注入接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /cpuInject/getCpuInjectPublic [get]
func (cpuInjectApi *CpuInjectApi) GetCpuInjectPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    cpuInjectService.GetCpuInjectPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的CPU故障注入接口信息",
    }, "获取成功", c)
}
// CpuInjectHandler CPU故障注入
// @Tags CpuInject
// @Summary CPU故障注入
// @Accept application/json
// @Produce application/json
// @Param data query faultInjectReq.CpuInjectSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /cpuInject/cpuInjectHandler [GET]
func (cpuInjectApi *CpuInjectApi) CpuInjectHandler(c *gin.Context) {
    // 请添加自己的业务逻辑
    err := cpuInjectService.CpuInjectHandler(c)
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据", c)
}


