export GO111MODULE=on

test:
	@go test -v ./...
tidy:
	@go mod tidy -compat=1.19
config: tidy
	go run ./app/generate/generate.go
run: tidy
	@go run main.go
