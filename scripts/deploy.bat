docker build -t simple-microservice ../ && docker run --network crypto-network -p 1323:1323  --rm --name simple-microservice simple-microservice
