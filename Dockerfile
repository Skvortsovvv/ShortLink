FROM golang:1.17

ARG mode

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o shortLinks ./cmd/shortLinks/main.go

CMD ["./shortLinks -mode=$mode"]
