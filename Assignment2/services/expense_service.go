package services

import (
	"database/sql"
	"errors"
	"modules/models"
	"modules/repositories"
)

func CreateExpense(expense models.Expense) (models.Expense, error) {
	if expense.Amount <= 0 {
		return expense, errors.New("Amount should be greater than 0")
	}
	if expense.Category == "" {
		return expense, errors.New("Category should not be empty")
	}
	if expense.UserID <= 0 {
		return expense, errors.New("User ID should be greater than 0")
	}
	return repositories.CreateExpense(expense)
}

func GetAllExpenses() ([]models.Expense, error) {
	return repositories.GetAllExpenses()
}

func GetExpenseByID(id int) (models.Expense, error) {
	if id <= 0 {
		return models.Expense{}, errors.New("Invalid expense ID")
	}
	return repositories.GetExpenseByID(id)
}

func UpdateExpense(id int, expense models.Expense) (models.Expense, error) {
	if id <= 0 {
		return models.Expense{}, errors.New("Invalid expense ID")
	}
	if expense.Amount <= 0 {
		return models.Expense{}, errors.New("Amount should be greater than 0")
	}
	if expense.Category == "" {
		return models.Expense{}, errors.New("Category should not be empty")
	}
	if expense.UserID <= 0 {
		return models.Expense{}, errors.New("User ID should be greater than 0")
	}

	updatedExpense, err := repositories.UpdateExpense(id, expense)
	if err == sql.ErrNoRows {
		return models.Expense{}, errors.New("Expense not found")
	}
	return updatedExpense, err
}

func DeleteExpense(id int) error {
	if id <= 0 {
		return errors.New("Invalid expense ID")
	}
	err := repositories.DeleteExpense(id)
	if err == sql.ErrNoRows {
		return errors.New("Expense not found")
	}
	return err
}

func GetExpensesByUserID(userID int) ([]models.Expense, error) {
	if userID <= 0 {
		return nil, errors.New("Invalid user ID")
	}
	return repositories.GetExpensesByUserID(userID)
}

func GetExpensesByCategory(category string) ([]models.Expense, error) {
	if category == "" {
		return nil, errors.New("Category should not be empty")
	}
	return repositories.GetExpensesByCategory(category)
}
