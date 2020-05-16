package models

// Path carries the file paths
type Path struct {
	ID          uint   `json:"id"`
	Alias       string `json:"alias"`
	LastTXFile  string
	PrevTXFile  string
	EmptyTXFile string
}
