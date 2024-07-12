package models

type FinanceData struct {
    Balance  float64 `json:"balance"`
    Expenses float64 `json:"expenses"`
    Income   float64 `json:"income"`
}
