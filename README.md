# API Go Full

## Visão Geral

API Go Full é uma API RESTful para gerenciamento de produtos, desenvolvida em Go. Esta API fornece endpoints para criar, listar, atualizar e deletar produtos, além de gerenciar usuários e autenticação via JWT.

## Características

- Gerenciamento completo de produtos (CRUD)
- Autenticação de usuários
- Geração de token JWT
- Paginação na listagem de produtos
- Documentação Swagger

## Pré-requisitos

- Go (versão recomendada: 1.16+)
- PostgreSQL (ou o banco de dados de sua escolha)
- Utilizei o SQLITE3 pra testes primeiro

## Instalação

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/api-go-full.git
   ```

2. Entre no diretório do projeto:
   ```
   cd api-go-full
   ```

3. Instale as dependências:
   ```
   go mod tidy
   ```

4. Configure as variáveis de ambiente (crie um arquivo `.env` na raiz do projeto):
   ```
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   DB_NAME=nome_do_banco
   DB_HOST=localhost
   DB_PORT=5432
   JWT_SECRET=seu_segredo_jwt
   ```

5. Execute a aplicação:
   ```
   go run main.go
   ```

A API estará disponível em `http://localhost:8080`.

## Uso

### Autenticação

A maioria dos endpoints requer autenticação. Para autenticar, você precisa incluir o token JWT no header `Authorization` de suas requisições.

Para obter um token JWT:

1. Crie um usuário (se ainda não tiver um):
   ```
   POST /users
   {
     "name": "Seu Nome",
     "email": "seu@email.com",
     "password": "sua_senha"
   }
   ```

2. Gere um token JWT:
   ```
   POST /users/generate_jwt
   {
     "email": "seu@email.com",
     "password": "sua_senha"
   }
   ```

3. Use o token retornado no header `Authorization` de suas requisições:
   ```
   Authorization: Bearer seu_token_jwt
   ```

### Endpoints

- **Produtos**:
  - Listar produtos: `GET /products`
  - Criar produto: `POST /products`
  - Obter produto: `GET /products/{id}`
  - Atualizar produto: `PUT /products/{id}`
  - Deletar produto: `DELETE /products/{id}`

- **Usuários**:
  - Criar usuário: `POST /users`
  - Gerar JWT: `POST /users/generate_jwt`

Para mais detalhes sobre os endpoints, parâmetros e respostas, consulte a documentação Swagger disponível em `/swagger/index.html` quando a API estiver em execução.

## Desenvolvimento

### Estrutura do Projeto

```
api-go-full/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── entity/
│   ├── handlers/
│   ├── infra/
│   └── service/
├── docs/
│   └── swagger.json
├── go.mod
├── go.sum
└── README.md
```

### Executando Testes

Para executar os testes:

```
go test ./...
```

### Gerando Documentação Swagger

A documentação Swagger é gerada automaticamente. Para atualizar:

1. Instale o swag:
   ```
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

2. Gere a documentação:
   ```
   swag init -g cmd/server/main.go
   ```
