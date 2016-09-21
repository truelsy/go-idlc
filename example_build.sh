#!/bin/bash

go build -ldflags "-s -w" -o go-idlc
./go-idlc -l=go example/example.idl
gofmt -w example/example.go