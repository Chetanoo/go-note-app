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

	userTodo.Display()
	err = saveData(userTodo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
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
