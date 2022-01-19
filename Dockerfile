FROM golang:1.14.4 as builder

WORKDIR /src/
COPY . /src/

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -tags netgo -mod vendor -o ./dapr-go-subscriber .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /src/dapr-go-subscriber .

ENTRYPOINT ["./dapr-go-subscriber"]
