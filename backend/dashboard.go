package app

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
	updateIntervalCurrency = 50 // Seconds
	updateIntervalBlocks   = 5  // Seconds
	updateIntervalPieChart = 60 // Seconds
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
		retryCounter := 1
		for {
			select {
			case <-a.killSignal:
				return
			default:
				time.Sleep(time.Duration(retryCounter) * time.Second) // Incremental backoff
				for retryCounter <= 20 && a.wallet.Address != "" {

					balance, err := a.GetTokenBalance()
					if err != nil {
						retryCounter++
						break
					}
					a.log.Infoln("Current Balance: ", balance)
					a.wallet.Balance, a.wallet.AvailableBalance, a.wallet.TotalBalance = balance, balance, balance
					a.RT.Events.Emit("token", a.wallet.Balance, a.wallet.AvailableBalance, a.wallet.TotalBalance)
					UpdateCounter(updateIntervalToken, "token_counter", time.Second, a.RT)
					time.Sleep(updateIntervalToken * time.Second)
				}
			}
		}
	}()
}

// pricePoller polls the min-api.cryptocompare REST API for DAG token value.
// Once polled, it'll Emit the token value to Dashboard.vue for full token
// balance evaluation against USD.
func (a *WalletApplication) pricePoller() {

	const (
		apiKey string = "17b10afdddc411087e2140ec91bd73d88d0c20294541838b192255fc574b1cb7"
		ticker string = "DAG"
		url    string = "https://min-api.cryptocompare.com/data/pricemulti?fsyms=" + ticker + "&tsyms=BTC,USD,EUR&api_key=" + apiKey
	)

	go func() {
		retryCounter := 1
		time.Sleep(3 * time.Second) // Give some space to pollTokenBalance

		for {
			select {
			case <-a.killSignal:
				return
			default:
				a.wallet.TokenPrice.DAG.USD = 0
				a.wallet.TokenPrice.DAG.EUR = 0
				a.wallet.TokenPrice.DAG.BTC = 0
				time.Sleep(time.Duration(retryCounter) * time.Second) // Incremental backoff
				for retryCounter <= 20 && a.wallet.Balance != 0 {
					a.log.Infoln("Contacting token evaluation API on: " + url + ticker)

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

						a.log.Infoln("Contacting alternate token evaluation API")
						a.wallet.TokenPrice.DAG.USD, a.wallet.TokenPrice.DAG.BTC, err = getTokenPriceAlternateRoute()

						if err != nil {
							a.log.Errorln("Failed to fetch token metrics using alternate endpoint. Reason: ", err)
							if retryCounter == 10 || retryCounter == 15 || retryCounter == 20 {
								warn := fmt.Sprintf("No data recieved from Token Price API. Will try again in %v seconds.", retryCounter)
								a.sendWarning(warn)
							}
							retryCounter++
							break
						}
					}

					a.log.Infof("Collected token price in USD: %v", a.wallet.TokenPrice.DAG.USD)
					a.log.Infof("Collected token price in EUR: %v", a.wallet.TokenPrice.DAG.EUR)
					a.log.Infof("Collected token price in BTC: %v", a.wallet.TokenPrice.DAG.BTC)

					totalCurrencyBalance := 0.0
					if a.wallet.Currency == "USD" {
						totalCurrencyBalance = float64(a.wallet.Balance) * a.wallet.TokenPrice.DAG.USD
					} else if a.wallet.Currency == "EUR" {
						totalCurrencyBalance = float64(a.wallet.Balance) * a.wallet.TokenPrice.DAG.EUR
					} else if a.wallet.Currency == "BTC" {
						totalCurrencyBalance = float64(a.wallet.Balance) * a.wallet.TokenPrice.DAG.BTC
					}
					a.RT.Events.Emit("totalValue", a.wallet.Currency, totalCurrencyBalance)

					UpdateCounter(updateIntervalCurrency, "value_counter", time.Second, a.RT)
					time.Sleep(updateIntervalCurrency * time.Second)
				}
			}
		}
	}()
}

// StoreTermsOfServiceStateDB stores the Terms of Service state in the user DB
func (a *WalletApplication) StoreTermsOfServiceStateDB(termsOfService bool) bool {
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("TermsOfService", termsOfService).Error; err != nil {
		a.log.Errorln("Unable to store termsOfService state. Reason: ", err)
		a.sendError("Unable to store termsOfService state persistently. Reason: ", err)
		return false
	}
	return true
}

type tokenPriceAlt struct {
	Code string `json:"code"`
	Data struct {
		Sequence    string `json:"sequence"`
		BestAsk     string `json:"bestAsk"`
		Size        string `json:"size"`
		Price       string `json:"price"`
		BestBidSize string `json:"bestBidSize"`
		Time        int64  `json:"time"`
		BestBid     string `json:"bestBid"`
		BestAskSize string `json:"bestAskSize"`
	} `json:"data"`
}

// getTokenPriceAlternateRoute will kick in in case the main token poller API will return
// broken payload.
func getTokenPriceAlternateRoute() (float64, float64, error) {

	var tokenpriceUSD, tokenpriceBTC float64
	var err error

	tpa := new(tokenPriceAlt)

	const (
		url       = "https://api.kucoin.com/api/v1/market/orderbook/level1?symbol="
		usdTicker = "DAG-USDT"
		btcTicker = "DAG-BTC"
	)

	tickers := []string{usdTicker, btcTicker}

	for _, tick := range tickers {
		resp, err := http.Get(url + tick)
		if err != nil {
			return 0, 0, err
		}

		// Example resp: {"code":"200000","data":
		// {"sequence":"1583079038860","bestAsk":"0.01058",
		// "size":"73908.9903","price":"0.01058",
		// "bestBidSize":"403.157","time":1589605888009,
		// "bestBid":"0.010539","bestAskSize":"79091.0097"}}

		if resp == nil {
			return 0, 0, err
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, 0, err
		}

		err = json.Unmarshal(bodyBytes, &tpa)
		if err != nil {
			return 0, 0, err
		}

		s := tpa.Data.Price
		balance, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, 0, err
		}
		switch {
		case tick == usdTicker:
			tokenpriceUSD = balance
		case tick == btcTicker:
			tokenpriceBTC = balance
		}
	}

	return tokenpriceUSD, tokenpriceBTC, err
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
