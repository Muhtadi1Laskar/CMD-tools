package main

import (
	"flag"
	"fmt"
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

func bToMB(b uint64) uint64 {
	return b / 1024 / 1024
}


func main() {
	info := flag.String("info", "", "The information type")

	flag.Parse()

	switch *info {
	case "basic info":
		basicInfo()
	case "memory info":
		memoryStat()
	default:
		fmt.Println("Invalid type")
	}
}