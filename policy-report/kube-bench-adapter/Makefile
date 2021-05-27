all: docker

build: fmt vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o policyreport .

fmt:
	go fmt ./...

vet:
	go vet ./...

docker: build
	docker build . -t mritunjay394/policyreport

codegen:
	./hack/update-codegen.sh