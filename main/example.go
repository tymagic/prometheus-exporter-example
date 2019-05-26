package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

var addr = flag.String("listen-address", ":9999", "the address to listen on for Http Request!")

var (
	opsQueued = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "example",
		Subsystem: "cpu_rate",
		Name:      "cpu_usage",
		Help:      "Cpu Usage Rate !",
	})
	opsQueuedVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "example1",
		Subsystem: "cpu_rate1",
		Name:      "cpu_usage1",
		Help:      "Cpu Usage Rate1 !",
	}, []string{"hostname"})
)

func init() {
	prometheus.MustRegister(opsQueued)
	prometheus.MustRegister(opsQueuedVec)
}

func collectCpuInfo() {
	for {
		fmt.Println(time.Now())
		opsQueued.Add(4)
		opsQueuedVec.WithLabelValues("hostname")
		time.Sleep(time.Second * 3)
	}
}

func main() {
	flag.Parse()
	go collectCpuInfo()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
