# Expense Report Go

Kata for Coding Dojos.

## Code smells

1. complex if clauses: [example1](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L28), [example2](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L43)
2. [single block of code that has too much logic](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L27-L50)
3. extensive reference of constants

## Analyze what you would have to change to implement the new requirement without refactoring the code

1. [create a mapper function for this part](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L32-L40). This will make things easier to test and understand
2. [create a function based on this logic](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L42-L47). This will make things easier to test and understand
3. [create a function that will return the results as strings](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L53-L54). This will make things easier to test and understand


## Write a characterization test. Take note of all design smells that you missed that made your life writing a test miserable

The [lines here](https://github.com/giannoul/expensereport/blob/trunk/expensereport-go/expenses/expensereport.go#L53-L54) need to use `fmt.Printf` instead of `fmt.Println`. 

The function `printReport` has all the compute and output logic in it. This means that it makes it too hard to test. One possible way to do so would be something like [this](https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4), which in no way we are going to implement.

## How to run the tests

```
ilias@CSMachine:[~/GitCode/Code-Kata/expensereport/expensereport-go]: go test -v expenses/*
```