package common

import (
	"log"
	"runtime"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// LogMemoryUsage shows runtime memory information - more information at https://golang.org/pkg/runtime/#MemStats
func LogMemoryUsage() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB | TotalAlloc = %v MiB | Sys = %v MiB | NumGC = %v", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}
