SHELL := /bin/bash

run:
	go run app/sales-api/main.go

test:
	go test ./... -count=1
	staticcheck ./...

sales-api:
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api-amd64:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.

kind-up:
	kind create cluster --image kindest/node:v1.18.8 --name ardan-starter-cluster --config zarf/k8s/dev/kind-config.yaml

kind-down:
	kind delete cluster --name ardan-starter-cluster

kind-load:
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster

kind-services:
	kustomize build zarf/k8s/dev | kubectl apply -f -

kind-logs:
	kubectl logs -lapp=sales-api --all-containers=true -f

kind-status:
	kubectl get nodes
	kubectl get pods --watch

kind-status-full:
	kubectl describe pod -lapp=sales-api

kind-sales-api: sales-api
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster
	kubectl delete pods -lapp=sales-api