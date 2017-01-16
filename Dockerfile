FROM alpine:3.4

RUN apk add --no-cache ca-certificates

ADD drone-email /bin/
ENTRYPOINT ["/bin/drone-email"]
