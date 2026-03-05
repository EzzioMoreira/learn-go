package main

import (
	"fmt"
	"os"
	"strconv"
)

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

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Arrays and Slices")
	salaryArrays := []int{1000, 2000, 3000} // array com 3 elementos
	//salaryArrays := make([]int, 3)            // é possível criar um array utilizando a função make
	// Array com tamanho fixo, não é possível adicionar ou remover elementos
	for i := 0; i < len(salaryArrays); i++ {
		salaryArrays[i] = 100 + i
	}
	for _, salaryArry := range salaryArrays {
		fmt.Println("Salary:", salaryArry)
	}
	// Slice sem tamanho fixo, é possível adicionar ou remover elementos utilizando a função append
	fmt.Println("")
	fmt.Println("Array with append")
	salaryAppend := []int{}
	for i := 0; i < 5; i++ {
		salaryAppend = append(salaryAppend, 100+i)
		//fmt.Println("Salary Append:", salaryAppend)
	}
	for _, SalasalaryAppend := range salaryAppend {
		fmt.Println("Salary Append:", SalasalaryAppend)
	}
	fmt.Println("")

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Maps")
	// Map é uma coleção de pares chave-valor, onde as chaves são únicas e os valores podem ser de qualquer tipo.
	personMap := make(map[string]int) // cria um map vazio
	personMap["Francisco"] = 30       // adiciona um par chave-valor ao map
	personMap["Maria"] = 25           // adiciona outro par chave-valor ao map
	for name, age := range personMap {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
	// Acessa o valor do map utilizando a chave e verifica se a chave existe
	sal, exists := personMap["Francisco"]
	fmt.Println("Salary Francisco:", sal, "Exists:", exists)

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Tratando Erros")
	// Em Go, os erros são tratados como valores, e a convenção é retornar um valor de erro como o último valor de uma função.
	if len(os.Args) != 2 { // verifica se o número de argumentos é diferente de 2
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1]) // converte o argumento para um inteiro utilizando a função Atoi do pacote strconv

	if err != nil {
		fmt.Println("Error converting string to int:", err)
		os.Exit(1)
	}
	fmt.Println("Converted number:", n)
	fmt.Println("")

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Defer")
	file, err := os.Open("file.txt") // abre um arquivo utilizando a função Open do pacote os
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // garante que o arquivo será fechado quando a função main terminar, mesmo que ocorra um erro

	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Print the pointer example")
	ponteiro()
	// Ponteiros são usados para passar referências a valores em vez de copiar os valores.
	var inteiro = 42
	var ponteiro *int = &inteiro // utiliza o operador & para obter o endereço de inteiro e atribui a ponteiro
	fmt.Println("Valor do inteiro:", inteiro)
	fmt.Println("Endereço do inteiro:", ponteiro)
	fmt.Println("Valor apontado por ponteiro:", *ponteiro) // utiliza o operador * para acessar o valor apontado por ponteiro
	fmt.Println("Exemplo 2: Ponteiros")
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
