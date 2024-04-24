FROM golang:1.22-alpine AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /build/app .

EXPOSE 8080

CMD ["./app"]
