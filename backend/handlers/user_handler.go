package handlers

import (
    "encoding/json"
    "net/http"

    "SmartFinanceAdvisor/models"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
    user := models.User{
        ID:    1,
        Name:  "John Doe",
        Email: "john.doe@example.com",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
