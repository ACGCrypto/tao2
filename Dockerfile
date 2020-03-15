FROM golang:1.12-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /tao2
RUN cd /tao2 && make tao

FROM alpine:latest

WORKDIR /tao2

COPY --from=builder /tao2/build/bin/tao /usr/local/bin/tao

RUN chmod +x /usr/local/bin/tao

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/local/bin/tao"]

CMD ["--help"]
