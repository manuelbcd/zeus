#!/bin/bash
environment=$1

if [ "$environment" == "dev" ]; then
	dlv debug --headless --listen=:2345 --api-version=2
else
    go run main.go
fi