FROM golang:1.22-alpine

WORKDIR /app

# Instalar git, tern e sqlc
RUN apk add --no-cache git bash \
  && go install github.com/jackc/tern@latest \
  && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

COPY . .

RUN go mod tidy 

COPY wait-for-it.sh /app/
RUN chmod +x /app/wait-for-it.sh

COPY wait-for-db.sh /app/
RUN chmod +x /app/wait-for-db.sh

# Construir o binário Go, mas não executar go generate ainda
RUN go build -o main ./cmd/wsrs/main.go

EXPOSE 8080

# Usar um script de entrada para rodar o go generate após o banco de dados estar disponível
CMD ["./wait-for-db.sh"]
