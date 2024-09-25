FROM golang:1.23

# Define o diretório de trabalho
WORKDIR /go/src/app

# Copia o arquivo .env para o contêiner
COPY .env ./


# Copia os arquivos do projeto para o contêiner
COPY . .

# Expõe a porta da aplicação
EXPOSE 8080

# Compila o binário da aplicação
RUN go build -o main cmd/server/main.go

# Executa o binário
CMD ["./main"]
