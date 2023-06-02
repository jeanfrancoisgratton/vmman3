#!/usr/bin/env bash

OUTPUT=/opt/bin

if [ "$#" -gt 0 ]; then
    OUTPUT=$1
fi
go build -o ${OUTPUT}/vmman3 .

