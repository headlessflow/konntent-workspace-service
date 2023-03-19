cid=$( \
  docker run -d \
  --name integration-yugabyte-cluster \
  -p 5433:5433 \
  -e \
  yugabytedb/yugabyte:latest bin/yugabyted start)