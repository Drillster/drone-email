all: test build

test:
	go vet
	go test -cover -coverprofile=coverage.out

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
