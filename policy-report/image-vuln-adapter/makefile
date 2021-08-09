SOURCES := $(shell find . -name '*.go')

## Builds the imgvuln binary
imgvuln: $(SOURCES)
	CGO_ENABLED=0 go build -o ./bin/imgvuln ./cmd/imgvuln/main.go
