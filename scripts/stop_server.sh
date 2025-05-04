#!/bin/bash
set -eux

# kill any old binary or "go run"
pkill -f "serve --dev" || true
pkill -f "/opt/vibecheckr/vibecheckr" || true
