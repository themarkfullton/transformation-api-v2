FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build

EXPOSE 3000
CMD ["./main.go"]