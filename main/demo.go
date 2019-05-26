package main

import "github.com/prometheus/client_golang/prometheus"

var (
	Collector prometheus.Collector
)

type JavaCollector struct {
	Collectors map[string]Collector
}
