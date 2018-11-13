
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

## Dependency Management

* Install the api dependencies using dep
```
dep init -v
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

## How did we implement Hot Reloading on GoLang?

https://github.com/oxequa/realize

## Deploy to EB
Upload the `Dockerrun.aws.json` to EB environment
