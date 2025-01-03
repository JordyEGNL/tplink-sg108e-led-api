FROM golang:1.23

ENV GIN_MODE=release

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

COPY ./src/*.go ./

RUN go build -o /tplink-sg108e-led-api

CMD ["/tplink-sg108e-led-api"]