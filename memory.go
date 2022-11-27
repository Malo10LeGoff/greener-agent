package main

import (

    "github.com/shirou/gopsutil/v3/mem"
	"math"
)

type MemoryUsageType struct {
	UsagePercent int `json:"usagePercent" jsonschema:"required"`
}

func getVirtualMemory() (MemoryUsageType) {
    v, _ := mem.VirtualMemory()

	memoryUsagePercent := int(math.Round(v.UsedPercent))

	memoryInfo := MemoryUsageType{memoryUsagePercent}

    return memoryInfo

}