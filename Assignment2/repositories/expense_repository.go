package repositories

import (
	"database/sql"
	"modules/config"
	"modules/models"
)

func CreateExpense(expense models.Expense) (models.Expense, error) {
	// returning id returns the last affected id 
	query := `INSERT INTO expenses (amount, category, description, date, user_id) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	err := config.DB.QueryRow(
		query,
		expense.Amount,
		expense.Category,
		expense.Description,
		expense.Date,
		expense.UserID,
	).Scan(&expense.ID)
	return expense, err
}

func GetAllExpenses() ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, user_id FROM expenses ORDER BY id DESC`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Amount,
			&expense.Category,
			&expense.Description,
			&expense.Date,
			&expense.UserID,
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, rows.Err()
}

func GetExpenseByID(id int) (models.Expense, error) {
	query := `SELECT id, amount, category, description, date, user_id FROM expenses WHERE id = $1`

	var expense models.Expense
	err := config.DB.QueryRow(query, id).Scan(
		&expense.ID,
		&expense.Amount,
		&expense.Category,
		&expense.Description,
		&expense.Date,
		&expense.UserID,
	)

	if err == sql.ErrNoRows {
		return expense, sql.ErrNoRows
	}
	return expense, err
}

func UpdateExpense(id int, expense models.Expense) (models.Expense, error) {
	query := `UPDATE expenses SET amount = $1, category = $2, description = $3, date = $4, user_id = $5 WHERE id = $6 RETURNING id, amount, category, description, date, user_id`

	var updatedExpense models.Expense
	err := config.DB.QueryRow(
		query,
		expense.Amount,
		expense.Category,
		expense.Description,
		expense.Date,
		expense.UserID,
		id,
	).Scan(
		&updatedExpense.ID,
		&updatedExpense.Amount,
		&updatedExpense.Category,
		&updatedExpense.Description,
		&updatedExpense.Date,
		&updatedExpense.UserID,
	)

	if err == sql.ErrNoRows {
		return updatedExpense, sql.ErrNoRows
	}
	return updatedExpense, err
}

func DeleteExpense(id int) error {
	query := `DELETE FROM expenses WHERE id = $1`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func GetExpensesByUserID(userID int) ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, user_id FROM expenses WHERE user_id = $1 ORDER BY id DESC`

	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Amount,
			&expense.Category,
			&expense.Description,
			&expense.Date,
			&expense.UserID,
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, rows.Err()
}

func GetExpensesByCategory(category string) ([]models.Expense, error) {
	query := `SELECT id, amount, category, description, date, user_id FROM expenses WHERE category = $1 ORDER BY id DESC`

	rows, err := config.DB.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Amount,
			&expense.Category,
			&expense.Description,
			&expense.Date,
			&expense.UserID,
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, rows.Err()
}
