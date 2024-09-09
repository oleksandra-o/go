package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	// Sample data
	data := []Person{
		{Name: "Alice", Age: 25, Salary: 5000},
		{Name: "Bob", Age: 30, Salary: 6000},
		{Name: "Charlie", Age: 35, Salary: 7000},
	}

	// Create CSV file
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Write data to CSV file
	for _, person := range data {
		err := writer.Write([]string{person.Name, fmt.Sprintf("%d", person.Age), fmt.Sprintf("%.2f", person.Salary)})
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Read data from CSV file
	file, err = os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert CSV records to Person structs
	var persons []Person
	for _, record := range records {
		age := 0
		salary := 0.0

		fmt.Sscanf(record[1], "%d", &age)
		fmt.Sscanf(record[2], "%f", &salary)

		persons = append(persons, Person{Name: record[0], Age: age, Salary: salary})
	}

	// Sort persons by name
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	})

	// Display sorted data
	fmt.Println("Sorted Data:")
	for _, person := range persons {
		fmt.Printf("Name: %s, Age: %d, Salary: %.2f\n", person.Name, person.Age, person.Salary)
	}
}
