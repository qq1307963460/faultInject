
// 自动生成模板CpuInject
package faultInject
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CPU故障注入 结构体  CpuInject
type CpuInject struct {
    global.GVA_MODEL
    CoreCount  int `json:"coreCount" form:"coreCount" gorm:"default:4;column:coreCount;comment:;"`  //CPU核心数
    Load  int `json:"load" form:"load" gorm:"default:100;column:load;comment:;"`  //负载百分比
    Duration  int `json:"duration" form:"duration" gorm:"default:30;column:duration;comment:;"`  //持续时间
}


// TableName CPU故障注入 CpuInject自定义表名 cpu_inject
func (CpuInject) TableName() string {
    return "cpu_inject"
}





