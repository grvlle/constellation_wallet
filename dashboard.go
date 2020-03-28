package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/wailsapp/wails"
)

const (
	dummyValue             = 1000
	updateIntervalToken    = 30 // Seconds
	updateIntervalUSD      = 50 // Seconds
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
		"9AM",
		"12AM",
		"3PM",
		"6PM",
		"9PM",
		"12PM",
		"3AM",
		"6AM"}
	cd.Throughput.SeriesOne = []int{287, 385, 490, 562, 594, 626, 698, 895, 952}
	cd.Throughput.SeriesTwo = []int{67, 152, 193, 240, 387, 435, 535, 642, 744}

	// Init chart widgets with data.
	go func() {
		select {
		case <-a.killSignal:
			return
		default:
			for i := 0; i < 2; i++ {
				a.RT.Events.Emit("tx_stats", cd.Transactions.SeriesOne, cd.Transactions.SeriesTwo, cd.Transactions.Labels)
				a.RT.Events.Emit("node_stats", cd.NodesOnline.Series, cd.NodesOnline.Labels)
				a.RT.Events.Emit("network_stats", cd.Throughput.SeriesOne, cd.Throughput.SeriesTwo, cd.Throughput.Labels)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return cd
}

// Populates the Nodes Online pie chart with data from the block explorer.
func (a *WalletApplication) nodeStats(cd *ChartData) {
	go func() {
		for {
			select {
			case <-a.killSignal:
				return
			default:
				// Will populate the chart with random data
				for i := range cd.NodesOnline.Series {
					cd.NodesOnline.Series[i] = rand.Intn(dummyValue)
				}
				a.RT.Events.Emit("node_stats", cd.NodesOnline.Series, cd.NodesOnline.Labels)
				UpdateCounter(updateIntervalPieChart, "chart_counter", time.Second, a.RT)
				time.Sleep(updateIntervalPieChart * time.Second)
			}
		}
	}()
}

func (a *WalletApplication) txStats(cd *ChartData) {
	go func() {
		for {
			select {
			case <-a.killSignal:
				return
			default:
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
		}
	}()
}

func (a *WalletApplication) networkStats(cd *ChartData) {

	go func() {
		for {
			select {
			case <-a.killSignal:
				return
			default:
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
		}
	}()
}

// BlockAmount is a temporary function
func (a *WalletApplication) blockAmount() {
	var randomNumber int
	go func() {
		for {
			select {
			case <-a.killSignal:
				return
			default:
				randomNumber = rand.Intn(dummyValue)
				a.RT.Events.Emit("blocks", randomNumber)
				UpdateCounter(updateIntervalBlocks, "block_counter", time.Second, a.RT)
				time.Sleep(updateIntervalBlocks * time.Second)
			}
		}
	}()
}

func (a *WalletApplication) pollTokenBalance() {
	go func() {
		var retryCounter int
		for {
			select {
			case <-a.killSignal:
				return
			default:
				time.Sleep(time.Duration(retryCounter) * time.Second) // Incremental backoff
				for retryCounter <= 20 && a.wallet.Address != "" {

					a.log.Debug("Contacting mainnet on: " + a.Network.URL + a.Network.Handles.Balance + " Sending the following payload: " + a.wallet.Address)

					resp, err := http.Get(a.Network.URL + a.Network.Handles.Balance + a.wallet.Address)
					if err != nil {
						a.log.Errorln("Failed to send HTTP request. Reason: ", err)
						retryCounter++
						break
					}
					if resp == nil {
						retryCounter++
						a.log.Errorln("Killing pollTokenBalance after 10 failed attempts to get balance from mainnet, Reason: ", err)
						a.sendWarning("Unable to showcase current balance. Please check your internet connectivity and restart the wallet application.")
						break
					}
					defer resp.Body.Close()

					bodyBytes, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						retryCounter++
						a.log.Error("Unable to update token balance. Reason: ", err)
						break
					}
					s := string(bodyBytes)
					if s == "" {
						s = "0" // Empty means zero
					}
					i, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						retryCounter++
						a.log.Warnln("Unable to parse balance. Reason:", err)
						break
					}
					f := fmt.Sprintf("%.2f", float64(i)/1e8) // Reverse normalized float

					balance, err := strconv.ParseFloat(f, 64)
					if err != nil {
						retryCounter++
						a.log.Warnln("Unable to type cast string to float for token balance poller. Check your internet connectivity")
						break
					}

					a.log.Debugln("Current Balance: ", f)
					a.wallet.Balance, a.wallet.AvailableBalance, a.wallet.TotalBalance = balance, balance, balance
					a.RT.Events.Emit("token", a.wallet.Balance, a.wallet.AvailableBalance, a.wallet.TotalBalance)
					UpdateCounter(updateIntervalToken, "token_counter", time.Second, a.RT)
					time.Sleep(updateIntervalToken * time.Second)
				}
			}
		}
	}()
}

// PricePoller polls the min-api.cryptocompare REST API for DAG token value.
// Once polled, it'll Emit the token value to Dashboard.vue for full token
// balance evaluation against USD.
func (a *WalletApplication) pricePoller() {

	const (
		apiKey string = "17b10afdddc411087e2140ec91bd73d88d0c20294541838b192255fc574b1cb7"
		ticker string = "DAG"
		url    string = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=" + ticker + "&tsyms=BTC,USD,EUR&api_key=" + apiKey
	)

	go func() {
		var retryCounter int
		time.Sleep(3 * time.Second) // Give some space to pollTokenBalance

		for {
			select {
			case <-a.killSignal:
				return
			default:
				a.wallet.TokenPrice.DAG.USD = 0
				time.Sleep(time.Duration(retryCounter) * time.Second) // Incremental backoff
				for retryCounter <= 20 && a.wallet.Balance != 0 {
					a.log.Debug("Contacting token evaluation API on: " + url + ticker)

					resp, err := http.Get(url)
					if err != nil {
						a.log.Warnln("Unable to poll token evaluation. Reason: ", err) // Log this
						retryCounter++
						break
					}

					if resp == nil {
						retryCounter++
						a.log.Errorln("Killing pollTokenBalance after 10 failed attempts to get balance from mainnet, Reason: ", err)
						a.sendWarning("Unable to showcase token USD evaluation. Please check your internet connectivity and restart the wallet application.")
						break
					}

					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						retryCounter++
						a.sendError("Unable to read HTTP resonse from Token API. Reason: ", err)
						a.log.Warnln("Unable to read HTTP resonse from Token API. Reason: ", err)
						break
					}
					err = json.Unmarshal([]byte(body), &a.wallet.TokenPrice)
					if err != nil {
						retryCounter++
						a.sendError("Unable to display token price. Reason: ", err)
						a.log.Warnln("Unable to display token price. Reason:", err)
						break
					}

					if a.wallet.Balance != 0 && a.wallet.TokenPrice.DAG.USD == 0 {
						if retryCounter == 10 || retryCounter == 15 || retryCounter == 20 {
							warn := fmt.Sprintf("No data recieved from Token Price API. Will try again in %v seconds.", retryCounter)
							a.sendWarning(warn)
						}
						retryCounter++
						break
					}

					a.log.Debugf("Collected token price in USD: %v", a.wallet.TokenPrice.DAG.USD)

					tokenUSD := int(float64(a.wallet.Balance) * a.wallet.TokenPrice.DAG.USD)
					a.RT.Events.Emit("totalValue", "USD", tokenUSD)
					UpdateCounter(updateIntervalUSD, "value_counter", time.Second, a.RT)
					time.Sleep(updateIntervalUSD * time.Second)

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
