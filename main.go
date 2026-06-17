package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const fileName = "todos.json"

func loadTodos() []Todo {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error loading file::: ", err)
	}

	var todos []Todo
	err = json.Unmarshal(file, &todos)
	if err != nil {
		fmt.Println("todos::: 2 ", todos)
		return []Todo{}
	}

	fmt.Println("todos::: 1 ", todos)

	return todos
}

func saveTodos(todos []Todo) {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("Error encoding todos:::", err)
		return
	}

	err = os.WriteFile(fileName, data, 0o644)
	if err != nil {
		fmt.Println("Something went wrong::: ", err)
		return
	}
}

func addTodo(title string) {
	todos := loadTodos()
	fmt.Println("todos::: ", todos)
	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)
	saveTodos(todos)

	fmt.Printf("Added todo: %s\n", title)
}

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
		addTodo(title)

	case "list":
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
