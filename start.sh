#/bin/bash

docker build -t pra .
docker run --name runpra -p 3030:3030 -d  pra