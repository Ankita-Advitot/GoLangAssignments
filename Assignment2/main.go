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
	r.HandleFunc("/api/expenses", controllers.CreateExpense).Methods("POST")
	r.HandleFunc("/api/expenses", controllers.GetAllExpenses).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.GetExpenseByID).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.UpdateExpense).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", controllers.DeleteExpense).Methods("DELETE")

	// Filter routes
	r.HandleFunc("/api/expenses/user/{user_id}", controllers.GetExpensesByUserID).Methods("GET")
	r.HandleFunc("/api/expenses/category/{category}", controllers.GetExpensesByCategory).Methods("GET")

	http.ListenAndServe(":8080", r)
}
