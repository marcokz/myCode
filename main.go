package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
	subcriber := magazine.Subscriber{Name: "Piter", Rate: 22, Active: true}
	fmt.Println(subcriber, subcriber.Active)

	var employee magazine.Employee
	employee.Name = "Alex"
	employee.Salary = 60000
	fmt.Println(employee.Name, employee.Salary)

	var subcriber2 magazine.Subscriber
	subcriber2.City = "Almaty"
	fmt.Println(subcriber2.Address.City)

	var sub3 magazine.Employee
	sub3.Name = "Erbol"
	sub3.Street = "Dauletkerey"

	fmt.Println(sub3.Name, sub3.Street)
}
