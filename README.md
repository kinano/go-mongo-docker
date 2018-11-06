# Sk-Next-API

Welcome to Skedaddle's Next Gen API (late fall 2018)


## Pre-requisites
* See [Docker pre-requisites](https://github.com/Skedaddle/infrastructure/tree/master/docker#pre-requisites)
* Run the Mongo Docker container
```
# Replace the path to point to the infrastructure repo on your machine
~/Development/infrastructure/docker/mongo/bin/run.sh
```
* DEP dependency manager
```
brew install dep && brew upgrade dep
```
* Define $GOPATH on your local box
```
# If you are using bash
echo 'export GOPATH=~/Development/go' >> ~/.bash_profile
source ~/.bash_profile
# If you are using zsh
echo 'export GOPATH=~/Development/go' >> ~/.zshrc
source ~/.zshrc
```
* Git clone the repo in $GOPATH/src/booking
```
cd $GOPATH
mkdir src
git clone git@github.com:Skedaddle/booking.git
```

## Dependency Management

* Install the booking dependencies using dep
```
cd $GOPATH/src/booking
dep init -v
```

## Running the container locally without Docker

```
cd $GOPATH/src/github.com/skedaddle/booking && APP_PORT=:5000 MONGO_URL=mongodb://skedaddle:skedaddle@localhost:27017 MONGO_DB_NAME=skedaddle_web MONGO_DB_BOOKINGS_COLLECTION=bookings MONGO_DB_LOGS_COLLECTION=logs go run main.go
```

## Running the container locally using Docker
```
cd $GOPATH/src/github.com/skedaddle/booking
./bin/run.sh
```
