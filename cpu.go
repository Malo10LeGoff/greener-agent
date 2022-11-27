package main

import (
	"runtime"
)

type CPUUsageType struct {
	NumCores int `json:"numCores" jsonschema:"required"`
}

func getCPU() (CPUUsageType) {

	cpuInfo := CPUUsageType{runtime.NumCPU()}

	return cpuInfo

}