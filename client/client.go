package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	f, err := os.OpenFile("cotacoes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var cotacao CotacaoResponse
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(f, "%s\n", cotacao.Bid)
	if err != nil {
		panic(fmt.Errorf("erro ao salvar cotação no arquivo: %w", err))
	}
	fmt.Println("Cotação: ", cotacao.Bid)
}
