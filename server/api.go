package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CotacaoRequest struct {
	USD struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type CotacaoResponse struct {
	Bid string `json:"bid"`
}

type API struct {
	database *Database
}

func NewAPI(database *Database) *API {
	return &API{database: database}
}
func (a *API) StartServer() {

	http.HandleFunc("/cotacao", a.handlerCotacao)
	fmt.Println("Server started on port 8080")
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		panic(error)
	}
}

func (a *API) handlerCotacao(w http.ResponseWriter, r *http.Request) {

	log.Println("Request received")
	bid, err := a.getCotacaoDolar()
	defer log.Println("request processed")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(CotacaoResponse{Bid: bid})

}

func (a *API) getCotacaoDolar() (string, error) {
	// Criar contexto com timeout para a requisição HTTP (1 segundo)
	httpCtx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(httpCtx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var cotacao CotacaoRequest
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		return "", err
	}
	_ = a.database.SaveCotacao(cotacao.USD.Bid)

	return cotacao.USD.Bid, nil
}
