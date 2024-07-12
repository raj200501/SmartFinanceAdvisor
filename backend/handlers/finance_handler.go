package handlers

import (
    "encoding/json"
    "net/http"

    "SmartFinanceAdvisor/models"
)

func GetFinanceData(w http.ResponseWriter, r *http.Request) {
    financeData := models.FinanceData{
        Balance:  5000.75,
        Expenses: 3000.40,
        Income:   7000.00,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(financeData)
}
