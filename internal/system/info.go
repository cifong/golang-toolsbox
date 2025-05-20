package system

import (
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// SystemInfo 定義系統資訊結構
type SystemInfo struct {
	CPUUsage    float64 `json:"cpu_usage"`
	TotalMemory uint64  `json:"total_memory"`
	UsedMemory  uint64  `json:"used_memory"`
	OS          string  `json:"os"`
	Arch        string  `json:"arch"`
	Version     string  `json:"version"`
}

// GetSystemInfo 取得 CPU 使用率與記憶體資訊
func GetSystemInfo() (*SystemInfo, error) {
	// 取得 CPU 使用率（1 秒平均）
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}

	// 取得記憶體資訊
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	info := &SystemInfo{
		CPUUsage:    cpuPercent[0],
		TotalMemory: vmStat.Total,
		UsedMemory:  vmStat.Used,
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Version:     runtime.Version(),
	}
	return info, nil
}
