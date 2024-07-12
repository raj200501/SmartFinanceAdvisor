package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "SmartFinanceAdvisor/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/api/finance", handlers.GetFinanceData).Methods("GET")
    router.HandleFunc("/api/user", handlers.GetUserData).Methods("GET")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
