FROM golang:1.19.12-alpine as base
RUN apk add --no-cache curl
RUN apk add --no-cache git ca-certificates openssh-client

FROM base as builder
WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./

RUN --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go/pkg/mod \
     go mod download

COPY . ./
RUN go build -o backend .

FROM alpine:3.18.2
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/backend /backend
CMD ["/backend"]
