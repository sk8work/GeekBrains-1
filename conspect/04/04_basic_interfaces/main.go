package main

type Interface interface {
	// Len возвращает коичество элементов в коллекции
	Len() int
	// Less отдает true, если элементы i, j
	// необходимо поменять местами
	Less(i, j int) bool
	// Swap меняет местами элементы с индексами i, j
	Swap(i, j int)
}

type ToSort []int

func (t ToSort) Len() int           { return len(t) }
func (t ToSort) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ToSort) Less(i, j int) bool { return (t[i] < t[j]) }

func main() {

}
