package main

import (
	"testing"

	"gopkg.in/yaml.v2"
	"sigs.k8s.io/yaml"
)

var collectorConfig = `
release-name: "examples"

cert-manager:
  installCRDs: true # <- this does not work as attended we must rely on the presync event within the the makefile

# otel collector helm docs can be found here:
#   https://artifacthub.io/packages/helm/opentelemetry-helm/opentelemetry-collector
opentelemetry-collector:
  mode: "deployment"

  containerLogs:
    enabled: true

      # place your collector configuration here [i.e https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/examples/demo/otel-collector-config.yaml]
  config:
    receivers:
      prometheus/collectd:
        use_start_time_metric: false
        start_time_metric_regex: '^(.+_)*process_start_time_seconds$'
        config:
          scrape_configs:
            - job_name: 'component-scraper'
              scrape_interval: 5s
              metrics_path: "/stats/prometheus"
              static_configs:
                - targets: [TEST]
    exporters:
      otlp:
        endpoint: ingest.lightstep.com:443
        headers:
          - lightstep-access-token: lightstep-secret
    service:
      pipelines:
        metrics:
          receivers: [prometheus/collectd]
          exporters: [otlp]

secrets:
  lightStepPlatformAccessToken: ${LIGHTSTEP_API_KEY}`

// TestMain tests whether if input args change the configuration
func TestMain(t *testing.T) {
	collectorConfig := new(CollectorConfig)
	collectorConfigBytes := yaml.Marshal(collectorConfig, collectorConfigBytes)
}

func TestInsertCollectorConfig(t *testing.T) {
	err := insertCollectorConfig("test", "test")
	if err != nil {
		t.Fatalf("Failed to insertCollectorConfig")
	}
}
