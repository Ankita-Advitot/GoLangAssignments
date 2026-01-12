package main 

import(
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)
const filepath="expense.json"

type Expense struct{
	Id int `json:"id"`
	Amount float64 `json:"amount"`
	Category string `json:category"`
	Description string `json:"description`
	Date string `json:"date"`
	
}

func loadExpenses()[]Expense{
	data,err:=os.ReadFile(filepath)
	if err!=nil{
		fmt.Println("Error while reading a file ",err)
	}

	var expenses []Expense
	json.Unmarshal(data,&expenses) // json to go struct format 
	return expenses
}

func saveExpenses(expenses []Expense){
	data,err:=json.MarshalIndent(expenses,""," ")
	if err!=nil{
		fmt.Println("Error exists while saving the expenses ", err)
	}
	os.WriteFile(filepath,data,0644)
	//0644 
	// 6 for read +write 110
	//4 fpr read  100
}

func addExpense(reader *bufio.Reader){
	expenses:=loadExpenses()
	fmt.Println("Enter amount")
	amt,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println("Error exists ",err)
		return
	}
	amount,_:=strconv.ParseFloat(strings.TrimSpace(amt), 64)
	if amount<=0{
		fmt.Println("enter valid amount")
		return 
	}
	fmt.Println("Enter category")

	category,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println("Error exists ",err)
		return
	}
	fmt.Println("Enter description")

	desc,err:=reader.ReadString('\n')
	if err!=nil{
		fmt.Println("Error exists ",err)
		return
	}
	id:=1
	if len(expenses)>0{
		id=expenses[len(expenses)-1].Id+1
	}
	expense := Expense{
		Id:          id,
		Amount:      amount,
		Category:    strings.TrimSpace(category),
		Description: strings.TrimSpace(desc),
		Date:        time.Now().Format("2006-01-02"),
	}
	expenses = append(expenses, expense)
	saveExpenses(expenses)

	fmt.Println("Expense added successfully!")
}
func viewExpenses() {
	expenses := loadExpenses()

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Println("\nID | Date       | Category | Amount | Description")
	fmt.Println("------------------------------------------------")

	for _, e := range expenses {
		fmt.Printf("%d  | %s | %s | %.2f | %s\n",
			e.Id, e.Date, e.Category, e.Amount, e.Description)
	}
}
func deleteExpense(reader *bufio.Reader) {
	expenses := loadExpenses()

	fmt.Print("Enter Expense ID to delete: ")
	idStr, err := reader.ReadString('\n')
	if err!=nil{
		fmt.Println("Error exists ",err)
		return
	}
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err!=nil{
		fmt.Println("Error exists ",err)
		return
	}

	var updated []Expense
	found := false

	for _, e := range expenses {
		if e.Id != id {
			updated = append(updated, e)
		} else {
			found = true
		}
	}

	if found {
		saveExpenses(updated)
		fmt.Println("Expense deleted.")
	} else {
		fmt.Println("Expense not found.")
	}
}
func main(){
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Personal Expense Tracker ---")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Delete Expense")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Println("Error exists ",err)
			return
		}
		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			addExpense(reader)
		case "2":
			viewExpenses()
		case "3":
			deleteExpense(reader)
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}

}