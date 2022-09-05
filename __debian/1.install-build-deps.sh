#!/usr/bin/env bash

echo "Installing dependencies";echo
apt-get update && apt-get install -y pkg-config libvirt-daemon libvirt-dev gcc
echo;echo;echo "Done. Now installing the Go binaries"
/opt/bin/install_golang.sh 1.19 amd64
