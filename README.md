# Sk-Next-API

Welcome to Skedaddle's Next Gen API (late fall 2018)


## Pre-requisites
* See [Docker pre-requisites](https://github.com/Skedaddle/infrastructure/tree/master/docker#pre-requisites)
* Define $GOPATH on your local box
```
# If you are using bash
echo 'export GOPATH=~/Development/go' >> ~/.bash_profile
source ~/.bash_profile
# If you are using zsh
echo 'export GOPATH=~/Development/go' >> ~/.zshrc
source ~/.zshrc
```
* Git clone the repo in $GOPATH/src/Sk-Next-API
```
cd $GOPATH
mkdir src/github/skedaddle/Sk-Next-API
git clone git@github.com:Skedaddle/Sk-Next-API.git
```

## Dependency Management

* Install the booking dependencies using dep
```
cd $GOPATH/src/github/skedaddle/Sk-Next-API
dep init -v
```

## Running the container locally using Docker
```
cd $GOPATH/src/github/skedaddle/Sk-Next-API
docker-compose up
```

## How do I shutdown the API locally?
```
cd $GOPATH/src/github/skedaddle/Sk-Next-API
docker-compose down
```

## How did we implement Hot Reloading on GoLang?

We are using realize to allow us to build the Go project on the fly when the developer saves Go files on the host machine.
See the `.realize.yml` file for the config we're using and also look at the GitHub repo: https://github.com/oxequa/realize
