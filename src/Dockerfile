FROM skedaddle/go

ENV SRC_DIR=${GOPATH}/src/github.com/skedaddle/api

WORKDIR $SRC_DIR

# Copy dependencies
COPY Gopkg.toml Gopkg.lock ./
RUN mkdir -p $SRC_DIR/vendor

# Compile dependencies
RUN dep ensure -vendor-only -v

# Copy source code
COPY . .
RUN go build -o ./build

ENTRYPOINT ["./build"]
