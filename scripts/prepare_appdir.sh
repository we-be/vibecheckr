#!/bin/bash
set -eux

# Ensure the deploy directory exists and is writable by the deploy user
mkdir -p /opt/vibecheckr
chown -R ec2-user:ec2-user /opt/vibecheckr