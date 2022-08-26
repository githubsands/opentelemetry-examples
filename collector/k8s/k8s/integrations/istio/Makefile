docker_build:
	eval $(minikube docker-env)
	cd cmd/machine && docker build -f Dockerfile -t machine .

deploy_application:
	helm install helm-application --generate-name

deploy_telemetry:
	helm install helm-otel-istio --generate-name

install_istio_demo:
	helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.4.2/charts/
	helm repo update
	helm repo list
	kubectl create namespace istio-system
 	helm install istio-base istio/base -n istio-system
	helm install istiod istio/istiod -n istio-system --wait
	helm install helm-application
	helm install helm-otel-istio

run_istio_demo:
	./scripts/run-istio-demo.sh
	./scripts/generator.sh

view_istio_demo:
	tmux attach-session -t istio-demo
	./scripts/generator.sh

exit_istio_demo:
	tmux detach-client istio-demo

minikube:
	eval $(minikube docker-env)
