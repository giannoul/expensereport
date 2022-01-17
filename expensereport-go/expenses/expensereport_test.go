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

/*
Test: produceStringTitle
*/
type produceStringTitleTest struct {
	arg1, expected string
}

var produceStringTitleTests = []produceStringTitleTest{
	{"whatever", "Expenses whatever\n"},
	{"2022-01-17", "Expenses 2022-01-17\n"},
}

func Test_produceStringTitle(t *testing.T) {

	for _, test := range produceStringTitleTests {
		if output := produceStringTitle(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

/*
Test: isMeal
*/
type isMealTest struct {
	arg1     Type
	expected bool
}

var isMealTests = []isMealTest{
	{DINNER, true},
	{BREAKFAST, true},
	{CAR_RENTAL, false},
}

func Test_isMeal(t *testing.T) {

	for _, test := range isMealTests {
		if output := isMeal(test.arg1); output != test.expected {
			t.Errorf("Output %t not equal to expected %t for %d", output, test.expected, test.arg1)
		}
	}
}

/*
Test: mealAmountFilter
*/
type mealAmountFilterTest struct {
	arg1     Expense
	expected int
}

var mealAmountFilterTests = []mealAmountFilterTest{
	{Expense{DINNER, 20}, 20},
	{Expense{BREAKFAST, 40}, 40},
	{Expense{CAR_RENTAL, 30}, 0},
}

func Test_mealAmountFilter(t *testing.T) {

	for _, test := range mealAmountFilterTests {
		if output := mealAmountFilter(test.arg1); output != test.expected {
			t.Errorf("Output %d not equal to expected %d for %d", output, test.expected, test.arg1.Type)
		}
	}
}
