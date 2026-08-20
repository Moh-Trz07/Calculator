package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func calculator(nbr1, nbr2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return nbr1 + nbr2, nil
	case "-":
		return nbr1 - nbr2, nil
	case "*":
		return nbr1 * nbr2, nil
	case "/":
		if nbr2 <= 0 {
			return 0, errors.New("division by zero\n")
		}
		return nbr1 / nbr2, nil
	default:
		return 0, fmt.Errorf("invalid operator\n")
	}
}

func parseFloat(input string) (float64, error) {
	input = strings.TrimSpace(input)
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a Number", input)
	}
	return val, nil
}

func main() {
	fmt.Println("===[/*-+{Calculator}+-*/]===")
	for {
		var nbr1, nbr2 float64
		var ope string
		var x string
		fmt.Print("\nEnter First Nbr ('q' for quit) : ")
		fmt.Scan(&x)

		if x == "q" || x == "Q" {
			break
		}

		nbr1, err := parseFloat(x)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Print("Enter Operator :")
		fmt.Scan(&ope)
		fmt.Print("Enter Second Nbr :")
		fmt.Scan(&x)

		nbr2, err = parseFloat(x)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		res, err := calculator(nbr1, nbr2, ope)

		if err != nil {
			fmt.Printf("Error : %v", err)
			continue
		}

		fmt.Printf("%.2f %s %.2f = %.2f\n", nbr1, ope, nbr2, res)
	}

	fmt.Println("End.")
}
