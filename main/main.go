package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"time"
)

type Exporter struct {
	gauge    prometheus.Gauge
	gaugeVec prometheus.GaugeVec
}

//Counter: 重點方法inc ，一個累計型的metric
//CounterVec: Counter支援Label
//Gauge: 重點方法set，自己設定各種value 最常用
//GaugeVec: Gauge支援Label
//Histogram: 重點方法Observe，集計型的metric
//HistogramVec: Histogram支援Label
//Summary: 重點方法Observe，集計型的metric
//SummaryVec: Summary支援Label
//构造函数构造Exporter 结构体
func NewExporter(metricsPrefix string) *Exporter {

	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "gauge_metric",
		Help:      "This is a dummy gauge metric"})

	gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix, //指标数据名称前缀
		Name:      "cpu_usc_pct",
		Help:      "This is a cpu pct metric"},
		[]string{"cpu_user","cpu_sys"})

	return &Exporter{
		gauge:    gauge,
		gaugeVec: gaugeVec,
	}
}


//這兩個方法不能省略，一定要有
//
//Describe理論上不用做什麼特別的事，只要讓exporter metrics呼叫Describe方法就好
//
//而Collect則是要實作對metrics的收集,


var chData chan float64
func risCpu() (ch chan float64,gaugeData float64, gaugeVecData float64){

	gaugeData = rand.Float64()
	gaugeVecData = rand.Float64()


	return
}

func (e *Exporter)  Collect(ch chan<- prometheus.Metric) {
	gaugeData,gaugeVecData:=risCpu()
	e.gauge.Set(gaugeData)
	e.gaugeVec.WithLabelValues("xxxx","aaaa").Set(gaugeVecData)
	//e.gaugeVec.WithLabelValues("cpu_sys").Set(float64(0.64))
	e.gauge.Collect(ch)
	e.gaugeVec.Collect(ch)
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.gauge.Describe(ch)
	e.gaugeVec.Describe(ch)
}

func main() {
	fmt.Println(`
  This is a dummy example of prometheus exporter
  Access: http://127.0.0.1:9999
  `)


	// Define parameters

	metricsPath := "/metrics"
	listenAddress := ":9999"
	metricsPrefix := "dummy"

	// Register dummy exporter, not necessary

	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	// Launch http service

	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Dummy Exporter</title></head>
             <body>
             <h1>Dummy Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}