#!/usr/bin/env bash

echo "Installing BuildRequires dependencies";echo
grep ^BuildRequires certificateManager.spec |awk -F\: '{print "sudo dnf install -y"$2}'|sed -e 's/,/ /g' | sh
echo;echo;echo "Done. Now installing the Go binaries"
#/opt/bin/install_golang.sh $(grep ^go src/go.mod | awk '{print $2}') amd64

if [ "$#" -lt 2 ]; then
	echo "You must specify the version number (ex: 1.18.3) and arch (ex: amd64) to download."
	exit 3
fi

export VER=${1}
export ARCH=${2}

echo "Fetching archive..."
sudo wget -q https://go.dev/dl/go${VER}.linux-${ARCH}.tar.gz -O /tmp/go.tar.gz -O /opt/go.tar.gz

echo "Unarchiving..."
cd /opt ; sudo rm -rf go;sudo tar zxf go.tar.gz; sudo rm -f go.tar.gz

echo "Completed."

