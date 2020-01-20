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
	dummyValue             = 1000
	updateIntervalToken    = 60 // Seconds
	updateIntervalBlocks   = 5  // Seconds
	updateIntervalPieChart = 6  // Seconds
)

// ChartData contains all the datapoints for the Charts
// on the Dashboard.
type ChartData struct {
	NodesOnline struct {
		Labels []string `json:"labels"`
		Series []int    `json:"series"`
	} `json:"nodes_online"`
	Transactions struct {
		Labels    []string `json:"labels"`
		SeriesOne []int    `json:"series_one"`
		SeriesTwo []int    `json:"series_two"`
	} `json:"transactions"`
	Throughput struct {
		Labels    []string `json:"labels"`
		SeriesOne []int    `json:"series_one"`
		SeriesTwo []int    `json:"series_two"`
	} `json:"throughput"`
}

// ChartDataInit initializes the ChartData struct with datapoints for
// the charts in the wallet. These are stored on the fs as chart_data.json
func (a *WalletApplication) ChartDataInit() *ChartData {
	cd := &ChartData{}

	cd.NodesOnline.Labels = []string{"30%", "20%", "50%"}
	cd.NodesOnline.Series = []int{30, 20, 50}

	cd.Transactions.Labels = []string{
		"Jan  ",
		"Feb  ",
		"Mar  ",
		"Apr  ",
		"Mai  ",
		"Jun  ",
		"Jul  ",
		"Aug  ",
		"Sep  ",
		"Oct  ",
		"Nov  ",
		"Dec  "}
	cd.Transactions.SeriesOne = []int{542, 543, 520, 680, 653, 753, 326, 434, 568, 610, 756, 895}
	cd.Transactions.SeriesTwo = []int{230, 293, 380, 480, 503, 553, 600, 664, 698, 710, 736, 795}

	cd.Throughput.Labels = []string{
		"9:00AM",
		"12:00AM",
		"3:00PM",
		"6:00PM",
		"9:00PM",
		"12:00PM",
		"3:00AM",
		"6:00AM"}
	cd.Throughput.SeriesOne = []int{287, 385, 490, 562, 594, 626, 698, 895, 952}
	cd.Throughput.SeriesTwo = []int{67, 152, 193, 240, 387, 435, 535, 642, 744}

	// Init chart widgets with data.
	go func() {
		for i := 0; i < 2; i++ {
			a.RT.Events.Emit("tx_stats", cd.Transactions.SeriesOne, cd.Transactions.SeriesTwo, cd.Transactions.Labels)
			a.RT.Events.Emit("node_stats", cd.NodesOnline.Series, cd.NodesOnline.Labels)
			a.RT.Events.Emit("network_stats", cd.Throughput.SeriesOne, cd.Throughput.SeriesTwo, cd.Throughput.Labels)
			time.Sleep(1 * time.Second)
		}
	}()

	return cd
}

// Populates the Nodes Online pie chart with data from the block explorer.
func (a *WalletApplication) nodeStats(cd *ChartData) {
	go func() {
		for {
			// Will populate the chart with random data
			for i := range cd.NodesOnline.Series {
				cd.NodesOnline.Series[i] = rand.Intn(dummyValue)
			}
			a.RT.Events.Emit("node_stats", cd.NodesOnline.Series, cd.NodesOnline.Labels)
			UpdateCounter(updateIntervalPieChart, "chart_counter", time.Second, a.RT)
			time.Sleep(updateIntervalPieChart * time.Second)
		}
	}()
}

func (a *WalletApplication) txStats(cd *ChartData) {
	go func() {
		for {
			// Will populate the chart with random data
			for i := range cd.Transactions.SeriesOne {
				cd.Transactions.SeriesOne[i] = rand.Intn(dummyValue)
			}
			for i := range cd.Transactions.SeriesTwo {
				cd.Transactions.SeriesTwo[i] = rand.Intn(dummyValue)
			}
			a.RT.Events.Emit("tx_stats", cd.Transactions.SeriesOne, cd.Transactions.SeriesTwo, cd.Transactions.Labels)
			//UpdateCounter(updateIntervalPieChart, "chart_counter", time.Second, a.RT)
			time.Sleep(updateIntervalPieChart * time.Second)
		}
	}()
}

