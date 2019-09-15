#!/bin/bash

set -e
COMP=$1
NS=$2
NAME=$3
export KUBECONFIG=$4
cd $COMP

echo "+PRE"
./PRE
echo "-return $?"

. ./INPUT
echo "+CREATE $NS $NAME"
./CREATE $NS $NAME
echo "-return $?"

. ./INPUT1
echo "+UPDATE_CHECK $NS $NAME"
./UPDATE_CHECK $NS $NAME
echo "-return $?"

. ./INPUT1
echo "+UPDATE $NS $NAME"
./UPDATE $NS $NAME
echo "-return $?"

. ./INPUT1
echo "+DELETE $NS $NAME"
./DELETE $NS $NAME
echo "-return $?"
