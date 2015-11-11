# Docker image for Drone's email notification plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-email .

FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-email /bin/
ENTRYPOINT ["/bin/drone-email"]