package main

import (
	"fmt"
	"os"
)

/*
var (
	defaultMetricsPath = "/metrics"
	defaultTargetHost  = "example.my-exampe.svc.cluster.local:1520"
)
*/

type CollectorConfig struct {
	metricsPath string `yaml:"metrics_path"`
	targets     string `yaml:"targets"`
}

func main() {
	if os.Args[1] == "" || os.Args[2] == "" {
		panic("must run command as ./receiver-config METRICSPATH HOSTTARGET")
	}
	if err := insertCollectorConfig(os.Args[1], os.Args[2]); err != nil {
		panic(err)
	}
	fmt.Println("Successfull collector config insertion complete")
}

func insertCollectorConfig(metricsPath, targetHost string) ([]byte, error) {

}
