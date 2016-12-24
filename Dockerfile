FROM ubuntu:16.04
MAINTAINER david amick <docker@davidamick.com>

ADD proxy-http-status-code-only /

ENTRYPOINT ["/proxy-http-status-code-only"]

