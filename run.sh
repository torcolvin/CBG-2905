#!/bin/bash
#
set -eux -o pipefail

cd ../sync_gateway && SG_EDITION=EE ./build.sh && cd -

./startup.sh

sleep 5 # wait for SG to start up
http --check-status --auth Administrator:password PUT :4985/db1/ bucket=bucket1 num_index_replicas:=0

http --check-status --auth Administrator:password PUT :4985/db2/ bucket=bucket2 num_index_replicas:=0

sleep 5 # replace with put DB online

./delete.sh
