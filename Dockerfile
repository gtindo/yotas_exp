FROM golang:alpine

LABEL maintainer="GUEKENG TINDO Yvan" version="1.0"

WORKDIR $GOPATH/src/github.com/gtindo/yotas_exp

RUN sudo apt-get update
RUN sudo apt-get install build-essential

COPY . .

RUN go get -d -v ./...
RUN go install -v ./... 

RUN make build

CMD make goose-up && make run
