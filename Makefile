# tests the app & returns coverage %
test:
	go test $$(go list ./... | grep -v 'mock\|cmd\|docs\|global') -coverprofile cover.out -covermode=atomic
	go tool cover -func cover.out | grep total | awk '{print $$3}'

# run integration test against runtime env
integration:
	go test ./cmd/integration_test -tags integration -host=$(HOST)

# builds none dependent binary file
build:
	export CGO_ENABLED=0
	export GOOS=linux
	export GOARCH=amd64
	go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/main ./cmd/main.go

# team2registry.azurecr.io
# make cicd USER=* PASSWORD=* HOST=* CONTAINER_REGISTRY=*
cicd:
	echo "start ci(compile, unit-tests & coverage)"
	docker login $(CONTAINER_REGISTRY) -u $(USER) -p $(PASSWORD)
	docker build . -t $(CONTAINER_REGISTRY)/ttlogin:v1.0.7
	echo "start cd(deployment to cloud)"
	docker push $(CONTAINER_REGISTRY)/ttlogin:v1.0.7
	echo "start integration tests.."
	sleep 10
	go test ./cmd/integration_test -tags integration -host=$(HOST)