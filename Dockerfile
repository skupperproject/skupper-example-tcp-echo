FROM --platform=$BUILDPLATFORM golang:1.15-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

RUN mkdir /build
ADD . /build
WORKDIR /build
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o tcp-go-echo .

FROM --platform=$TARGETPLATFORM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/tcp-go-echo /app
CMD ["/app/tcp-go-echo"]
