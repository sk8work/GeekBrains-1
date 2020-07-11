package fifo

var x []string

// Push - adding a new element into stack
func Push(str string) {
	x = append(x, str)
}

// Unshift - returns lust adding element
func Unshift() string {
	numOfElements := len(x)
	// When stack is empty, returns empty string
	if numOfElements == 0 {
		return "Нет элементов"
	}
	unshiftElem := x[0]
	x = x[1:numOfElements]
	return unshiftElem
}
