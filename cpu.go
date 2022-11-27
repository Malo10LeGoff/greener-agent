package main

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/cpu"
)

func getCPU() {

	p, _ := cpu.Info()

    fmt.Println(p)

}