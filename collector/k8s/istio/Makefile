include ../Makefile.k8s

prep: local_docker install

install:
	helm repo add istio.io https://storage.googleapis.com/istio-release/releases/1.4.2/charts/
	helm repo update
	helm repo list
	kubectl create namespace istio-system
	helm install istiod istio/istiod -n istio-system --wait

inject_side_cars:


local_docker:
	(cd cmd/machine && docker build -f Dockerfile -t machine .)
	kind load docker-image machine:latest
