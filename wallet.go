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

// WailsInit initializes the Client and Server side bindings
func (w *Wallet) WailsInit(runtime *wails.Runtime) error {
	w.RT = runtime
	w.BlockAmount()
	w.TokenAmount()
	w.PricePoller()
	return nil
}

// TokenAmount polls the token balance and stores it in the Wallet.Balance object
func (w *Wallet) TokenAmount() {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(3000000)
			w.RT.Events.Emit("token", randomNumber)
			w.UpdateTokenCounter(20)
			time.Sleep(20 * time.Second)
		}
	}()
}

// BlockAmount is a temporary function
func (w *Wallet) BlockAmount() {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(300)
			w.RT.Events.Emit("blocks", randomNumber)
			w.UpdateBlockCounter(5)
			time.Sleep(5 * time.Second)
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

// UpdateTokenCounter will count up from the last time a card was updated.
func (w *Wallet) UpdateTokenCounter(countFrom int) {
	go func() {
		for i := countFrom; i > 0; i-- {
			w.RT.Events.Emit("counter", i)
			time.Sleep(time.Second)
			continue
		}
	}()
}

// UpdateBlockCounter will count up from the last time a card was updated.
func (w *Wallet) UpdateBlockCounter(countFrom int) {
	go func() {
		for i := countFrom; i > 0; i-- {
			w.RT.Events.Emit("block_counter", i)
			time.Sleep(time.Second)
			continue
		}
	}()
}
