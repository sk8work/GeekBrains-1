package main

import (
	"fmt"
	"runtime"
)

func somefunc() bool {
	return true
}
func main() {
	runtime.GOMAXPROCS(1)

	done := false

	done = somefunc()
	for !done {
		fmt.Println("Not done")
	}
	fmt.Println("finished")
}
