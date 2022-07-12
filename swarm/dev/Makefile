.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/janotchain/janotchain/cmd/swarm
	go build -o bin/tomo     github.com/janotchain/janotchain/cmd/tomo
	go build -o bin/bootnode github.com/janotchain/janotchain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/janotchain/janotchain/swarm/...
