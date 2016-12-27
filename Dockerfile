FROM scratch
MAINTAINER david amick <docker@davidamick.com>

ADD ca-certificates.crt /etc/ssl/certs/
ADD proxy-http-status-code-only /

ENTRYPOINT ["/proxy-http-status-code-only"]

