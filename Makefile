build:
	go build -o ./projectx-blockchain


run: build
	./projectx-blockchain

test:
	go test -v ./...