FROM golang:1.11

WORKDIR /go/src/app
COPY . .

RUN wget -P /tmp/dep https://raw.githubusercontent.com/golang/dep/master/install.sh
RUN chmod +rwx /tmp/dep/install.sh
RUN /tmp/dep/install.sh
RUN rm /tmp/dep/install.sh
RUN dep ensure -vendor-only

CMD go run main.go