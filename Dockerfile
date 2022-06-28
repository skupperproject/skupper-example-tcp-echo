FROM golang:1.15-alpine AS builder
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build -o tcp-go-echo .


FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/tcp-go-echo /app
CMD ["/app/tcp-go-echo"]

