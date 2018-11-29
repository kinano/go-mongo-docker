
## Pre-requisites
* Define $GOPATH on your local box
```
# If you are using bash
echo 'export GOPATH=~/Development/go' >> ~/.bash_profile
source ~/.bash_profile
# If you are using zsh
echo 'export GOPATH=~/Development/go' >> ~/.zshrc
source ~/.zshrc
```

## Running the container locally using Docker
```
docker-compose build # This builds all the images you need to run the API locally
docker-compose up
```

## How do I shutdown the API locally?
```
docker-compose down
```

## How did we implement Hot Reloading for GoLang?
We are using [CompileDaemon](https://github.com/githubnemo/CompileDaemon) on local DEV. See [Dockerfile.local](https://github.com/kinano/go-mongo-docker/blob/develop/src/Dockerfile.local#L11) for details. 

## Deploy to EBS
* Build the docker containers for AWS
```
docker build -t kinano/api-go ./src
docker build -t kinano/api-nginx ./nginx
docker image push kinano/api-go
docker image push kinano/api-nginx
```
* Create a DB on Mongo DB Atlas (or any Mongo DB cloud provider)
* Build a medium EBS environment (MongoDB driver used in this app was failing to compile on small instances)
* Remember to generate and use an EC2 key pair to be able to ssh into the created EC2 instance
* Upload the `Dockerrun.aws.json` to EBS
* Add the following config keys on EBS
```
APP_PORT
MONGO_DB_NAME
MONGO_DB_BOOKINGS_COLLECTION
MONGO_DB_LOGS_COLLECTION
MONGO_USERNAME
MONGO_PASSWORD
MONGO_HOST=HOST:PORT
```
* Create a security group for your application and add the following inbound rules
```
type    protocol    port range  source
SSH     TCP         22          Your IP address
HTTP    TCP         80          0.0.0.0/0 (This is created by default when you create a web app on EBS)
Custom  TCP         5000        YOUR SECURITY GROUP ID (This allows the nginx container to forward api traffic to the API container using port 5000. Replace the port with whatever APP_PORT env variable you used on EBS above)
```
