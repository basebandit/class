SHELL := /bin/bash

run:
	go run app/sales-api/main.go

test:
	go test ./... -count=1
	staticcheck ./...
