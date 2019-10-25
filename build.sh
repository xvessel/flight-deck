#!/usr/bin/env bash
cd `dirname $0`

export GOPATH=${PWD}
export GO111MODULE=off

go build -o ./bin/deck  deck
go build -o ./bin/dkctl client
