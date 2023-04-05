FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /api

EXPOSE 8585
ENTRYPOINT ["/api"]
