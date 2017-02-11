#!/bin/bash

set -e

# Build the binary
docker run -it --rm -v "$PWD":/go/src/proxy-http-status-code-only -w /go/src/proxy-http-status-code-only golang:1.7 bash -c "go get github.com/tools/godep && godep restore -v && CGO_ENABLED=0 GOOS=linux go build -a -v"
chmod ug+x proxy-http-status-code-only

# Get a recent copy of ca-certificates.crt
docker pull ubuntu:latest
docker build -f Dockerfile.certs -t temp-certs:latest .
ID=$(docker run -d temp-certs:latest)
docker cp $ID:/etc/ssl/certs/ca-certificates.crt .
docker rm $ID
docker rmi temp-certs:latest

docker build -t snarlysodboxer/proxy-http-status-code-only:latest .

rm -f proxy-http-status-code-only

docker push snarlysodboxer/proxy-http-status-code-only:latest

