FROM skedaddle/go

WORKDIR /go/src/app

# Copy the source code
COPY . .

# Compile dependencies using dep
RUN dep ensure -vendor-only

CMD go run main.go