export GO111MODULE=on

test:
	@go test -v ./...
tidy:
	@go mod tidy -compat=1.17

run: tidy
	@go run main.go
