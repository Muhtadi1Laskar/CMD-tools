package main

import (
	"fmt"
	"os"
	"runtime"
)

func basicInfo() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Go Version:", runtime.Version())
}

func memoryStat() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB\n", bToMB(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\n", bToMB(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\n", bToMB(m.Sys))
	fmt.Printf("NumGC = %v\n", m.NumGC)
}

func hostInfo() {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hostname:", hostName)
}

func bToMB(b uint64) uint64 {
	return b / 1024 / 1024
}


func main() {
	fmt.Printf("\nBasic Info\n")
	basicInfo()

	fmt.Printf("\nMemory Stat\n")
	memoryStat()

	fmt.Printf("\nHost Info\n")
	hostInfo()
}