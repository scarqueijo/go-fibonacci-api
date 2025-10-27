# -------------------------------
# Etapa 1: Build da aplicação Go
# -------------------------------
FROM golang:1.24 AS builder

WORKDIR /app

# Copiar ficheiros de dependências e descarregar módulos
COPY go.mod ./
RUN go mod download

# Copiar o código-fonte
COPY . .

# Compilar o binário de forma ESTÁTICA (sem dependências glibc)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o fibonacci-api .

# -------------------------------
# Etapa 2: Imagem final minimalista
# -------------------------------
FROM scratch

# Copiar apenas o binário estático compilado
COPY --from=builder /app/fibonacci-api /fibonacci-api

# Expor a porta usada pelo servidor Go
EXPOSE 8080

# Comando de arranque
ENTRYPOINT ["/fibonacci-api"]
