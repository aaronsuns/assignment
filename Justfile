build:
  go build -o bin/resolveRange cmd/resolveRange/resolveRange.go
  go build -o bin/restAPI cmd/restAPI/restAPI.go

lint:
  golangci-lint run --timeout 1m

testall:
  go test ./...
