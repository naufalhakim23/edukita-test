################################################################################
# BASE
################################################################################
FROM golang:1.24-alpine AS base

RUN apk add --no-cache git openssh-client bash curl gcc g++ make libc6-compat git openssh-client ca-certificates vips vips-dev libc-dev libheif

################################################################################
# DEPENDENCY
################################################################################
FROM base AS dependency

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download -x && go mod verify

################################################################################
# BUILDER
################################################################################
FROM dependency AS builder

COPY . .

# build swagger file
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.0
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "static"' -tags "netgo" -installsuffix netgo -o edukita-teaching-grading -v

################################################################################
# WORKER
################################################################################
FROM alpine:latest

WORKDIR /usr/local/bin

RUN apk --update add ca-certificates

# copy docs (swagger API docs)
COPY --from=builder /build/docs /usr/local/bin/docs

COPY --from=builder /build/migrations  /usr/local/bin/migrations

COPY --from=builder /build/edukita-teaching-grading /usr/local/bin/edukita-teaching-grading


CMD ["./edukita-teaching-grading"]
