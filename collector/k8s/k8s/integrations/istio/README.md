# Istio telemetry integrations:

Istio provides a solution for telemetry through its telemetry API (or CRUD resource). A brief overview of this Istio feature is that fundamentally utilizes a set of envoy proxies. Each application deployed in an istio environment can have an envoy injected adjacent to an application. This adjacent envoy then can deliver fine grain telemetry for each living application container within.

This repositorty provides a brief example to get you up to speed with istio telemtry and the open telemetry.

The following directory provides two helm charts for an Open telemetry Istio integration.

(1) A sample application to drive traffic for observing purposes. Some notable istio and telemetry specific nuances within this chart are:

* automated side car injections istio annotations within kubernetes namespace templates where pods, replicasets, or deployments are deployed.

(2) An open telemetry specific helm chart needed to drive telemtry for that specific application. Some major components within this chart are:

* application specific telemetry resources (a CRUD provided by the Istio project)
* istio mesh-configuration

# Building the project:

. Have a based Lighstep open telemetry setup deployed noted in the `github.com/opentelemetry-examples/collector/k8s` directory
. (cd cmd/machine && docker build . && kind load docker-image machine:latest)
. `helm install helm-application --generate-name`
. `helm install helm-otel-istio --generate-name`
