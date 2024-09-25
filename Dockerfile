FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia o arquivo go.mod e go.sum (se existirem)
COPY go.mod go.sum* ./

# Baixa as dependências
RUN go mod download

# Copia o restante dos arquivos do projeto
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Imagem final
FROM alpine:latest

WORKDIR /app

# Copia o binário da imagem de compilação
COPY --from=builder /app/main .

# Copia o arquivo .env se necessário
COPY .env .

# Expõe a porta da aplicação
EXPOSE 8080

# Executa o binário
CMD ["./main"]