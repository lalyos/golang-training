package main

import "fmt"

type Employee struct {
	ID        int
	Name      string
	Address   string
	Position  string
	Salary    int
	ManagerID int
}

type Point struct{ X, Y int }

func points() {
	p1 := Point{1, 1}
	p2 := Point{Y: 5}
	fmt.Println(p1)
	fmt.Println(p2)
}
func main() {
	var dilbert Employee
	dilbert = Employee{
		Name:      "dilbert",
		ID:        123,
		Salary:    1000,
		Position:  "DevOps",
		ManagerID: 1,
	}
	dilbert.Salary += 200
	fmt.Println(dilbert)
	employeeOfTheMonth := &dilbert
	employeeOfTheMonth.Position += " !"
	fmt.Println(dilbert)
	points()
}
