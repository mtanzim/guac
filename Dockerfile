FROM golang:latest

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


RUN go build -o rest-server ./cmd/server/main.go
CMD ["./rest-server"]