**crypto-service**

crypto-service

Microservice that has one endpoint, 1 input (address) and
an output of JSON format for gas price, block number, and balance

This project runs on assigned port of 1323, and has following dependencies:
- echo (high-performant, minimalist web server) - github.com/labstack/echo/v4 v4.12.0
- mongodb (storage) - go.mongodb.org/mongo-driver/v2
- mock (unit testing) - github.com/golang/mock v1.6.0
- gopkg (reading config files) - gopkg.in/yaml.v2 v2.4.0

Security.
 I implemented a simple way of securing the webservice from (XSS attack),
 content-type sniffing, clickjacking, insecure connection and other code injection attacks
 using echo (https://echo.labstack.com/docs/middleware/secure)


Running the project.
1. clone the repository.
2. provide the TOKEN in the config file on simple-microservice/internal/config
3. make sure you have docker/docker-desktop installed and running on your machine
4. run the _deploy.bat_ to build the image and run an instance in docker
5. you may test the service on _localhost:1323/etherium/getBalance_ with a POST request. You can use this JSON payload
   `{
   "address":"0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae"
   }`

Optional:
 in the _/scripts_ folder, you can run each individually to:
  - run mongodb (run run-mongodb.bat)
  - deploy the project (run _deploy.bat_)
 pulling images:
  - mongodb and redis (run _docker-pull.bat_)


Known issues/limitations:
 - mongodb and redis implementation for demo only
 - test are found in mock folder but not used yet

Thank you