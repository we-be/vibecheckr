#!/bin/bash
set -eux

cd /opt/vibecheckr
# start in background
nohup ./vibecheckr serve --dev > /opt/vibecheckr/server.log 2>&1 &
