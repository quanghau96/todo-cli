package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		return []Todo{}
	}

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
	newTodo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)
	saveTodos(todos)

	fmt.Printf("Added todo: %s\n", title)
}

func listTodos() {
	todos := loadTodos()
	if len(todos) == 0 {
		fmt.Println("No todos found")
		return
	}

	for _, todo := range todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}

		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Title)
	}
}

func markDone(id int) {
	var todos []Todo
	todos = loadTodos()

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true

			fmt.Println("Marked as done:::", todos[i].Title)
			saveTodos(todos)

			return
		}
	}
}

func delete(id int) {
	var todos []Todo
	todos = loadTodos()

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos(todos)
			fmt.Println("Deleted::: ", todo.Title)
			return
		}
	}
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
		listTodos()
	case "complete":
		if len(args) < 3 {
			fmt.Println("Please provide the ID of the todo to complete")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Conversion error:", err)
		}

		markDone(id)
	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide the ID of the todo to delete")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Conversion error:", err)
		}

		delete(id)
	default:
		printHelp()
	}
}
