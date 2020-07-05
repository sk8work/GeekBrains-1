package main

import (
	"errors"
	"fmt"
)

func someFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[3])
}

func panicFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("Panic!!!")
}

func errorFunc() error {
	return errors.New("Some error occured")
}

func complexFunc(a int) (int, error) {
	if a == 0 {
		return 0, errors.New("a is a zero")
	}
	return a, nil
}

func main() {
	// someFunc()
	// fmt.Println("End of program")
	// panicFunc()
	res, err := complexFunc(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
