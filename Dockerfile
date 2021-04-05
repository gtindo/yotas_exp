FROM golang:1.16-buster

LABEL maintainer="GUEKENG TINDO Yvan" version="1.0"

WORKDIR $GOPATH/src/github.com/gtindo/yotas_exp

RUN apt-get update
RUN apt-get install build-essential -y

COPY . .

RUN go get -d -v ./... 

RUN go get -u github.com/pressly/goose/cmd/goose

RUN make build

CMD make goose-up && make run
