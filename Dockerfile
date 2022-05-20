FROM golang:1.17

ENV WORKMODE=memory

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get clean
RUN apt-get update
RUN apt-get -y install postgresql-client

# RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o shortlinks ./cmd/shortlinks/main.go

CMD ["./shortlinks"]
