.SILENT:

run: linter
	go run ./cmd/grpc-client/main.go
linter:
	golangci-lint run ./... --config=./.golangci.yaml