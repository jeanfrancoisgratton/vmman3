#!/usr/bin/env bash

echo "Installing BuildRequires dependencies";echo
grep ^BuildRequires vmman3.spec |awk -F\: '{print "sudo dnf install -y"$2}'|sed -e 's/,/ /g' | sh
echo;echo;echo "Done. Now installing the Go binaries"
/opt/bin/install_golang.sh `grep ^go ../source/go.mod | awk '{print $2}'` amd64
