# desafio-client-server

Projeto em Go com **server HTTP** e **client**:

- **Server**: expõe `GET /cotacao` na porta `8080`, consulta a cotação USD/BRL na AwesomeAPI e salva o valor em um banco SQLite (`cotacoes.db`).
- **Client**: consome `http://localhost:8080/cotacao` e grava o `bid` retornado em `cotacoes.txt`.

## Sobre o desafio (enunciado)

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos, banco de dados e manipulação de arquivos com Go.

Você precisará entregar dois sistemas em Go:

- `client.go`
- `server.go`

Requisitos do desafio:

- O `client.go` deve realizar uma requisição HTTP no `server.go` solicitando a cotação do dólar.
- O `server.go` deve consumir a API contendo o câmbio de Dólar e Real em `https://economia.awesomeapi.com.br/json/last/USD-BRL` e retornar no formato JSON o resultado para o cliente.
- Usando o package `context`, o `server.go` deve registrar no banco de dados SQLite cada cotação recebida:
  - timeout máximo para chamar a API de cotação do dólar: **200ms**
  - timeout máximo para persistir os dados no banco: **10ms**
- O `client.go` deve receber do `server.go` apenas o valor atual do câmbio (campo `bid` do JSON). Usando `context`, o `client.go` deve ter timeout máximo de **300ms** para receber o resultado do `server.go`.
- Os 3 contextos devem retornar erro nos logs caso o tempo de execução seja insuficiente.
- O `client.go` deve salvar a cotação atual em um arquivo `cotacao.txt` no formato: `Dólar: {valor}`.
- O endpoint do `server.go` deve ser `/cotacao` e a porta do servidor HTTP deve ser a **8080**.

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
