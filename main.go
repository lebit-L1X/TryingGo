package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type item struct {
	ID    string
	Name  string
	Price int
}

func main() {
	price, err := decimal.NewFromString("136.02")
	if err == nil {
		fmt.Print(price)
	}
}
