package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/wailsapp/wails"
)

const (
	dummyValue           = 300000
	updateIntervalToken  = 10
	updateIntervalBlocks = 5
)

// TokenAmount polls the token balance and stores it in the Wallet.Balance object
func (w *Wallet) TokenAmount(runtime *wails.Runtime) {
	go func() {
		for {
			w.Balance = rand.Intn(dummyValue)
			runtime.Events.Emit("token", w.Balance)
			w.UpdateTokenCounter(updateIntervalToken, runtime)
			time.Sleep(updateIntervalToken * time.Second)
		}
	}()
}

// RetrieveTokenAmount is a user initiated function for updating current balance
func (w *Wallet) RetrieveTokenAmount() int {
	w.Balance = rand.Intn(dummyValue)
	return w.Balance
}

// BlockAmount is a temporary function
func (w *Wallet) BlockAmount(runtime *wails.Runtime) {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(dummyValue)
			runtime.Events.Emit("blocks", randomNumber)
			w.UpdateBlockCounter(updateIntervalBlocks, runtime)
			time.Sleep(updateIntervalBlocks * time.Second)
		}
	}()
}

// PricePoller polls the min-api.cryptocompare REST API for DAG token value.
// Once polled, it'll Emit the token value to Dashboard.vue for full token
// balance evaluation against USD.
func (w *Wallet) PricePoller(runtime *wails.Runtime) {

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

			runtime.Events.Emit("price", "$", w.TokenPrice.DAG.USD)
			time.Sleep(updateIntervalToken * time.Second)
		}
	}()
}

// UpdateTokenCounter will count up from the last time a card was updated.
func (w *Wallet) UpdateTokenCounter(countFrom int, runtime *wails.Runtime) {
	go func() {
		for i := countFrom; i > 0; i-- {
			runtime.Events.Emit("counter", i)
			time.Sleep(time.Second)
			continue
		}
	}()
}

// UpdateBlockCounter will count up from the last time a card was updated.
func (w *Wallet) UpdateBlockCounter(countFrom int, runtime *wails.Runtime) {
	go func() {
		for i := countFrom; i > 0; i-- {
			runtime.Events.Emit("block_counter", i)
			time.Sleep(time.Second)
			continue
		}
	}()
}
