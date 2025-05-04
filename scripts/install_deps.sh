#!/bin/bash
set -eux

cd /opt/vibecheckr
# download all modules
go mod download
