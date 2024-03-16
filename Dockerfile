FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/api/main.go  

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY ./configs/api.toml  ./configs/api.toml
COPY ./configs/.env ./configs/.env  

CMD ["./app"]