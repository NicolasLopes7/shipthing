#!/bin/bash

find ./deployer-service ./lib ./config -name "*.go" | entr -r sh -c 'echo "Restarted deployer-service" && go run ./deployer-service'
