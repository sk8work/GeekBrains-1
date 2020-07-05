package structs

// Person is a basictype
type Person struct {
	Name    string
	surname string
	// SetSurname()
}

// SetSurname - setter
func (p *Person) SetSurname(surname string) {
	p.surname = surname
}
