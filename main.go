package main

import (
    "fmt"
	"time"
)

func main() {

	running := true

	for running {

		cpuInfo := getCPU()

		memoryInfo := getVirtualMemory()

		fmt.Printf("Total Resource Consumed: NumCores : %v , UsedPercentMemory: %v\n", cpuInfo.NumCores, memoryInfo.UsagePercent)

		time.Sleep(8 * time.Second)

	}

}