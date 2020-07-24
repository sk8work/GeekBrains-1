// В этом разделе я пробую некоторые штуки, не обращать внимания!

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Rectangle struct {
	A float64
	B float64
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (r Rectangle) Perimeter() float64 {
	return r.A + r.B
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

type GetPerimeter interface {
	Perimeter() float64
}
type Message struct {
	M string
}

// func AllPerimsSum(perims ...GetPerimeter) float64 {
// 	res := 0.0
// 	for _, perim := range perims {
// 		res += perim.Perimeter()
// 	}
// 	return res
// }

// with nil
func AllPerimsSum(perims ...GetPerimeter) float64 {
	res := 0.0
	for _, perim := range perims {
		if perim == nil {
			continue
		}
		res += perim.Perimeter()
	}
	return res
}

func main() {
	// firstRect := Rectangle{10.0, 5.0}
	// firstTriangle := Triangle{3, 4, 5}
	// fmt.Println(firstRect.Perimeter())
	// fmt.Println(firstTriangle.Perimeter())
	// res := AllPerimsSum(firstRect, firstTriangle)
	// fmt.Println(res)

	// sl1 := []int{2, 3, 1, 5, 4, 8, 5, 5, 2, 3, 1}
	// sort.Ints(sl1)
	// fmt.Println(sl1)
	// fmt.Println(os.Args[0])
	// pb := string(os.Args[0])
	// fmt.Println(pb)
	// b := []byte(pb)
	// fmt.Println(b)
	// bb, err := json.Marshal(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(bb)
	// fmt.Println(string(bb))
	// fmt.Println([]byte(string(bb)))

	// cc := string(bb)
	// fmt.Println(cc)
	// ccb := []byte(cc)
	// fmt.Println(ccb)
	// fmt.Println(string(ccb))
	// fmt.Println([]byte(string(ccb)))
	// m := Message{}
	// err = json.Unmarshal(ccb, &m)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(m)

	str := "Hello"
	fmt.Println(str)
	jsonStrToByte, err := json.Marshal(str)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(jsonStrToByte)

	m := Message{}
	err = json.Unmarshal(jsonStrToByte, &m)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(m)

}
