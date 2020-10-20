package models

// TXHistory stores inidividual transactions
type TXHistory struct {
	ID                 uint   `json:"id"`
	Alias              string `json:"alias"`
	Amount             int64  `json:"amount"`
	Sender             string `json:"sender"`
	Receiver           string `json:"receiver"`
	Fee                int64  `json:"fee"`
	Timestamp					 string `json:"timestamp"`
	Hash               string `json:"hash"`
	LastTransactionRef struct {
		Hash    string `json:"prevHash"`
		Ordinal int    `json:"ordinal"`
	} `json:"lastTransactionRef"`
	TS     string `json:"date"`
	Status string `json:"status"`
	Failed bool
}
