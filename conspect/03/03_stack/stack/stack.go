package stack

var x []string

// Push добавит новый элемент в стек
func Push(str string) {
	x = append(x, str)
}

// Pop
func Pop() string {
	numOfElems := len(x)
	// Когда стек будет пустым, он вернет пустую строку
	if numOfElems == 0 {
		return ""
	}
	popElem := x[numOfElems-1]
	x = x[:numOfElems-1]
	return popElem
}
