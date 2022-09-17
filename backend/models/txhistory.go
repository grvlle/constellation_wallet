package models

// TXHistory stores inidividual transactions
type TXHistory struct {
	ID          uint   `json:"id"`
	Alias       string `json:"alias"`
	Amount      int64  `json:"amount"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Fee         int64  `json:"fee"`
	Timestamp   string `json:"timestamp"`
	Hash        string `json:"hash"`
	Parent      struct {
		Hash    string `json:"hash"`
		Ordinal int    `json:"ordinal"`
	} `json:"parent"`
	Status string `json:"status"`
	Failed bool
}
