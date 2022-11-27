package main


import "time"

func main() {

	running := true

	for running {

		getCPU()

		getVirtualMemory()

		time.Sleep(8 * time.Second)

	}



}