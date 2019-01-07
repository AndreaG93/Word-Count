#!/bin/sh

# Building image from a Dockerfile...
# --file 		-> Name of the Dockerfile (Default is ‘PATH/Dockerfile’)
# --tag      	-> Name and optionally a tag in the ‘name:tag’ format
docker build --file Dockerfile.server --tag wc_server_image .
docker build --file Dockerfile.worker --tag wc_worker_image .

# Create new containers...
# --name 		-> Assign a name to the container.
# --network		-> Connect a container to a network.
# --env 		-> Set environment variables

# server
docker create --network host --name word_count_server wc_server_image

# Workers
docker create --network host --name word_count_worker0 --env WORKER_ADDRESS=localhost:1000 wc_worker_image
docker create --network host --name word_count_worker1 --env WORKER_ADDRESS=localhost:1001 wc_worker_image
docker create --network host --name word_count_worker2 --env WORKER_ADDRESS=localhost:1002 wc_worker_image
docker create --network host --name word_count_worker3 --env WORKER_ADDRESS=localhost:1003 wc_worker_image
docker create --network host --name word_count_worker4 --env WORKER_ADDRESS=localhost:1004 wc_worker_image
docker create --network host --name word_count_worker5 --env WORKER_ADDRESS=localhost:1005 wc_worker_image
docker create --network host --name word_count_worker6 --env WORKER_ADDRESS=localhost:1006 wc_worker_image
docker create --network host --name word_count_worker7 --env WORKER_ADDRESS=localhost:1007 wc_worker_image
docker create --network host --name word_count_worker8 --env WORKER_ADDRESS=localhost:1008 wc_worker_image
docker create --network host --name word_count_worker9 --env WORKER_ADDRESS=localhost:1009 wc_worker_image

# Start containers...

docker container start word_count_server
# Wait some time...
sleep 3
docker container start word_count_worker0
docker container start word_count_worker1
docker container start word_count_worker2
docker container start word_count_worker3
docker container start word_count_worker4
docker container start word_count_worker5
docker container start word_count_worker6
docker container start word_count_worker7
docker container start word_count_worker8
docker container start word_count_worker9
