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
	nextID := 1
	tasks := []Task{}

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

		switch choice {
		case 1:
			var description string
			fmt.Print("Enter task description: ")
			fmt.Scan(&description)

			task := Task{nextID, description, false}
			tasks = append(tasks, task)
			fmt.Printf("Task added: %v\n", task)
			nextID++
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				fmt.Println("Tasks:")
				for _, task := range tasks {
					fmt.Printf("ID: %d, Description: %s, Completed: %t\n", task.ID, task.Description, task.Completed)
				}
			}

		}
	}
}