func (a *WalletApplication) networkStats(cd *ChartData) {

	go func() {
		for {
			// Will populate the chart with random data
			for i := range cd.Throughput.SeriesOne {
				cd.Throughput.SeriesOne[i] = rand.Intn(dummyValue)
			}
			for i := range cd.Throughput.SeriesTwo {
				cd.Throughput.SeriesTwo[i] = rand.Intn(dummyValue)
			}
			a.RT.Events.Emit("network_stats", cd.Throughput.SeriesOne, cd.Throughput.SeriesTwo, cd.Throughput.Labels)
			//UpdateCounter(updateIntervalPieChart, "chart_counter", time.Second, a.RT)
			time.Sleep(updateIntervalPieChart * time.Second)
		}
	}()
}

// TokenAmount polls the token balance and stores it in the Wallet.Balance object
func (a *WalletApplication) tokenAmount(wallet *Wallet) {
	go func() {
		for {
			a.RT.Events.Emit("token", wallet.Balance)
			UpdateCounter(updateIntervalToken, "token_counter", time.Second, a.RT)
			time.Sleep(updateIntervalToken * time.Second)
		}
	}()
}

// BlockAmount is a temporary function
func (a *WalletApplication) blockAmount() {
	var randomNumber int
	go func() {
		for {
			randomNumber = rand.Intn(dummyValue)
			a.RT.Events.Emit("blocks", randomNumber)
			UpdateCounter(updateIntervalBlocks, "block_counter", time.Second, a.RT)
			time.Sleep(updateIntervalBlocks * time.Second)
		}
	}()
}

// PricePoller polls the min-api.cryptocompare REST API for DAG token value.
// Once polled, it'll Emit the token value to Dashboard.vue for full token
// balance evaluation against USD.
func (a *WalletApplication) pricePoller(wallet *Wallet) {

	const (
		apiKey string = "17b10afdddc411087e2140ec91bd73d88d0c20294541838b192255fc574b1cb7"
		ticker string = "DAG"
		url    string = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=" + ticker + "&tsyms=BTC,USD,EUR&api_key=" + apiKey
	)

	go func() {
		retryCounter := 0
		for {
			resp, err := http.Get(url)
			if err != nil {
				a.log.Warnf("Unable to poll token evaluation. Reason: ", err) // Log this
			}
			if resp != nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					a.sendError("Unable to read HTTP resonse from Token API. Reason: ", err)
					a.log.Warnf("Unable to read HTTP resonse from Token API. Reason: ", err)
				}
				err = json.Unmarshal([]byte(body), &wallet.TokenPrice)
				if err != nil {
					a.sendError("Unable to display token price. Reason: ", err)
					a.log.Warnf("Unable to display token price. Reason:", err)
				}
				a.log.Debugf("Collected token price in USD: %v", wallet.TokenPrice.DAG.USD)

				tokenUSD := int(float64(wallet.Balance) * wallet.TokenPrice.DAG.USD)
				a.RT.Events.Emit("price", "$", tokenUSD)
				time.Sleep(updateIntervalToken * time.Second)
			} else {
				retryCounter++
				time.Sleep(1 * time.Second)
				if retryCounter >= 10 {
					a.sendError("Unable to poll token evaluation. Reason: ", err)
					break
				}
			}
		}
	}()
}

// UpdateCounter will count up from the last time a card was updated.
func UpdateCounter(countFrom int, counter string, unit time.Duration, runtime *wails.Runtime) {
	go func() {
		for i := countFrom; i > 0; i-- {
			runtime.Events.Emit(counter, i)
			time.Sleep(unit)
			continue
		}
	}()
}
