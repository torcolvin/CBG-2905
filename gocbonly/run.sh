#!/bin/bash
#
set -eux -o pipefail

/bin/bash ../start_server.sh

docker exec couchbase couchbase-cli bucket-create --cluster localhost --username Administrator --password password --bucket bucket1 --bucket-type couchbase --bucket-ramsize 200
docker exec couchbase couchbase-cli bucket-create --cluster localhost --username Administrator --password password --bucket bucket2 --bucket-type couchbase --bucket-ramsize 200

go run main.go
