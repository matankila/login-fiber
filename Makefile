# tests the app & returns coverage %
test:
	go test ./... -coverprofile cover.out -covermode=atomic
	go tool cover -func cover.out | grep total | awk '{print $$3}'

# builds none dependent binary file
build:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/main ./cmd/main.go