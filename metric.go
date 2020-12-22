package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/cclin81922/go-3rd-party/go-diskinfo"
)

var (
	diskAll = promauto.NewGauge(prometheus.GaugeOpts{
		Name:      "disk_all",
		Help:      "All disk space.",
	})
	diskUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name:      "disk_used",
		Help:      "Used disk space.",
	})
	diskFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name:      "disk_free",
		Help:      "Free disk space.",
	})
)

func Metric(disk diskinfo.DiskStatus) {
	diskAll.Set(float64(disk.All))
	diskUsed.Set(float64(disk.Used))
	diskFree.Set(float64(disk.Free))
}
