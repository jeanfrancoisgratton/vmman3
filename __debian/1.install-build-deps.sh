#!/usr/bin/env bash

echo "Installing dependencies";echo
sudo apt-get update && sudo apt-get install -y pkg-config libvirt-daemon libvirt-dev gcc
echo;echo;echo "Done. Now installing the Go binaries"
sudo /opt/bin/install_golang.sh 1.19 amd64
