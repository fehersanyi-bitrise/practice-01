#/bin/bash

docker build -t infradb .
echo "Container build successfully"
sleep 2
docker run --name infradb -e POSTGRES_USER=infra -d -p 5432:5432 infradb
echo "database up and running"
sleep 2

docker exec -it infradb psql -d infra -U infra -f db.sql
echo "table created successfully"

