# Desafio Lab-Observabilidade

## Pré-requisitos

- Docker
- Go
- Crie um arquivo `.env` na raiz do projeto de acordo com o arquivo `.env.example` fornecido.

## Comandos Necessários para Executar o Sistema

### Execução Inicial

```bash
docker compose up --build -d
```

### Execuções Futuras

```bash
docker compose up -d
```

## Como Usar

- Execute o comando abaixo para consultar um CEP:

```bash
   curl -X POST http://localhost:8181 -H "Content-Type: application/json" -d '{"cep": "01153000"}'
   ```

- Acesse o Zipkin em `http://localhost:9411/zipkin` para visualizar os spans e verificar o tracing da requisição.
