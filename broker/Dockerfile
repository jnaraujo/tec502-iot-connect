FROM golang:1.21.6 as builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia os arquivos .go para o diretório de trabalho
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux go build -o broker .

# Cria uma imagem alpine
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /root/

# Copia o binário do builder para o diretório de trabalho
COPY --from=builder /app/broker .

# Define a porta de exposição
EXPOSE 8080

# Define a variável de ambiente GIN_MODE para release
ENV GIN_MODE=release

# Define o comando de execução
CMD ["./broker"]