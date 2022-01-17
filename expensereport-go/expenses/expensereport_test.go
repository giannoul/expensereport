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

/*
Test: expenseNameMapper
*/
type expenseNameMapperTest struct {
	arg1     Type
	expected string
}

var expenseNameMapperTests = []expenseNameMapperTest{
	{DINNER, "Dinner"},
	{BREAKFAST, "Breakfast"},
	{CAR_RENTAL, "Car Rental"},
}

func Test_expenseNameMapper(t *testing.T) {

	for _, test := range expenseNameMapperTests {
		if output := expenseNameMapper(test.arg1); output != test.expected {
			t.Errorf("Output %s not equal to expected %s for %d", output, test.expected, test.arg1)
		}
	}
}

/*
Test: mealOverExpensesIdentifier
*/
type mealOverExpensesIdentifierTest struct {
	arg1     Expense
	expected string
}

var mealOverExpensesIdentifierTests = []mealOverExpensesIdentifierTest{
	{Expense{DINNER, 20}, " "},
	{Expense{DINNER, 2000}, " "},
	{Expense{DINNER, 5000}, " "},
	{Expense{DINNER, 5001}, "X"},
	{Expense{BREAKFAST, 40}, " "},
	{Expense{BREAKFAST, 999}, " "},
	{Expense{BREAKFAST, 1000}, " "},
	{Expense{BREAKFAST, 1001}, "X"},
	{Expense{CAR_RENTAL, 450000}, " "},
}

func Test_mealOverExpensesIdentifier(t *testing.T) {

	for _, test := range mealOverExpensesIdentifierTests {
		if output := mealOverExpensesIdentifier(test.arg1); output != test.expected {
			t.Errorf("Output %s not equal to expected %s for %d", output, test.expected, test.arg1)
		}
	}
}

/*
Test: expenseLineText
*/
type expenseLineTextTest struct {
	arg1, arg2 string
	arg3       int
	expected   string
}

var expenseLineTextTests = []expenseLineTextTest{
	{"Dinner", " ", 200, "Dinner\t200\t \n"},
	{"Breakfast", " ", 300, "Breakfast\t300\t \n"},
	{"Car Rental", " ", 400, "Car Rental\t400\t \n"},
}

func Test_expenseLineText(t *testing.T) {

	for _, test := range expenseLineTextTests {
		if output := expenseLineText(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}

/*
Test: mealExpensesLineText
*/
type mealExpensesLineTextTest struct {
	arg1     int
	expected string
}

var mealExpensesLineTextTests = []mealExpensesLineTextTest{
	{200, "Meal expenses: 200\n"},
	{600, "Meal expenses: 600\n"},
}

func Test_mealExpensesLineText(t *testing.T) {

	for _, test := range mealExpensesLineTextTests {
		if output := mealExpensesLineText(test.arg1); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}

/*
Test: totalExpensesLineText
*/
type totalExpensesLineTextTest struct {
	arg1     int
	expected string
}

var totalExpensesLineTextTests = []totalExpensesLineTextTest{
	{100, "Total expenses: 100\n"},
	{6500, "Total expenses: 6500\n"},
}

func Test_totalExpensesLineTextText(t *testing.T) {

	for _, test := range totalExpensesLineTextTests {
		if output := totalExpensesLineText(test.arg1); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}
