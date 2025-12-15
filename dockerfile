FROM golang:1.25-alpine AS builder

WORKDIR /app

# dependências
COPY go.mod go.sum ./
RUN go mod download

# código
COPY . .

# build do binário (usa cmd/api)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o api ./cmd/api

# imagem final
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/api /app/api

EXPOSE 8080

CMD ["/app/api"]