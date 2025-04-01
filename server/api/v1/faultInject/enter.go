package faultInject

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ CpuInjectApi }

var cpuInjectService = service.ServiceGroupApp.FaultInjectServiceGroup.CpuInjectService
