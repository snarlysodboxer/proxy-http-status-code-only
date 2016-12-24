#!/bin/bash

docker run --rm -v "$PWD":/usr/src/proxy-http-status-code-only -w /usr/src/proxy-http-status-code-only golang:1.6 go build -v
chmod ug+x proxy-http-status-code-only

docker build -t snarlysodboxer/proxy-http-status-code-only:latest .

docker push snarlysodboxer/proxy-http-status-code-only:latest

