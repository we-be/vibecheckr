#!/bin/bash
set -eux

cd /opt/vibecheckr
# build binary
go build -o /opt/vibecheckr/vibecheckr cmd/main.go
