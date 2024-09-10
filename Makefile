build:
	@go build -o bin/ecom cmd/api/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/ecom