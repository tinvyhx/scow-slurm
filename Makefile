ARCH ?= amd64

protos: 
	buf generate --template buf.gen.yaml ../protos

run: 
	go run *.go 

build:
	CGO_BUILD=0 GOARCH=${ARCH} go build -o scow-slurm-adapter-${ARCH}

test:
	go test
