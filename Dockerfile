FROM golang:1.17 AS builder
ENV GO111MODULE=off \	
    CGO_ENABLED=0 \	
    GOOS=linux \	
    GOARCH=amd64
WORKDIR /build
COPY main.go .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 8000
ENTRYPOINT ["/httpserver"]