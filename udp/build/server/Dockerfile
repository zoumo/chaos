FROM cargo.caicloudprivatetest.com/caicloud/alpine:3.6

WORKDIR /root

RUN apk add --no-cache tcpdump

COPY bin/linux_amd64/server /usr/bin/server

ENTRYPOINT ["/usr/bin/server"]

EXPOSE 9090
