package main

import (
	"errors"
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Zero Division")
	}
	return a / b, nil
}

func main() {
	div1, err := div(1, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(div1)

	div2, err := div(1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(div2)
}
