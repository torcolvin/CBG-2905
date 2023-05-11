#!/bin/bash

set -eux -o pipefail

http --check-status --auth Administrator:password DELETE :4985/db1/

docker exec couchbase couchbase-cli bucket-delete --cluster localhost --username Administrator --password password --bucket bucket1
