#/bin/bash

echo "gyasz up image"
docker build -t pra .
sleep 2
echo "gyaszing container from image"
docker run --name runpra -p 3030:3030 -d  pra