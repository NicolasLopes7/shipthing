#!/bin/bash

find ./uploader-service ./lib ./config -name "*.go" | entr -r sh -c 'echo "Restarted uploader-service" && go run ./uploader-service'