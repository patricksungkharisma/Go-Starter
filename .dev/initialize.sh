#!/bin/bash

echo "=== Initialize Requirements ==="

echo "Installing gotest..."
go get github.com/rakyll/gotest

echo "Installing mockgen..."
go install github.com/golang/mock/mockgen@latest

echo "Installing air..."
go install github.com/air-verse/air@latest

echo "Downloading dependencies..."
go mod download

echo "Vendor setup..."
go mod vendor

echo "=== Initialization Complete ==="