#clean:
#	rm server
build:
	cd ./cmd/server/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./server.go
all: build
