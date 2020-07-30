package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	someInt  = 5
	someInt2 = flag.Int("some_int", 42, "hello world!")
	needHelp = flag.Bool("help", false, "Need help?")
	dur      = flag.Duration("some_duration", time.Second, "Duration")
)

func main() {
	flag.Parse()
	if *needHelp {
		flag.Usage()
		return
	}
	fmt.Println(someInt)
	fmt.Println(*someInt2)
	fmt.Println(*dur)
}
