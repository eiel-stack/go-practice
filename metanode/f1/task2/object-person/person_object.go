package main

import "fmt"

// Person represents a person with name and age.
type Person struct {
	Name string
	Age  int
}

// Employee represents an employee, composed of a Person and an EmployeeID.
type Employee struct {
	Person
	EmployeeID int
}

// PrintInfo prints the employee's information.
func (this Employee) PrintInfo() {
	fmt.Printf("Employee ID:%d, Name:%s, Age:%d\n", this.EmployeeID, this.Name, this.Age)
}

func main() {
	// Create an Employee instance using composition
	emp := Employee{
		Person: Person{
			Name: "Blues",
			Age:  33,
		},
		EmployeeID: 100001,
	}

	// Call the PrintInfo method
	emp.PrintInfo()
}
