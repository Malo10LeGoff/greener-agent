package main

import (
	"runtime"
)

type CPUUsage struct {
	NumCores int
}

func getCPU() (CPUUsage) {

	cpuInfo := CPUUsage{runtime.NumCPU()}

	return cpuInfo

}