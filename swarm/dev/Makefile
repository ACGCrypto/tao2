.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/tao2-core/tao2-core/cmd/swarm
	go build -o bin/tao     github.com/tao2-core/tao2-core/cmd/tao
	go build -o bin/bootnode github.com/tao2-core/tao2-core/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/tao2-core/tao2-core/swarm/...
