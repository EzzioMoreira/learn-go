package main

import "fmt"

// Golang não é orientado a objetos, mas é possível criar estruturas e métodos para simular o comportamento de classes e objetos.
type Person struct {
	fullName string
	age      int
	salary   int
}

func main() {
	person := &Person{ // utiliza o operador & para passar o endereço
		fullName: "Francisco da Silva",
		age:      30,
		salary:   1000,
	}

	salary := 1000
	firstName, lastName := getName("Francisco", "da Silva")
	newSalary, bonus := getSalary(salary, 320)

	fmt.Println("Hello,", firstName, lastName, "your salary is", newSalary, "and your bonus is", bonus)

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Print the pointer example")
	ponteiro()

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Print the person struct")
	fmt.Printf("Name: %s, Age: %d, Salary: %d\n", person.fullName, person.age, person.salary)
	fmt.Println(person)

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Adding bonus to salary appointing to the function addSalaryPerson")
	addSalaryPerson(person, 30)
	fmt.Println("Name:", person.fullName, "-", "Salary:", person.salary)

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Adding person struct to the function addSalaryPersonStruct")
	person.addSalaryPersonStruct(50)
	fmt.Println("Name:", person.fullName, "-", "Salary:", person.salary)
	fmt.Println("")
}

// Utiliza o * para acessar o valor
func addSalaryPerson(p *Person, bonus int) {
	p.salary += bonus
}

// Adiciona a função a estrutura Person e utiliza o valor da estrutura para acessar o campo salary
func (p *Person) addSalaryPersonStruct(bonus int) {
	p.salary += bonus
}

func getName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func getSalary(valueSalary int, bonus int) (int, int) {
	return valueSalary + bonus, bonus
}

func ponteiro() {
	// Ponteiros são úteis quando você deseja modificar o valor de uma variável dentro de uma função ou quando deseja economizar memória ao passar grandes estruturas de dados.
	x := 10
	p := &x // utiliza o operador & para obter o endereço de x e atribui a p
	// Imprime o valor de x, o endereço de x e o valor apontado por p
	fmt.Println("Valor de x:", x)
	// Imprime o endereço de x e o valor apontado por p
	fmt.Println("Endereço de x:", p)
	// Imprime o valor apontado por p
	fmt.Println("Valor apontado por p:", *p) // utiliza o operador * para acessar o valor apontado por p
}
