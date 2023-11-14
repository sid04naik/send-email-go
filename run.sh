#!/bin/bash

echo "Building Go executable..."

# Create a build folder if it doesn't exist
buildFolder="bin"
mkdir -p $buildFolder

# Build the Go executable
go build -o $buildFolder/send-email cmd/*.go
echo "Building Executable file: ./$buildFolder/send-email"

