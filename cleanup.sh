#/bin/bash

docker-compose down
docker rmi practice-01_db practice-01_backend practice-01_agent -f
