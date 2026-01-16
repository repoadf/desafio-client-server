# desafio-client-server

Projeto em Go com **server HTTP** e **client**:

- **Server**: expõe `GET /cotacao` na porta `8080`, consulta a cotação USD/BRL na AwesomeAPI e salva o valor em um banco SQLite (`cotacoes.db`).
- **Client**: consome `http://localhost:8080/cotacao` e grava o `bid` retornado em `cotacoes.txt`.

## Requisitos

- Go instalado (veja a versão em `go.mod`)

## Como rodar o server

Em um terminal, a partir da raiz do projeto:

```bash
cd server
go run .
```

Saída esperada:

- `Server started on port 8080`

Observações:

- O arquivo `cotacoes.db` é criado/atualizado **no diretório atual** (por isso o `cd server`).
- O handler tem timeouts curtos (requisição externa e persistência), então dependendo da rede a chamada pode falhar com `500`.

## Como rodar o client

Em outro terminal:

```bash
cd client
go run .
```

O client:

- imprime o JSON retornado pelo server
- grava apenas o valor do `bid` (uma linha por execução) em `client/cotacoes.txt`

## Testando o endpoint manualmente

Com o server rodando:

```bash
curl -s http://localhost:8080/cotacao
```

Resposta esperada (exemplo):

```json
{"bid":"5.1234"}
```
