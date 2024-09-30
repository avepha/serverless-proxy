#!/bin/bash

# Build the project
GOOS=linux GOARCH=amd64 go build -o .build/bootstrap main.go

# shellcheck disable=SC2164
cd .build
zip lambda-handler.zip bootstrap