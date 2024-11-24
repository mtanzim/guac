FROM golang:bookworm as build
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o rest-server ./cmd/server/main.go

FROM chromedp/headless-shell:131.0.6778.86
WORKDIR /go/src/app
EXPOSE 8090
RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
COPY --from=build /go/src/app/rest-server .
COPY . .
CMD ["./rest-server"]