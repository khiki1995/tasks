package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrMeasureNotFound = errors.New("measure not found")
var ErrIncorrectAmount = errors.New("amount is incorrect")
var ErrIncorrectAnswer = errors.New("answer is incorrect")

type System struct {
	Measure string
	Amount  float64
}

func main() {
	measures := GetDefaultMeasures()
	var all, answer string

	for {
		all, answer = "", ""
		fmt.Println("1. Enter 'add' to add new measure: ")
		fmt.Println("2. Enter measure to convert to another measure: ")
		for _, v := range measures {
			all += v.Measure + " | "
		}
		fmt.Println(all)

		fmt.Scan(&answer)

		switch answer {
		case "add":
			var name string
			var amount float64
			fmt.Print("Enter name of measure: ")
			fmt.Scan(&name)
			fmt.Printf("Enter how much meter in one %v: ", name)
			fmt.Scan(&amount)

			measures = append(measures, &System{Measure: name, Amount: amount})
		default:
			if !strings.Contains(all, answer) || answer == "" {
				fmt.Println(ErrMeasureNotFound)
				continue
			}
			Calculate(answer, measures)
		}
	}
}

func Calculate(firstM string, measures []*System) {
	var amount float64
	var answer, secondM string
	fmt.Printf("Enter amount of %v : ", firstM)
	fmt.Scan(&answer)
	amount, err := strconv.ParseFloat(answer, 64)
	if err != nil {
		fmt.Println(ErrIncorrectAmount)
		return
	}
	fmt.Print("Enter type of measure you whant to conver to: ")
	fmt.Scan(&secondM)

	for _, first := range measures {
		if first.Measure == firstM {
			for _, second := range measures {
				if second.Measure == secondM {
					amount = (amount * first.Amount) / second.Amount
					fmt.Printf("Answer : %v\n", amount)
					break
				}
			}
			break
		}
	}

}

func GetDefaultMeasures() []*System {
	return []*System{
		{
			Measure: "inch",
			Amount:  0.0254,
		},
		{
			Measure: "foot",
			Amount:  0.3048,
		},
		{
			Measure: "yard",
			Amount:  0.9144,
		},
		{
			Measure: "mile",
			Amount:  1609.34,
		},
		{
			Measure: "kilometer",
			Amount:  1000,
		},
		{
			Measure: "meter",
			Amount:  1,
		},
	}
}
