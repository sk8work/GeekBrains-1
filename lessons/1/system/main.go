package main

import (
	"fmt"
	"runtime"
)

func main() {
	myOs := runtime.GOOS
	myProcs := runtime.GOMAXPROCS(10)
	fmt.Println("My OS is:", myOs)
	fmt.Println("My processors number is:", myProcs)
}
