package circle

type Position struct {
	X int
	Y int
}

// Структура position представлена в виде поля другой структуры
type Circle struct {
	Point  Position
	Radius int
}
