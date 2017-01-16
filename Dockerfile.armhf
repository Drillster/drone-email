FROM armhfbuild/alpine:3.4

RUN apk update && \
    apk add --no-cache ca-certificates

ADD drone-email /bin/
ENTRYPOINT ["/bin/drone-email"]