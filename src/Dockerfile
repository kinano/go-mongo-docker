FROM kinano/go

ENV SRC_DIR=${GOPATH}/src/github.com/kinano/go-mongo-docker

WORKDIR $SRC_DIR

# Copy source code
COPY . .

RUN go build -o ./build/compiled_app

ENTRYPOINT ["./build/compiled_app"]
