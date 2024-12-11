package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func storeToFile(ebt, profit, ratio float64) error {
	// Format the float64 values into a string
	data := fmt.Sprintf("EBT: %.2f, Profit: %.2f, Ratio: %.2f\n", ebt, profit, ratio)

	// Write the string to a file
	err := os.WriteFile("financial_data.txt", []byte(data), 0644)
	if err != nil {
		return err
	}

	return nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	storeToFile(ebt, profit, ratio)
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("InvalidInput")
	}
	return userInput, nil
}
