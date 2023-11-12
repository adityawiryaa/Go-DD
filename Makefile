BINARY=engine
build:
	go build -o go-driven-design cmd/main.go
start:
	go run cmd/main.go
test:
	go test -v -cover -coverprofile=coverage.out -covermode=atomic ./... && go tool cover -func coverage.out
