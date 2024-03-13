.DEFAULT_GOAL := build

# This line keeps make from getting confused if a directory or file in your project has the same name as one of the
# listed targets
.PHONY:fmt vet build

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build

clean:
	go clean ./..
