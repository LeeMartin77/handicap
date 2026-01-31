FROM golang:1.25.6-alpine AS build

RUN mkdir /build
WORKDIR /app

COPY . .
RUN go build -o /build/server cmd/server/main.go

FROM alpine AS run

WORKDIR /app
COPY --from=build /build/server /app/server
RUN /app/server