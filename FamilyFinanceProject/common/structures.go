package common

// User represents a user in the system
type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Relation  string  `json:"relation"`
	Income    float64 `json:"income"`
	Expenses  float64 `json:"expenses"`
	Balance   float64 `json:"balance"`
	Budget    float64 `json:"budget"`
	Bill      float64 `json:"bill"`
}

// Incoming represents an income transaction
type Incoming struct {
	Income_ID  int     `json:"income_id"`
	Date       string  `json:"date"`
	Account_ID int     `json:"id"`
	Vendor     string  `json:"vendor"`
	Category   string  `json:"category"`
	Deposit    float64 `json:"deposit"`
	Balance    float64 `json:"balance"`
	Notes      string  `json:"notes"`
}

// Expenses represents an expenses transaction
type Expenses struct {
	Expenses_ID int     `json:"expenses_id"`
	Date        string  `json:"date"`
	Account_ID  int     `json:"id"`
	Category    string  `json:"category"`
	Vendor      string  `json:"vendor"`
	Payment     float64 `json:"payment"`
	Balance     float64 `json:"balance"`
	Notes       string  `json:"notes"`
}
