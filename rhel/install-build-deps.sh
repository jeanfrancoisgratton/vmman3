#!/usr/bin/env bash

grep ^BuildRequires vmman3.spec |awk -F\: '{print "sudo microdnf install -y"$2}'|sed -e 's/,/ /g' | sh
/opt/bin/install_golang.sh 1.19 amd64
