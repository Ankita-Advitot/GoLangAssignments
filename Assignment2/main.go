package main

import (
	"net/http"

	"modules/config"
	"modules/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDB()

	r := mux.NewRouter()

	// Expense routes
	r.HandleFunc("/api/expenses", controllers.Create).Methods("POST")
	r.HandleFunc("/api/expenses", controllers.Index).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.GetExpenseByID).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.Update).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", controllers.Destroy).Methods("DELETE")

	// Filter routes
	r.HandleFunc("/api/expenses/user/{user_id}", controllers.GetExpensesByUserID).Methods("GET")
	r.HandleFunc("/api/expenses/category/{category}", controllers.GetExpensesByCategory).Methods("GET")

	http.ListenAndServe(":8080", r)
}
