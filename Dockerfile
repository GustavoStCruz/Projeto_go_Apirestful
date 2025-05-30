FROM golang:1.24

WORKDIR /app

# Primeiro copie os arquivos de módulo
COPY go.mod go.sum ./

COPY . .

EXPOSE 8000

# Compile o projeto
RUN go build -o main .

# Comando para rodar
CMD ["./main"]