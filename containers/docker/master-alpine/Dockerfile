FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone https://github.com/taoblockchain/tao2 tao && \
  (cd tao && make tao) && \
  cp tao/build/bin/tao /tao && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /toa && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 20202

ENTRYPOINT ["/tao"]
