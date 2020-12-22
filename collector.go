package main

import (
	"log"
	"time"
    "github.com/cclin81922/go-3rd-party/go-diskinfo"
)

func statusUpdate() {
    disk := diskinfo.DiskUsage("/")
    log.Printf("All: %.2f GB %.2f GiB\n", float64(disk.All)/float64(diskinfo.GB), float64(disk.All)/float64(diskinfo.GiB))
    log.Printf("Used: %.2f GB %.2f GiB\n", float64(disk.Used)/float64(diskinfo.GB), float64(disk.Used)/float64(diskinfo.GiB))
    log.Printf("Free: %.2f GB %.2f GiB\n", float64(disk.Free)/float64(diskinfo.GB), float64(disk.Free)/float64(diskinfo.GiB))

	Metric(disk)
}

func Collect() {
	go func() {
		for range time.Tick(time.Minute) {
			statusUpdate()
		}
	}()
}
