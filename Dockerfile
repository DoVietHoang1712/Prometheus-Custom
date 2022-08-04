FROM golang:1.17-alpine AS build
WORKDIR /
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /app

FROM alpine:latest
WORKDIR /

COPY --from=build /app /app

ENTRYPOINT ["/app"]