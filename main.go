package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
}

func validateChoice(choice int) (bool, error) {
	if choice >= 1 && choice <= 5 {
		return true, nil
	}
	return false, fmt.Errorf("Invalid choice: %v", choice)
}

func toggleComplete(tasks *[]Task, id int) error {
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Completed = !(*tasks)[i].Completed
			return nil
		}
	}

	return fmt.Errorf("Task with ID %d not found", id)
}

func deleteTask(tasks *[]Task, id int) error {
	index := -1

	for i, task := range *tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Task with ID %d not found", id)
	}

	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)

	for i := range *tasks {
		(*tasks)[i].ID = i + 1
	}

	return nil
}

func main() {
	nextID := 1
	tasks := []Task{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Task Manager!\n\nPlease choose an option:\n1. Add Task\n2. View Tasks\n3. Toggle Task Completion\n4. Delete Task\n5. Exit")

	for {
		var input string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&input)

		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		} else if _, err := validateChoice(choice); err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			task := Task{nextID, description, false}
			tasks = append(tasks, task)
			fmt.Printf("Task added successfully with ID %d.\n", nextID)
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
		case 3:
			var id int

			for {
				fmt.Print("Enter task ID to toggle completion: ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				id, err = strconv.Atoi(input)
				if err != nil {
					fmt.Println("Please enter a valid number")
					continue
				}

				if err := toggleComplete(&tasks, id); err != nil {
					fmt.Println(err)
					continue
				} else {
					fmt.Printf("Task with ID %d toggled successfully.\n", id)
					break
				}

			}
		case 4:
			var id int

			for {
				fmt.Print("Enter task ID to delete: ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				id, err = strconv.Atoi(input)
				if err != nil {
					fmt.Println("Please enter a valid number")
					continue
				}

				if err := deleteTask(&tasks, id); err != nil {
					fmt.Println(err)
					continue
				} else {
					fmt.Printf("Task with ID %d deleted successfully.\n", id)
					nextID = len(tasks) + 1
					break
				}

			}
		case 5:
			fmt.Println("Exiting the Task Manager. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
