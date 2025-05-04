#!/bin/bash
set -eux

# wait up to 30s for service
timeout=30
count=0
while ! curl -sf http://localhost:8090/; do
  ((count++))
  if [ $count -ge $timeout ]; then
    echo "ERROR: vibecheckr failed to come up"
    exit 1
  fi
  sleep 1
done
