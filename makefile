run:
	echo "running App"
	go run cmd/main.go

test:
	echo "testing App"
	cd ./
	go test -v ./...

build:
	echo "Building Go executable..."
	# Create a build folder if it doesn't exist
	mkdir -p bin
	# Build the Go executable
	go build -o bin/send-email cmd/*.go
	echo "build Successful"

.PHONY: run test build