package expenses

import (
	"fmt"
	"time"
)

type Type int

const (
	DINNER Type = iota + 1
	BREAKFAST
	CAR_RENTAL
)

type Expense struct {
	Type   Type
	Amount int
}

func produceStringTitle() string {
	return fmt.Sprintf("Expenses %s\n", time.Now().Format("2006-01-02"))
}

func isMeal(t Type) bool {
	mealTypes := []Type{DINNER, BREAKFAST}
	for _, m := range mealTypes {
		if m == t {
			return true
		}
	}
	return false
}

func mealAmountFilter(e Expense) (i int) {
	if isMeal(e.Type) {
		return e.Amount
	}
	return 0
}

func expenseNameMapper(i Type) (s string) {
	names := []string{"Dinner", "Breakfast", "Car Rental"}
	return names[i-1]
}

func mealOverExpensesIdentifier(e Expense) (s string) {
	m := make(map[Type]int)
	m[DINNER] = 5000
	m[BREAKFAST] = 1000
	if _, ok := m[e.Type]; ok {
		if e.Amount > m[e.Type] {
			return "X"
		}
	}
	return " "
}

func expenseLineText(expenseName, mealOverExpensesMarker string, amount int) string {
	return fmt.Sprintf("%s\t%d\t%s\n", expenseName, amount, mealOverExpensesMarker)
}

func mealExpensesLineText(mealExpenses int) string {
	return fmt.Sprintf("Meal expenses: %d\n", mealExpenses)
}

func totalExpensesLineText(total int) string {
	return fmt.Sprintf("Total expenses: %d\n", total)
}

func printReport(expenses []Expense) {
	var total, mealExpenses int

	fmt.Printf(produceStringTitle())

	for _, expense := range expenses {
		mealExpenses += mealAmountFilter(expense)
		expenseName := expenseNameMapper(expense.Type)
		mealOverExpensesMarker := mealOverExpensesIdentifier(expense)
		fmt.Printf(expenseLineText(expenseName, mealOverExpensesMarker, expense.Amount))
		total += expense.Amount
	}

	fmt.Printf(mealExpensesLineText(mealExpenses))
	fmt.Printf(totalExpensesLineText(total))
}
