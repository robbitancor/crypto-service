docker build -t simple-microservice ../. && docker run --net=bridge -p 1323:1323  --rm --name simple-microservice simple-microservice
