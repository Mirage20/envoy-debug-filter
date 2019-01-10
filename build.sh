#!/usr/bin/env bash

export CGO_ENABLED=0

go build -o envoy-debug-filter -x ./main.go

docker build -t mirage20/envoy-debug-filter .

docker push mirage20/envoy-debug-filter
