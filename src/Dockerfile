FROM kinano/go

ENV SRC_DIR=${GOPATH}/src/github.com/kinano/go-mongo-docker

WORKDIR $SRC_DIR

# Copy source code
COPY . .
# Disable CGO in order to be able to use go modules
RUN go build -o ./build

ENTRYPOINT ["./build"]
