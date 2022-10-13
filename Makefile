build:
	go build -o ./bin/projectx-blockchain


run: build
	./bin/projectx-blockchain

test:
	go test ./...