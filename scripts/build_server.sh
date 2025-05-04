#!/bin/bash
set -eux
# Ensure Go binary is in PATH when run by CodeDeploy agent
export PATH=$PATH:/usr/local/go/bin

cd /opt/vibecheckr
# build binary
go build -o /opt/vibecheckr/vibecheckr cmd/main.go
