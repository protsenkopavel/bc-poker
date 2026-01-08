build:
	@go build -o bin/bcpoker

run: build
	@./bin/bcpoker

test:
	go test -v ./...