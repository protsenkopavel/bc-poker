build:
	@go build -o bin/bcpocker

run: build
	@./bin/bcpocker

test:
	go test -v ./...