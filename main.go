package main

import "fmt"

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func validateChoice(choice int) (bool, error) {
	if choice >= 1 && choice <= 4 {
		return true, nil
	}
	return false, fmt.Errorf("Invalid choice: %v", choice)
}

func main() {
	// tasks := []Task{}

	fmt.Println("Welcome to the Task Manager!\n\nPlease choose an option:\n1. Add Task\n2. View Tasks\n3. Mark Task as Completed\n4. Exit")

	for {
		var choice int
		fmt.Print("Enter your choice:")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Please enter a valid number")
			continue
		} else if _, err := validateChoice(choice); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
