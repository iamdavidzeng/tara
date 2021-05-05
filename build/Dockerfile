FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main cmd/tara/main.go

#

FROM scratch

COPY configs/config.yaml /app/configs/

COPY --from=builder /build/main /app/

WORKDIR /app

CMD ["./main"]
