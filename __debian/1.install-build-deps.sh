#!/usr/bin/env bash

echo "Installing dependencies";echo
sudo apt-get update && sudo apt install -y pkg-config libvirt-daemon libvirt-dev gcc
echo;echo;echo "Done. Now installing the Go binaries"
sudo /opt/bin/install_golang.sh `head -1 ../go.version` amd64
