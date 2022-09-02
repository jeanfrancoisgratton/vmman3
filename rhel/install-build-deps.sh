#!/usr/bin/env bash

echo "Installing BuildRequires dependencies";echo
grep ^BuildRequires ../vmman3.spec |awk -F\: '{print "sudo microdnf install -y"$2}'|sed -e 's/,/ /g' | sh
echo;echo;echo "Done. Now installing the Go binaries"
/opt/bin/install_golang.sh 1.19 amd64
