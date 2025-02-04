# Lightstep OpenTelemetry Examples

This repository contains example code and resources for configuring a variety of languages with OpenTelemetry. The directory structure for each language looks like this:

```
go
├── launcher - these applications are using the opentelemetry launcher, a config layer built
               on top of opentelemetry
├── opentelemetry - these applications are configured and instrumented using opentelemetry
                    directly
└── opentracing - these applications are using the OpenTracing bridge available in
                  OpenTelemetry. The code is instrumented using OpenTracing.
```

### Running examples

All the applications in this repository can be launched using docker-compose. In order to send data to Lightstep, you can update the configuration using the steps below before starting docker-compose:

```bash
git clone https://github.com/lightstep/opentelemetry-examples && cd opentelemetry-examples

# copy the example environment variable file
# and update the access token
cp ./config/example.env .env
sed -i '' 's/<ACCESS TOKEN>/YOUR TOKEN HERE/' .env
# if using the AWS OpenTelemetry Distro, use `example-aws-collector-config.yaml`
cp ./config/example-collector-config.yaml ./config/collector-config.yaml
sed -i '' 's/<ACCESS TOKEN>/YOUR TOKEN HERE/' ./config/collector-config.yaml

docker-compose up
```

The following client/server applications use different mechanism for sending data to Lightstep. The following examples are configured in the `docker-compose.yml` file:

| name             | description                                                  |
| ---------------- | ------------------------------------------------------------ |
| go-opentracing   | client/server example instrumented via lightstep-tracer-go   |
| go-opentelemetry | client/server example instrumented via OpenTelemetry and the OTLP exporter |
| go-launcher      | client/server example instrumented via OpenTelemetry and the Launcher for configuration |
| py-collector     | client/server example instrumented via OpenTelemetry and the OTLP exporter combined with the OpenTelemetry Collector |
| py-opentelemetry | client/server example instrumented via OpenTelemetry and the OTLP exporter |
| py-launcher      | client/server example instrumented via OpenTelemetry and the Launcher for configuration |
| java             | client/server example instrumented via special agent         |
| java-otlp        | client/server example instrumented via OpenTelemetry and the OTLP exporter |
| java-launcher    | client/server example instrumented via OpenTelemetry and the Launcher for configuration |
| js-ot-shim       | client/server example instrumented via OpenTelemetry and JS Launcher with OpenTracing |

### OpenTelemetry Collector Integrations

Additionally, the repo contains example for producing telemetry via various OpenTelemetry Collector Integrations. The integrations can be found under `./collector/<integration>`. To produce telemetry using these integrations:

1. Clone this repository and `cd` into the directory for the integration of choice:

    `cd ./collector/mysql`

2. Export your Lightstep Access Token

    `export LS_ACCESS_TOKEN=<your token>`.

3. Run the environment using Docker compose

    `docker compose up`

4. Access the Lightstep dashboard for the integration.

------

*Made with* ![:heart:](https://a.slack-edge.com/production-standard-emoji-assets/10.2/apple-medium/2764-fe0f.png) *@ [Lightstep](http://lightstep.com/)*