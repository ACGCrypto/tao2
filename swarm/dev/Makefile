.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/taoblockchain/tao2/cmd/swarm
	go build -o bin/tao     github.com/taoblockchain/tao2/cmd/tao
	go build -o bin/bootnode github.com/taoblockchain/tao2/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/taoblockchain/tao2/swarm/...
