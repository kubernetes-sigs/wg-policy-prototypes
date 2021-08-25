SOURCES := $(shell find . -name '*.go')

## Builds the imgvuln binary
trivy-adapter: $(SOURCES)
	CGO_ENABLED=0 go build -o ./trivy-adapter ./cmd/trivy-adapter/main.go
install:	
	sudo mv trivy-adapter /usr/local/bin
