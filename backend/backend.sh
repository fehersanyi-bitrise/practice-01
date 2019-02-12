#/bin/bash

echo "gyasz up image"
docker build -t backend .
sleep 2
echo "gyaszing container from image"
docker run --name backend-server -p 3030:3030 -d  backend