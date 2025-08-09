package main

import (
	"bufio"
	"fmt"
	"goNote/note"
	"goNote/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

//
//type displayer interface {
//	Display()
//}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	//switch value.(type) {
	//case int:
	//	fmt.Println("int")
	//	break
	//case string:
	//	fmt.Println("string")
	//	break
	//case float64:
	//	fmt.Println("float64")
	//	break
	//default:
	//	fmt.Println("unknown")
	//	break
	//}

	typedValue, ok := value.(int)
	if ok {
		fmt.Println("int", typedValue)
	}
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving failed", err)
		return err
	}

	fmt.Println("Saved successfully")
	return nil
}

func getTodoData() string {
	return getUserInput("todo text: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
