FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /taoblockchain
RUN cd /taoblockchain && make tao

FROM alpine:latest

WORKDIR /taoblockchain

COPY --from=builder /taoblockchain/build/bin/tao /usr/local/bin/tao

RUN chmod +x /usr/local/bin/tao

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/tao"]

CMD ["--help"]
