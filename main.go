package main

import (
	"fmt"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const fileName = "todos.json"

func printHelp() {
	fmt.Printf("Todo CLI \n\n")
	fmt.Println("Usage:")
	fmt.Println("  add <title>       Add a new todo")
	fmt.Println("  list              List all todos")
	fmt.Println("  complete <id>     Mark a todo as completed")
	fmt.Println("  delete <id>       Delete a todo")
}

func main() {
	args := os.Args
	fmt.Printf("Args::: %d\n", len(args))

	if len(args) < 2 {
		fmt.Println("Please provide a task todo")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task todo")
			return
		}
		title := args[2]
		fmt.Printf("Adding todo: %s\n", title)
	case "list":
		fmt.Println("Listing all todos")
	case "complete":
		if len(args) < 3 {
			fmt.Println("Please provide the ID of the todo to complete")
			return
		}
		id := args[2]
		fmt.Printf("Completing todo with ID: %s\n", id)
	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide the ID of the todo to delete")
			return
		}
		id := args[2]
		fmt.Printf("Deleting todo with ID: %s\n", id)
	default:
		printHelp()
	}
}
