package _714220044

import (
	"fmt"
	"testing"
)

func TestInsertEmployee(t *testing.T) {
	name := "Jimmy"
	position := "Software Engineer"
	salary := 5000000.0

	hasil := InsertEmployee(name, position, salary) 
	fmt.Println(hasil)

}

func TestInsertHonor(t *testing.T) {
	name := "Jimmy"
	amount := 25000000.0

	hasil := InsertHonor(name, amount)
	fmt.Println(hasil)
}

func TestInsertSalary(t *testing.T) {
    name := Employee{ 
        Name:     "jimmy",
        Position: "Sofware Engineer",
        Salary:   5000000.0,
    }
    month := 6
    year := 2023

    hasil := InsertSalary(name, month, year)
    fmt.Println(hasil)
}