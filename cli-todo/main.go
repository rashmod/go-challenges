package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type Todos struct {
	Todos  map[int]*Todo `json:"todos"`
	NextId int           `json:"nextId"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	todos := Todos{NextId: 1, Todos: make(map[int]*Todo)}

	fmt.Printf("Welcome to the todo app! \n\n")
	fmt.Printf("Type 'help' for a list of commands.\n\n")
	fmt.Printf("Do you want to load previous todos? (y/n): ")

	load, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.TrimSpace(load) == "y" {
		todos.load()
	}

	for {
		fmt.Println()
		fmt.Println("what would you like to do? (list, toggle, add, clear, exit, help, save)")
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

		case "save":
			todos.save()

		case "load":
			todos.load()

		case "help":
			fmt.Println("add - add a todo")
			fmt.Println("list - list all todos")
			fmt.Println("toggle - toggle a todo")
			fmt.Println("clear - clear all todos")
			fmt.Println("exit - exit the app")
			fmt.Println("save - save all todos to a file")

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

func (t *Todos) add(text string) {
	t.Todos[t.NextId] = &Todo{Id: t.NextId, Text: text, Done: false}
	t.NextId++
}

func (t *Todos) list() {
	if len(t.Todos) == 0 {
		fmt.Println("no todos")
		return
	}

	fmt.Println("\ntodos:")
	for _, todo := range t.Todos {
		status := "[ ]"
		if todo.Done {
			status = "[x]"
		}
		fmt.Printf("%s %d %s\n", status, todo.Id, todo.Text)
	}
}

func (t *Todos) toggle(id int) {
	todo, ok := t.Todos[id]
	if !ok {
		fmt.Printf("no todo with id %d\n", id)
		return
	}
	t.Todos[id].Done = !todo.Done
}

func (t *Todos) clear() {
	t.Todos = make(map[int]*Todo)
}

func (t *Todos) save() {
	out, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile("todos.json", out, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("todos saved")
}

func (t *Todos) load() {
	in, err := os.ReadFile("todos.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.clear()

	err = json.Unmarshal(in, t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("todos loaded")
}
