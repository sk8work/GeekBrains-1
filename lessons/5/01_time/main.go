package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Format(time.RFC3339))
	// fmt.Println(time.Now().Format(time.RFC822Z))

	// fmt.Println("Hello")
	// time.Sleep(time.Second * 4)
	// fmt.Println("world")
	// t1 := time.Now()
	// time.Sleep(time.Second * 4)
	// fmt.Println(time.Since(t1))

	// s := "Hello world"
	// fmt.Println(strings.ToUpper(s))
	// fmt.Println(strings.ToLower(s))
	// fmt.Println(s)
	// b := []byte(s)
	// fmt.Println(b)
	// sl := make([]string, 0, 0)
	// fmt.Println(sl)
	// for _, val := range b {
	// 	sl = append(sl, string(val))
	// }
	// fmt.Println(sl)

	// str := "some string"
	// split := strings.Split(str, " ")
	// fmt.Println(split)
	// for _, v := range split {
	// 	fmt.Println(v)
	// }

	str := "some string"
	fmt.Println(strings.Compare(str, "some string 1"))
	fmt.Println(strings.Contains(str, "some"))
	fmt.Println(strings.Index(str, "str"))
	fmt.Println(strings.Count(str, "s"))
	fmt.Println(strings.Join([]string{"Hello", "World"}, ", "))
	fmt.Println(strings.Replace(str, "some", "televizor", -1))

}
