# review-job 微服务镜像
FROM golang:1.26.9-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPROXY=${GOPROXY}

WORKDIR /build

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o review-job ./cmd/review-job

FROM alpine:3.22

COPY --from=builder /build/review-job /

ENTRYPOINT ["/review-job"]
