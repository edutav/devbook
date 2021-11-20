#!/bin/sh

# removendo versão anterior
rm -rf bin

go build -o bin/devbook-api

# copiando dockerfile para o diretório bin
cp Dockerfile .env ./bin

echo "******************************"
echo "| Binário criado com sucesso |"
echo "******************************"
