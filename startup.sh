/bin/bash ./start_server.sh

docker exec couchbase couchbase-cli bucket-create --cluster localhost --username Administrator --password password --bucket bucket1 --bucket-type couchbase --bucket-ramsize 200
docker exec couchbase couchbase-cli bucket-create --cluster localhost --username Administrator --password password --bucket bucket2 --bucket-type couchbase --bucket-ramsize 200

pkill sync_gateway || true

sleep 5
../sync_gateway/bin/sync_gateway config.json > sg.log 2>&1 &
