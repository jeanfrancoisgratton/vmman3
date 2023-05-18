#!/usr/bin/env bash

egrep "gotest|golang|github" go.mod|awk '{print "go get "$1}'|sh
