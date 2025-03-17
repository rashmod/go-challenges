package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	id   int
	text string
	done bool
}

type Todos struct {
	todos  map[int]*Todo
	nextId int
}

func (t *Todos) add(text string) {
	t.todos[t.nextId] = &Todo{id: t.nextId, text: text, done: false}
	t.nextId++
}

func (t *Todos) list() {
	if len(t.todos) == 0 {
		fmt.Println("no todos")
		return
	}

	fmt.Println("\ntodos:")
	for _, todo := range t.todos {
		status := "[ ]"
		if todo.done {
			status = "[x]"
		}
		fmt.Printf("%s %d %s\n", status, todo.id, todo.text)
	}
}

func (t *Todos) toggle(id int) {
	todo, ok := t.todos[id]
	if !ok {
		fmt.Printf("no todo with id %d\n", id)
		return
	}
	t.todos[id].done = !todo.done
}

func (t *Todos) clear() {
	t.todos = make(map[int]*Todo)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	todos := Todos{nextId: 1, todos: make(map[int]*Todo)}

	for {
		fmt.Println()
		fmt.Println("what would you like to do? (list, toggle, add, clear, exit, help)")
		fmt.Printf("> ")
		cmd, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "list":
			todos.list()

		case "toggle":
			fmt.Printf("> Enter todo id: ")
			id, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			i, err := strconv.Atoi(strings.TrimSpace(id))
			if err != nil {
				fmt.Println("invalid todo id")
				continue
			}

			todos.toggle(i)

		case "add":
			fmt.Printf("> Enter todo text: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			todos.add(strings.TrimSpace(text))

		case "help":
			fmt.Println("add - add a todo")
			fmt.Println("list - list all todos")
			fmt.Println("toggle - toggle a todo")
			fmt.Println("clear - clear all todos")
			fmt.Println("exit - exit the app")

		case "clear":
			todos.clear()
			fmt.Println("todos cleared")

		case "exit":
			fmt.Println("bye")
			os.Exit(0)

		default:
			fmt.Println("unknown command")
		}
	}
}
