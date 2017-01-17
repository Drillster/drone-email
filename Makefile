all: test build publish

test:
	go vet
	go test -cover -coverprofile=coverage.out

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build -t drillster/drone-email:latest .

publish:
	bash publish.sh