build:
	@go build -o ./bin/goEcom ./cmd/main.go

run:build
	@./bin/goEcom