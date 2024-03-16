FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/api/main.go

FROM alpine:latest

COPY --from=builder /app/app /app/app
COPY configs/api.toml /configs/api.toml

CMD ["/app/app"]