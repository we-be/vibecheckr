#!/bin/bash
set -eux
# Ensure Go binary is in PATH when run by CodeDeploy agent
export PATH=$PATH:/usr/local/go/bin

cd /opt/vibecheckr
# download all modules
go mod download
