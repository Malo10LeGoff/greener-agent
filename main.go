package main

import (
    "fmt"
	"time"
	"encoding/json"

)

func main() {

	running := true

	type ResourceUsage struct {
		CPUUsage CPUUsageType `json:"cpuUsage"`
		MemoryUsage MemoryUsageType `json:"memoryUsage"`
	}

	for running {

		cpuInfo := getCPU()

		memoryInfo := getVirtualMemory()

		fmt.Printf("Total Resource Consumed: NumCores : %v , UsedPercentMemory: %v\n", cpuInfo.NumCores, memoryInfo.UsagePercent)

		resourceInfo := ResourceUsage{cpuInfo, memoryInfo}

		serializedResourceUsage, err := json.Marshal(resourceInfo)

		fmt.Println(string(serializedResourceUsage))

		if err != nil {
			fmt.Println(err)
			return
		}

		// Create the schema : 
		//jsonschema.Reflect(&ResourceUsage{})

		var deserializedResourceUsage ResourceUsage

		json.Unmarshal(serializedResourceUsage, &deserializedResourceUsage)

		time.Sleep(8 * time.Second)

	}

}