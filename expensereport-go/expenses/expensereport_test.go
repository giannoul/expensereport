package expenses

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestCharacterization(t *testing.T) {
	expensesInput := []Expense{
		{
			Type:   DINNER,
			Amount: 12,
		},
		{
			Type:   DINNER,
			Amount: 20,
		},
		{
			Type:   BREAKFAST,
			Amount: 10,
		},
		{
			Type:   CAR_RENTAL,
			Amount: 150,
		},
	}
	printReport(expensesInput)
	got := "the above returns no value"
	want := "?"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
