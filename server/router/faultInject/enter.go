package faultInject

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ CpuInjectRouter }

var cpuInjectApi = api.ApiGroupApp.FaultInjectApiGroup.CpuInjectApi
