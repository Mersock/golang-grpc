#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc greet/greetpb/greet.proto --go_out=. --go-grpc_out=.