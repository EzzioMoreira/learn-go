package employee

// Golang não é orientado a objetos, mas é possível criar estruturas e métodos para simular o comportamento de classes e objetos.
type Person struct {
	fullName string
	age      int
	salary   int
}

// Utiliza o * para acessar o valor
func addSalaryPerson(p *Person, bonus int) {
	p.salary += bonus
}
