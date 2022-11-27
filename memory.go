package main

import (

    "github.com/shirou/gopsutil/v3/mem"
	"math"
)

type MemoryUsage struct {
	UsagePercent int
}

func getVirtualMemory() (MemoryUsage) {
    v, _ := mem.VirtualMemory()

	memoryUsagePercent := int(math.Round(v.UsedPercent))

	memoryInfo := MemoryUsage{memoryUsagePercent}

    return memoryInfo

}