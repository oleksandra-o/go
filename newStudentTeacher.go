package main

import "fmt"

type Address struct {
	Street   string
	City     string
	State    string
	ZipCode  int
}

type Student struct {
	StudentID int
	Name      string
	Address   Address
}

type Teacher struct {
	EmployeeID int
	Name       string
	Salary     float64
	Address    Address
}

func populateStudent() Student {
	return Student{
		StudentID: 123,
		Name:      "John Doe",
		Address: Address{
			Street:  "123 Main Street",
			City:    "Cityville",
			State:   "Stateville",
			ZipCode: 12345,
		},
	}
}

func populateTeacher() Teacher {
	return Teacher{
		EmployeeID: 456,
		Name:       "Jane Smith",
		Salary:     5000.0,
		Address: Address{
			Street:  "456 Elm Street",
			City:    "Townville",
			State:   "Stateville",
			ZipCode: 67890,
		},
	}
}

func printStudent(student Student) {
	fmt.Println("Student ID:", student.StudentID)
	fmt.Println("Name:", student.Name)
	fmt.Println("Address:")
	fmt.Println("Street:", student.Address.Street)
	fmt.Println("City:", student.Address.City)
	fmt.Println("State:", student.Address.State)
	fmt.Println("Zip Code:", student.Address.ZipCode)
}

func printTeacher(teacher Teacher) {
	fmt.Println("Employee ID:", teacher.EmployeeID)
	fmt.Println("Name:", teacher.Name)
	fmt.Println("Salary:", teacher.Salary)
	fmt.Println("Address:")
	fmt.Println("Street:", teacher.Address.Street)
	fmt.Println("City:", teacher.Address.City)
	fmt.Println("State:", teacher.Address.State)
	fmt.Println("Zip Code:", teacher.Address.ZipCode)
}

func main() {
	student := populateStudent()
	teacher := populateTeacher()

	fmt.Println("Student Information:")
	printStudent(student)

	fmt.Println("\nTeacher Information:")
	printTeacher(teacher)
}
