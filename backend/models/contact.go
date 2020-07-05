package models

// Contact holds the DAG address and other contact details
type Contact struct {
	ID          uint   `json:"id" gorm:"AUTO_INCREMENT"`
	Alias       string `json:"alias"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}
