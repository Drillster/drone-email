# Docker image for the Drone Email plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-email
#     make deps build
#     docker build --rm=true -t plugins/drone-email .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-email /bin/
ENTRYPOINT ["/bin/drone-email"]
