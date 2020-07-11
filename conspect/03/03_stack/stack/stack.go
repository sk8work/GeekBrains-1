package stack

// x := make([]string, 0, 0)
var x []string

// Push - adding a new element into stack
func Push(str string) {
	x = append(x, str)
}

// Pop - returns lust adding element
func Pop() string {
	numOfElements := len(x)
	// When stack is empty, returns empty string
	if numOfElements == 0 {
		return ""
	}
	popElem := x[numOfElements-1]
	x = x[:numOfElements-1]
	return popElem
}
