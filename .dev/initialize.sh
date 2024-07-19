#!/bin/bash

echo "=== Initialize Requirements ==="

echo "Installing gotest..."
go get github.com/rakyll/gotest

echo "Installing mockgen..."
go get github.com/golang/mock/mockgen

echo "Installing air..."
go install github.com/air-verse/air@latest

echo "=== Initialization Complete ==="