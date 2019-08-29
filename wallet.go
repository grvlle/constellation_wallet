package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/wailsapp/wails"
)

// Wallet holds all wallet information.
type Wallet struct {
	RT         *wails.Runtime // Client/Server Event Bus
	Balance    int            `json:"balance"`
	Address    string         `json:"address"`
	TokenPrice struct {
		DAG struct {
			BTC float64 `json:"BTC"`
			USD float64 `json:"USD"`
			EUR float64 `json:"EUR"`
		} `json:"DAG"`
	}
}

func (w *Wallet) WailsInit(runtime *wails.Runtime) error {
	w.RT = runtime
	w.BlockAmount()
	w.TokenAmount()
	w.PricePoller()
	return nil
}

func (w *Wallet) TokenAmount() {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(3000000000)
			w.RT.Events.Emit("token", randomNumber)
			time.Sleep(20 * time.Second)
		}
	}()
}

func (w *Wallet) BlockAmount() {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(300)
			w.RT.Events.Emit("blocks", randomNumber)
			time.Sleep(2 * time.Second)
		}
	}()
}

// PricePoller polls the min-api.cryptocompare REST API for DAG token value.
// Once polled, it'll Emit the token value to Dashboard.vue for full token
// balance evaluation against USD.
func (w *Wallet) PricePoller() {

	const (
		apiKey string = "17b10afdddc411087e2140ec91bd73d88d0c20294541838b192255fc574b1cb7"
		ticker string = "DAG"
		url    string = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=" + ticker + "&tsyms=BTC,USD,EUR&api_key=" + apiKey
	)

	go func() {
		for {
			resp, err := http.Get(url)
			if err != nil {
				// handle error
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			json.Unmarshal([]byte(body), &w.TokenPrice)

			w.RT.Events.Emit("price", "$", w.TokenPrice.DAG.USD)

			time.Sleep(20 * time.Second)
		}
	}()
}
