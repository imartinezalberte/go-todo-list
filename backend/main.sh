#!/bin/bash
#

reflex -r '\.go$' -s -- sh -c 'go mod tidy && go run ./... -addr=:8080'
