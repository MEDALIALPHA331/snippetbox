build: 
	@go build -o ./tmp/main ./cmd/web/

run: build
	@./tmp/main