package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go_introduction/commands"
	"github.com/go_introduction/expenses"
)

func main() {

	var exps []float64

	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}
		if input == "cls" {
			break
		}

		exp, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}

		exps = append(exps, exp)

	}

	fmt.Println(expenses.SumTotal(exps...))
	fmt.Println(expenses.Max(exps...))
	fmt.Println(expenses.Min(exps...))

	fmt.Println("Closing...")
}
