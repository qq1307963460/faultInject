
package faultInject

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/faultInject"
    faultInjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/faultInject/request"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"runtime"
	"time"
)

type CpuInjectService struct {}
// CreateCpuInject 创建CPU故障注入记录
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService) CreateCpuInject(cpuInject *faultInject.CpuInject) (err error) {
	err = global.GVA_DB.Create(cpuInject).Error
	return err
}

// DeleteCpuInject 删除CPU故障注入记录
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService)DeleteCpuInject(ID string) (err error) {
	err = global.GVA_DB.Delete(&faultInject.CpuInject{},"id = ?",ID).Error
	return err
}

// DeleteCpuInjectByIds 批量删除CPU故障注入记录
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService)DeleteCpuInjectByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]faultInject.CpuInject{},"id in ?",IDs).Error
	return err
}

// UpdateCpuInject 更新CPU故障注入记录
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService)UpdateCpuInject(cpuInject faultInject.CpuInject) (err error) {
	err = global.GVA_DB.Model(&faultInject.CpuInject{}).Where("id = ?",cpuInject.ID).Updates(&cpuInject).Error
	return err
}

// GetCpuInject 根据ID获取CPU故障注入记录
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService)GetCpuInject(ID string) (cpuInject faultInject.CpuInject, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cpuInject).Error
	return
}
// GetCpuInjectInfoList 分页获取CPU故障注入记录
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService)GetCpuInjectInfoList(info faultInjectReq.CpuInjectSearch) (list []faultInject.CpuInject, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&faultInject.CpuInject{})
    var cpuInjects []faultInject.CpuInject
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&cpuInjects).Error
	return  cpuInjects, total, err
}
func (cpuInjectService *CpuInjectService)GetCpuInjectPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}

// CpuInjectHandler CPU故障注入
// Author [yourname](https://github.com/yourname)
func (cpuInjectService *CpuInjectService) CpuInjectHandler(c *gin.Context) (err error) {
	var req faultInject.CpuInject

	// 绑定JSON参数
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败"})
		return
	}

	// 参数校验
	if req.CoreCount <= 0 || req.Load < 0 || req.Load > 100 || req.Duration <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}

	// 启动 CPU 压测 goroutine
	go simulateCPULoad(req.CoreCount, req.Load, req.Duration)

	// 写入数据库记录
	req.CreatedAt = time.Now()
	if err = global.GVA_DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库插入失败"})
		return
	}

	return
}

func simulateCPULoad(cores int, load int, duration int) {
	runtime.GOMAXPROCS(cores)
	for i := 0; i < cores; i++ {
		go func() {
			busyTime := time.Duration(load) * time.Millisecond
			idleTime := time.Duration(100-load) * time.Millisecond
			stopTime := time.Now().Add(time.Duration(duration) * time.Second)

			for time.Now().Before(stopTime) {
				start := time.Now()
				for time.Since(start) < busyTime {
					_ = math.Sqrt(123.456)
				}
				time.Sleep(idleTime)
			}
		}()
	}
}