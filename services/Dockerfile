FROM golang:1.13.4 as builder
WORKDIR /go/src
COPY . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o server main.go

FROM scratch
WORKDIR /go/bin
COPY --from=builder /go/src .
CMD ["/go/bin/server"]