#!/usr/bin/env bash

grep ^BuildRequires vmman3.spec |awk -F\: '{print "microdnf install -y"$2}'|sed -e 's/,/ /g' | sh