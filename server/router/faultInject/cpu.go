package faultInject

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CpuInjectRouter struct{}

func (s *CpuInjectRouter) InitCpuInjectRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	cpuInjectRouter := Router.Group("cpuInject").Use(middleware.OperationRecord())
	cpuInjectRouterWithoutRecord := Router.Group("cpuInject")
	cpuInjectRouterWithoutAuth := PublicRouter.Group("cpuInject")
	{
		cpuInjectRouter.POST("createCpuInject", cpuInjectApi.CreateCpuInject)
		cpuInjectRouter.DELETE("deleteCpuInject", cpuInjectApi.DeleteCpuInject)
		cpuInjectRouter.DELETE("deleteCpuInjectByIds", cpuInjectApi.DeleteCpuInjectByIds)
		cpuInjectRouter.PUT("updateCpuInject", cpuInjectApi.UpdateCpuInject)
	}
	{
		cpuInjectRouterWithoutRecord.GET("findCpuInject", cpuInjectApi.FindCpuInject)
		cpuInjectRouterWithoutRecord.GET("getCpuInjectList", cpuInjectApi.GetCpuInjectList)
	}
	{
		cpuInjectRouterWithoutAuth.GET("getCpuInjectPublic", cpuInjectApi.GetCpuInjectPublic)
		cpuInjectRouterWithoutAuth.GET("cpuInjectHandler", cpuInjectApi.CpuInjectHandler)
	}
}
