package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: content,
	}, nil
}
